package builder

import (
	"github.com/goexl/guc/internal/collection/queue"
	"github.com/goexl/guc/internal/internal/param"
)

type Blocking[T any] struct {
	params *param.Blocking
}

func NewBlocking[T any]() *Blocking[T] {
	return &Blocking[T]{
		params: param.NewBlocking(),
	}
}

func (b *Blocking[T]) Capacity(capacity int) (blocking *Blocking[T]) {
	b.params.Capacity = capacity
	blocking = b

	return
}

func (b *Blocking[T]) Build() *queue.Blocking[T] {
	return queue.NewBlocking[T](b.params)
}
