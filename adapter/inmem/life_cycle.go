package inmem

import "context"

type LifeCycle struct{}

func NewLifeCycle() *LifeCycle {
	return &LifeCycle{}
}

func (lc *LifeCycle) Begin(ctx context.Context) context.Context {
	return ctx
}

func (lc *LifeCycle) End(_ context.Context, _ error) {}
