package core

import (
	"github.com/goexl/guc/internal/sync/limit"
)

type Atomic[T limit.Atomic] interface {
	Load() T
	Store(T)
	Add(T) T
	Increment() T
	CompareAndSwap(T, T) bool
	Swap(T) T
}
