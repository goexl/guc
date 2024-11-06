package queue

import (
	"github.com/goexl/guc/internal/builder"
)

func NewBlocking[T any]() *builder.Blocking[T] {
	return builder.NewBlocking[T]()
}
