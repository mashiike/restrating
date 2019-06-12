package usecase_test

import (
	"context"
	"testing"
)

type testLifeCycle struct {
	t       *testing.T
	lastErr error
}

func newLifeCycle(t *testing.T) *testLifeCycle {
	return &testLifeCycle{
		t: t,
	}
}

func (lc *testLifeCycle) Begin(ctx context.Context) context.Context {
	lc.t.Log("start life cycle")
	return ctx
}

func (lc *testLifeCycle) End(ctx context.Context, err error) {
	if err == nil {
		lc.t.Log("end life cycle")
	} else {
		lc.t.Logf("end life cycle with error = %v", err)
	}
	lc.lastErr = err
}
