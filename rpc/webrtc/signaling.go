package rpcwebrtc

import (
	"context"
	"errors"
	"io"
	"sync"
	"time"

	"github.com/edaniels/golog"
	gwebrtc "github.com/edaniels/gostream/webrtc"
	"go.uber.org/multierr"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	webrtcpb "go.viam.com/core/proto/rpc/webrtc/v1"
	"go.viam.com/core/rpc/dialer"
	"go.viam.com/core/utils"
)

// A SignalingServer implements a signaling service for WebRTC by exchanging
// SDPs (https://webrtcforthecurious.com/docs/02-signaling/#what-is-the-session-description-protocol-sdp)
// via gRPC. The service consists of a many-to-one interaction where there are many callers
// and a single answerer. The callers provide an SDP to the service which asks a waiting
// answerer to provide an SDP in exchange in order to establish a P2P connection between
// the two parties.
type SignalingServer struct {
	webrtcpb.UnimplementedSignalingServiceServer
	callQueue chan callOffer
}

// NewSignalingServer makes a new signaling server that uses an in memory
// call queue.
// TODO(https://github.com/viamrobotics/core/issues/79): abstraction to be able to use
// MongoDB as a distributed call queue. This will enable many signaling services to
// run acting as effectively operators on as switchboard.
func NewSignalingServer() *SignalingServer {
	return &SignalingServer{callQueue: make(chan callOffer)}
}

// callOffer is the offer to start a call where information about the caller
// and how it wishes to speak is contained in the SDP.
type callOffer struct {
	sdp      string
	response chan<- callAnswer
}

// callAnswer is the response to an offer. An agreement to start the call
// will contain an SDP about how the answerer wishes to speak.
type callAnswer struct {
	sdp string
	err error
}

// Call is a request/offer to start a caller with the connected answerer.
func (srv *SignalingServer) Call(ctx context.Context, req *webrtcpb.CallRequest) (*webrtcpb.CallResponse, error) {
	response := make(chan callAnswer)
	offer := callOffer{sdp: req.Sdp, response: response}
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case srv.callQueue <- offer:
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case answer := <-response:
			if answer.err != nil {
				return nil, answer.err
			}
			return &webrtcpb.CallResponse{Sdp: answer.sdp}, nil
		}
	}
}

// Answer listens on call/offer queue forever responding with SDPs to agreed to calls.
// TODO(erd): This should be authenticated to internal servers only with client certificates.
func (srv *SignalingServer) Answer(server webrtcpb.SignalingService_AnswerServer) error {
	for {
		select {
		case <-server.Context().Done():
			return server.Context().Err()
		case offer := <-srv.callQueue:
			if err := server.Send(&webrtcpb.AnswerRequest{Sdp: offer.sdp}); err != nil {
				return err
			}
			answer, err := server.Recv()
			if err != nil {
				offer.response <- callAnswer{err: err}
				return err
			}
			respStatus := status.FromProto(answer.Status)
			if respStatus.Code() != codes.OK {
				offer.response <- callAnswer{err: respStatus.Err()}
				continue
			}
			select {
			case offer.response <- callAnswer{sdp: answer.Sdp}:
			case <-server.Context().Done():
				return server.Context().Err()
			}
		}
	}
}

// A SignalingAnswerer listens for and answers calls with a given signaling service. It is
// directly connected to a Server that will handle the actual calls/connections over WebRTC
// data channels.
type SignalingAnswerer struct {
	address                 string
	client                  webrtcpb.SignalingService_AnswerClient
	server                  *Server
	activeBackgroundWorkers sync.WaitGroup
	cancelBackgroundWorkers func()
	closeCtx                context.Context
	logger                  golog.Logger
}

// NewSignalingAnswerer makes an answerer that will connect to and listen for calls at the given
// address. Note that using this assumes that the connection at the given address is secure and
// assumed that all calls are authenticated. Random ports will be opened on this host to establish
// connections as a means to service ICE (https://webrtcforthecurious.com/docs/03-connecting/#how-does-it-work).
func NewSignalingAnswerer(address string, server *Server, logger golog.Logger) *SignalingAnswerer {
	closeCtx, cancel := context.WithCancel(context.Background())
	return &SignalingAnswerer{
		address:                 address,
		server:                  server,
		cancelBackgroundWorkers: cancel,
		closeCtx:                closeCtx,
		logger:                  logger,
	}
}

// Start connects to the signaling service and listens forever until instructed to stop
// via Stop.
func (ans *SignalingAnswerer) Start() error {
	setupCtx, timeoutCancel := context.WithTimeout(ans.closeCtx, 2*time.Second)
	defer timeoutCancel()

	var conn dialer.ClientConn
	var err error
	dialOpts := []grpc.DialOption{grpc.WithBlock(), grpc.WithInsecure()}
	if ctxDialer := dialer.ContextDialer(setupCtx); ctxDialer != nil {
		conn, err = ctxDialer.Dial(setupCtx, ans.address, dialOpts...)
	} else {
		conn, err = grpc.DialContext(setupCtx, ans.address, dialOpts...)
	}
	if err != nil {
		return err
	}

	client := webrtcpb.NewSignalingServiceClient(conn)
	answerClient, err := client.Answer(ans.closeCtx)
	if err != nil {
		return multierr.Combine(err, conn.Close())
	}
	ans.client = answerClient

	ans.activeBackgroundWorkers.Add(1)
	utils.ManagedGo(func() {
		for {
			select {
			case <-ans.closeCtx.Done():
				return
			default:
			}
			if err := ans.answer(); err != nil && utils.FilterOutError(err, context.Canceled) != nil {
				ans.logger.Errorw("error answering", "error", err)
			}
		}
	}, func() {
		defer ans.activeBackgroundWorkers.Done()
		defer func() {
			if err := conn.Close(); err != nil {
				ans.logger.Errorw("error closing signaling connection", "error", err)
			}
		}()
		defer func() {
			if err := ans.client.CloseSend(); err != nil {
				ans.logger.Errorw("error closing send side of answering client", "error", err)
			}
		}()
	})

	return nil
}

// Stop waits for the answer to stop listening and return.
func (ans *SignalingAnswerer) Stop() {
	ans.cancelBackgroundWorkers()
	ans.activeBackgroundWorkers.Wait()
}

// answer accepts a single call offer, responds with a corresponding SDP, and
// attempts to establish a WebRTC connection with the caller via ICE. Once established,
// the designated WebRTC data channel is passed off to the underlying Server which
// is then used as the server end of a gRPC connection.
func (ans *SignalingAnswerer) answer() (err error) {
	resp, err := ans.client.Recv()
	if err != nil {
		if errors.Is(err, io.EOF) {
			return nil
		}
		return err
	}

	pc, dc, err := newPeerConnectionForServer(ans.closeCtx, resp.Sdp, ans.logger)
	if err != nil {
		return ans.client.Send(&webrtcpb.AnswerResponse{
			Status: ErrorToStatus(err).Proto(),
		})
	}
	var successful bool
	defer func() {
		if !successful {
			err = multierr.Combine(err, pc.Close())
		}
	}()

	encodedSDP, err := gwebrtc.EncodeSDP(pc.LocalDescription())
	if err != nil {
		return ans.client.Send(&webrtcpb.AnswerResponse{
			Status: ErrorToStatus(err).Proto(),
		})
	}

	ans.server.NewChannel(pc, dc)

	successful = true
	return ans.client.Send(&webrtcpb.AnswerResponse{
		Status: ErrorToStatus(nil).Proto(),
		Sdp:    encodedSDP,
	})
}
