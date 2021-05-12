package rpc

import (
	"context"
	"sync"

	"go.uber.org/multierr"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// A Dialer is responsible for making connections to gRPC endpoints.
type Dialer interface {
	// Dial makes a connection to the given target with the supplied options.
	Dial(ctx context.Context, target string, opts ...grpc.DialOption) (ClientConn, error)

	// Close ensures all connections made are cleanly closed.
	Close() error
}

// A ClientConn is a wrapper around the gRPC client connection interface but ensures
// there is a way to close the connection.
type ClientConn interface {
	grpc.ClientConnInterface
	Close() error
}

type ctxKey int

const ctxKeyDialer = ctxKey(iota)

// ContextWithDialer attaches a Dialer to the given context.
func ContextWithDialer(ctx context.Context, d Dialer) context.Context {
	return context.WithValue(ctx, ctxKeyDialer, d)
}

// ContextDialer returns a Dialer. It may be nil if the value was never set.
func ContextDialer(ctx context.Context) Dialer {
	dialer := ctx.Value(ctxKeyDialer)
	if dialer == nil {
		return nil
	}
	return dialer.(Dialer)
}

type cachedDialer struct {
	mu    sync.Mutex // Note(erd): not suitable for highly concurrent usage
	conns map[string]*refCountedConnWrapper
}

// NewCachedDialer returns a Dialer that returns the same connection if it
// already has been established at a particular target (regardless of the
// options used).
func NewCachedDialer() Dialer {
	return &cachedDialer{conns: map[string]*refCountedConnWrapper{}}
}

func (cd *cachedDialer) Dial(ctx context.Context, target string, opts ...grpc.DialOption) (ClientConn, error) {
	cd.mu.Lock()
	c, ok := cd.conns[target]
	cd.mu.Unlock()
	if ok {
		return c.Ref(), nil
	}

	// assume any difference in opts does not matter
	conn, err := grpc.DialContext(ctx, target, opts...)
	if err != nil {
		return nil, err
	}
	refConn := newRefCountedConnWrapper(conn)
	cd.mu.Lock()
	defer cd.mu.Unlock()

	// someone else might have already connected
	c, ok = cd.conns[target]
	if ok {
		if err := conn.Close(); err != nil {
			return nil, err
		}
		return c.Ref(), nil
	}
	cd.conns[target] = refConn
	return refConn.Ref(), nil
}

func (cd *cachedDialer) Close() error {
	cd.mu.Lock()
	defer cd.mu.Unlock()
	var err error
	for _, c := range cd.conns {
		if closeErr := c.actual.Close(); closeErr != nil && status.Convert(closeErr).Code() != codes.Canceled {
			err = multierr.Combine(err, closeErr)
		}
	}
	return err
}

func newRefCountedConnWrapper(conn ClientConn) *refCountedConnWrapper {
	return &refCountedConnWrapper{NewRefCountedValue(conn), conn}
}

type refCountedConnWrapper struct {
	ref    RefCountedValue
	actual ClientConn
}

func (w *refCountedConnWrapper) Ref() ClientConn {
	return &reffedConn{ClientConn: w.ref.Ref().(ClientConn), deref: w.ref.Deref}
}

type reffedConn struct {
	ClientConn
	derefOnce sync.Once
	deref     func() bool
}

// Close will deref the reference and if it is the last to do so, will close
// the underlying connection.
func (rc *reffedConn) Close() error {
	var err error
	rc.derefOnce.Do(func() {
		if unref := rc.deref(); unref {
			if closeErr := rc.ClientConn.Close(); closeErr != nil && status.Convert(closeErr).Code() != codes.Canceled {
				err = closeErr
			}
		}
	})
	return err
}

// RefCountedValue is a utility to "reference count" values in order
// to destruct them once no one references them.
// If you don't require that kind of logic, just rely on golang's
// garbage collection.
type RefCountedValue interface {
	// Ref increments the reference count and returns the value.
	Ref() interface{}

	// Deref decrements the reference count and returns if this
	// dereference resulted in the value being unreferenced.
	Deref() (unreferenced bool)
}

type refCountedValue struct {
	mu    sync.Mutex
	count int
	val   interface{}
}

// NewRefCountedValue returns a new reference counted value for the given
// value. Its reference count starts at zero but is not released. It is
// assumed the caller of this will reference it at least once.
func NewRefCountedValue(val interface{}) RefCountedValue {
	return &refCountedValue{val: val}
}

func (rcv *refCountedValue) Ref() interface{} {
	rcv.mu.Lock()
	defer rcv.mu.Unlock()
	if rcv.count == -1 {
		panic("already released")
	}
	rcv.count++
	return rcv.val
}

func (rcv *refCountedValue) Deref() bool {
	rcv.mu.Lock()
	defer rcv.mu.Unlock()
	if rcv.count <= 0 {
		panic("deref when count already zero")
	}
	rcv.count--
	if rcv.count == 0 {
		rcv.count = -1
		return true
	}
	return false
}
