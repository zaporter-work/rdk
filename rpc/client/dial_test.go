package client_test

import (
	"context"
	"net"
	"testing"
	"time"

	"github.com/edaniels/golog"
	"go.viam.com/test"

	pb "go.viam.com/core/proto/rpc/examples/echo/v1"
	"go.viam.com/core/rpc/client"
	echoserver "go.viam.com/core/rpc/examples/echo/server"
	"go.viam.com/core/rpc/server"
)

// TODO(erd): test local.* once available
func TestDial(t *testing.T) {
	logger := golog.NewTestLogger(t)

	// pure failure cases
	_, err := client.Dial(context.Background(), "::", client.DialOptions{}, logger)
	test.That(t, err, test.ShouldNotBeNil)
	test.That(t, err.Error(), test.ShouldContainSubstring, "too many")

	ctx1, cancel1 := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel1()
	_, err = client.Dial(ctx1, "127.0.0.1:1", client.DialOptions{}, logger)
	test.That(t, err, test.ShouldResemble, context.DeadlineExceeded)

	ctx2, cancel2 := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel2()
	_, err = client.Dial(ctx2, "host.unknown", client.DialOptions{}, logger)
	test.That(t, err, test.ShouldResemble, context.DeadlineExceeded)

	// working and fallbacks

	rpcServer, err := server.New(logger)
	test.That(t, err, test.ShouldBeNil)

	es := echoserver.Server{}
	err = rpcServer.RegisterServiceServer(
		context.Background(),
		&pb.EchoService_ServiceDesc,
		&es,
		pb.RegisterEchoServiceHandlerFromEndpoint,
	)
	test.That(t, err, test.ShouldBeNil)

	httpListener, err := net.Listen("tcp", "localhost:0")
	test.That(t, err, test.ShouldBeNil)

	errChan := make(chan error)
	go func() {
		errChan <- rpcServer.Serve(httpListener)
	}()

	conn, err := client.Dial(context.Background(), httpListener.Addr().String(), client.DialOptions{}, logger)
	test.That(t, err, test.ShouldBeNil)
	test.That(t, conn.Close(), test.ShouldBeNil)

	test.That(t, rpcServer.Stop(), test.ShouldBeNil)
	err = <-errChan
	test.That(t, err, test.ShouldBeNil)

	rpcServer, err = server.NewWithOptions(server.Options{server.WebRTCOptions{Enable: true}}, logger)
	test.That(t, err, test.ShouldBeNil)

	err = rpcServer.RegisterServiceServer(
		context.Background(),
		&pb.EchoService_ServiceDesc,
		&es,
		pb.RegisterEchoServiceHandlerFromEndpoint,
	)
	test.That(t, err, test.ShouldBeNil)

	httpListener, err = net.Listen("tcp", "localhost:0")
	test.That(t, err, test.ShouldBeNil)

	errChan = make(chan error)
	go func() {
		errChan <- rpcServer.Serve(httpListener)
	}()

	conn, err = client.Dial(context.Background(), httpListener.Addr().String(), client.DialOptions{}, logger)
	test.That(t, err, test.ShouldBeNil)
	test.That(t, conn.Close(), test.ShouldBeNil)

	test.That(t, rpcServer.Stop(), test.ShouldBeNil)
	err = <-errChan
	test.That(t, err, test.ShouldBeNil)
}
