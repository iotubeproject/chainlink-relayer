package runner

import (
	"context"
	"time"
)

type ErrHandler func(error)

type ProducerConsumer interface {
	Produce(context.Context) error
	Consume(context.Context) error
}

type ProducerConsumerRunner struct {
	ctx        context.Context
	r          Runner
	pc         ProducerConsumer
	errHandler ErrHandler
}

func NewProducerConsumerRunner(
	interval time.Duration,
	pc ProducerConsumer,
	errHandler ErrHandler,
) (*ProducerConsumerRunner, error) {
	pcr := &ProducerConsumerRunner{
		ctx:        context.Background(),
		pc:         pc,
		errHandler: errHandler,
	}
	r, err := NewRunner(interval, pcr.run)
	if err != nil {
		return nil, err
	}
	pcr.r = r

	return pcr, nil
}

func (pcr *ProducerConsumerRunner) run() {
	if err := pcr.pc.Produce(pcr.ctx); err != nil {
		pcr.errHandler(err)
	}
	if err := pcr.pc.Consume(pcr.ctx); err != nil {
		pcr.errHandler(err)
	}
}

func (pcr *ProducerConsumerRunner) Start(ctx context.Context) error {
	pcr.ctx = ctx
	return pcr.r.Start(ctx)
}

func (pcr *ProducerConsumerRunner) Stop(ctx context.Context) error {
	return pcr.r.Stop(ctx)
}
