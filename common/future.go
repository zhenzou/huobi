package common

import (
	"context"
)

func NewFuture(data interface{}) *Future {
	return &Future{
		Data: data,
		ch:   make(chan struct{}),
	}
}

type Future struct {
	Data interface{}
	ch   chan struct{}
	err  error
}

func (f *Future) Resolve() {
	close(f.ch)
}

func (f *Future) Reject(err error) {
	f.err = err
	close(f.ch)
}

func (f *Future) Error() error {
	return f.err
}

func (f *Future) Await(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-f.ch:
		return f.err
	}
}
