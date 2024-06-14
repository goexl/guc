package factory

import (
	"github.com/goexl/guc/internal/sync/core"
	"github.com/goexl/guc/internal/sync/wait"
)

type Wait struct {
	// 方法
}

func NewWait() *Wait {
	return new(Wait)
}

func (w *Wait) Group(delta int) core.Waiter {
	return wait.NewGroup(delta)
}
