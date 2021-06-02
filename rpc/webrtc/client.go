// Package rpcwebrtc providers client/server functionality for gRPC serviced over
// WebRTC data channels. The work is adapted from https://github.com/jsmouret/grpc-over-webrtc.
package rpcwebrtc

import (
	"context"
	"time"

	"github.com/edaniels/golog"
	gwebrtc "github.com/edaniels/gostream/webrtc"
	"github.com/pion/webrtc/v3"
	"go.uber.org/multierr"
	"google.golang.org/grpc"

	webrtcpb "go.viam.com/core/proto/rpc/webrtc/v1"
	"go.viam.com/core/rpc/dialer"
)

// Dial connects to the signaling service at the given address and attempts to establish
// a WebRTC connection with the corresponding peer reflected in the address.
func Dial(ctx context.Context, address string, logger golog.Logger) (ch *ClientChannel, err error) {
	dialCtx, timeoutCancel := context.WithTimeout(ctx, 20*time.Second)
	defer timeoutCancel()

	logger.Debugw("connecting to signaling server", "address", address)

	var conn dialer.ClientConn
	{
		var err error
		dialOpts := []grpc.DialOption{grpc.WithBlock(), grpc.WithInsecure()}
		if ctxDialer := dialer.ContextDialer(dialCtx); ctxDialer != nil {
			conn, err = ctxDialer.Dial(dialCtx, address, dialOpts...)
		} else {
			conn, err = grpc.DialContext(dialCtx, address, dialOpts...)
		}
		if err != nil {
			return nil, err
		}
	}
	defer func() {
		err = multierr.Combine(err, conn.Close())
	}()

	logger.Debug("connected")

	signalingClient := webrtcpb.NewSignalingServiceClient(conn)

	pc, dc, err := newPeerConnectionForClient(ctx, logger)
	if err != nil {
		return nil, err
	}
	var successful bool
	defer func() {
		if !successful {
			err = multierr.Combine(err, pc.Close())
		}
	}()

	encodedSDP, err := gwebrtc.EncodeSDP(pc.LocalDescription())
	if err != nil {
		return nil, err
	}

	answerResp, err := signalingClient.Call(ctx, &webrtcpb.CallRequest{Sdp: encodedSDP})
	if err != nil {
		return nil, err
	}

	answer := webrtc.SessionDescription{}
	if err := gwebrtc.DecodeSDP(answerResp.Sdp, &answer); err != nil {
		return nil, err
	}

	err = pc.SetRemoteDescription(answer)
	if err != nil {
		return nil, err
	}

	clientCh := NewClientChannel(pc, dc, logger)

	select {
	case <-ctx.Done():
		return nil, multierr.Combine(ctx.Err(), clientCh.Close())
	case <-clientCh.Ready():
	}
	successful = true
	return clientCh, nil
}
