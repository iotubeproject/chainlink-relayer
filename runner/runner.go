package runner

import (
	"context"
	"errors"
	"time"
)

var (
	ErrNegTime    = errors.New("wait time cannot be negative")
	ErrNilHandler = errors.New("handler cannot be nil")
)

type (
	// RunFunc is the handler to run
	RunFunc func()

	// Runner defines an interface which calls a callback function periodically
	Runner interface {
		// Start starts the runner
		Start(context.Context) error
		// Stop signals the runner to quit
		Stop(context.Context) error
	}

	// runner implements the Runner interface
	runner struct {
		start chan struct{} // signal to start
		quit  chan struct{} // signal to quit
		wait  time.Duration // wait time before run next round
		run   RunFunc
	}
)

// NewRunner creates a new runner with a duration and a callback function
func NewRunner(wait time.Duration, run RunFunc) (Runner, error) {
	if wait < 0 {
		return nil, ErrNegTime
	}
	if run == nil {
		return nil, ErrNilHandler
	}
	r := runner{
		start: make(chan struct{}),
		quit:  make(chan struct{}),
		wait:  wait,
	}

	go func() {
		<-r.start
		for {
			select {
			case <-r.quit:
				return
			default:
				run()
				time.Sleep(r.wait)
			}
		}
	}()
	return &r, nil
}

// Start starts the runner
func (r *runner) Start(ctx context.Context) error {
	r.start <- struct{}{}
	return nil
}

// Close signals the runner to quit
func (r *runner) Stop(context.Context) error {
	close(r.quit)
	return nil
}
