package main

import (
	"context"
	"io"

	"github.com/edaniels/golog"
	"github.com/go-errors/errors"

	echopb "go.viam.com/core/proto/rpc/examples/echo/v1"
	"go.viam.com/core/rlog"
	rpcwebrtc "go.viam.com/core/rpc/webrtc"
	"go.viam.com/core/utils"
)

func main() {
	utils.ContextualMain(mainWithArgs, logger)
}

var logger = rlog.Logger.Named("client")

// Arguments for the command.
type Arguments struct {
	SignalingAddress string `flag:"0,default=localhost:8080"`
}

func mainWithArgs(ctx context.Context, args []string, logger golog.Logger) error {
	var argsParsed Arguments
	if err := utils.ParseFlags(args, &argsParsed); err != nil {
		return err
	}

	cc, err := rpcwebrtc.Dial(ctx, argsParsed.SignalingAddress, logger)
	if err != nil {
		return err
	}
	defer cc.Close()

	echoClient := echopb.NewEchoServiceClient(cc)
	resp, err := echoClient.Echo(ctx, &echopb.EchoRequest{Message: "hello?"})
	if err != nil {
		return errors.Wrap(err, 0)
	}
	logger.Infow("echo", "resp", resp.Message)

	multiClient, err := echoClient.EchoMultiple(ctx, &echopb.EchoMultipleRequest{Message: "hello?"})
	if err != nil {
		return errors.Wrap(err, 0)
	}
	for {
		resp, err := multiClient.Recv()
		if err != nil {
			if !errors.Is(err, io.EOF) {
				return errors.Wrap(err, 0)
			}
			break
		}
		logger.Infow("echo multi", "resp", resp.Message)
	}

	biDiClient, err := echoClient.EchoBiDi(ctx)
	if err != nil {
		return errors.Wrap(err, 0)
	}
	biDiClient.Send(&echopb.EchoBiDiRequest{Message: "one"})
	for i := 0; i < 3; i++ {
		resp, err := biDiClient.Recv()
		if err != nil {
			if !errors.Is(err, io.EOF) {
				return errors.Wrap(err, 0)
			}
			break
		}
		logger.Infow("echo bidi", "resp", resp.Message)
	}

	biDiClient.Send(&echopb.EchoBiDiRequest{Message: "two"})

	for i := 0; i < 3; i++ {
		resp, err := biDiClient.Recv()
		if err != nil {
			if !errors.Is(err, io.EOF) {
				return errors.Wrap(err, 0)
			}
			break
		}
		logger.Infow("echo bidi", "resp", resp.Message)
	}

	biDiClient.CloseSend()

	return nil
}
