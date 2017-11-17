package context

import (
	"context"
	"net"
	"time"

	docker "github.com/fsouza/go-dockerclient"
	"github.com/kr/beanstalk"
	"github.com/nuveo/gofn/iaas"
	"github.com/nuveo/gofn/provision"
	"github.com/nuveo/log"
)

// A Context carries a deadline, cancelation signal, and request-scoped values
// across API boundaries. Its methods are safe for simultaneous use by multiple
// goroutines.
type Context interface {
	// Done returns a channel that is closed when this Context is canceled
	// or times out.
	Done() <-chan struct{}

	// Err indicates why this context was canceled, after the Done channel
	// is closed.
	Err() error

	// Deadline returns the time when this Context will be canceled, if any.
	Deadline() (deadline time.Time, ok bool)

	// Value returns the value associated with key or nil if none.
	Value(key interface{}) interface{}
}

// Background returns an empty Context. It is never canceled, has no deadline,
// and has no values. Background is typically used in main, init, and tests,
// and as the top-level Context for incoming requests.
func Background() Context {}

// WithCancel returns a copy of parent whose Done channel is closed as soon as
// parent.Done is closed or cancel is called.
func WithCancel(parent Context) (ctx Context, cancel CancelFunc) {}

// A CancelFunc cancels a Context.
type CancelFunc func()

// WithDeadline returns a copy of the parent context with the deadline adjusted
// to be no later than d. If the parent's deadline is already earlier than d,
// WithDeadline(parent, d) is semantically equivalent to parent. The returned
// context's Done channel is closed when the deadline expires, when the returned
// cancel function is called, or when the parent context's Done channel is
// closed, whichever happens first.
//
// Canceling this context releases resources associated with it, so code should
// call cancel as soon as the operations running in this Context complete.
func WithDeadline(parent Context, d time.Time) (Context, CancelFunc) {}

// WithTimeout returns WithDeadline(parent, time.Now().Add(timeout)).
//
// Canceling this context releases resources associated with it, so code should
// call cancel as soon as the operations running in this Context complete
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc) {
	return WithDeadline(parent, time.Now().Add(timeout))
}

// WithValue returns a copy of parent in which the value associated with key is
// val.
//
// Use context Values only for request-scoped data that transits processes and
// APIs, not for passing optional parameters to functions.
//
// The provided key must be comparable and should not be of type
// string or any other built-in type to avoid collisions between
// packages using context. Users of WithValue should define their own
// types for keys. To avoid allocating when assigning to an
// interface{}, context keys often have concrete type
// struct{}. Alternatively, exported context key variables' static
// type should be a pointer or interface.
func WithValue(parent Context, key, val interface{}) Context {}

func FromContext(ctx context.Context) (net.IP, bool) {
	// ctx.Value returns nil if ctx has no value for the key;
	// the net.IP type assertion returns ok=false for nil.
	userIP, ok := ctx.Value(userIPKey).(net.IP)
	return userIP, ok
}

// Run runs the designed image
func Run(ctx context.Context, buildOpts *provision.BuildOptions, containerOpts *provision.ContainerOptions) (stdout string, stderr string, err error) {
	var client *docker.Client
	var container *docker.Container
	var machine *iaas.Machine
	done := make(chan struct{})
	go func(ctx context.Context, done chan struct{}) {
		// omited code
		done <- struct{}{}
	}(ctx, done)
	select {
	case <-ctx.Done():
		log.Printf("trying to destroy container %v\n", ctx.Err())
	case <-done:
		log.Println("trying to destroy container process done")
	}
	// omited code
	return
} // fim

func queueLoop(conn *beanstalk.Conn) {
	// omited code
	// ajusta o status do job
	ctx, cancelFunc := context.WithTimeout(context.Background(), item.Timeout)
	// omited code
	if err = sc.Err(); err != nil {
		decrementJobCount(customer)
		log.Errorln(err)
		cancelFunc()
		continue
	}
	// executa o runner

	go func(ctx context.Context, cancelFunc context.CancelFunc) {
		err = runFn(ctx, runJob, item, id, conn)
		if err != nil {
			log.Errorln(err)
		}
		decrementJobCount(customer)
		cancelFunc()
	}(ctx, cancelFunc, customer, item, id, conn)
} // fim
