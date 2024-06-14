//go:build !386 && !arm && !mips && !mipsle

package atomic

import (
	"sync/atomic"

	"github.com/goexl/guc/internal/sync/core"
)

var _ core.Atomic[int64] = (*Int64)(nil)

type Int64 struct {
	value int64
}

func NewInt64(value int64) *Int64 {
	return &Int64{
		value: value,
	}
}

func (i *Int64) Load() int64 {
	return atomic.LoadInt64(&i.value)
}

func (i *Int64) Store(val int64) {
	atomic.StoreInt64(&i.value, val)
}

func (i *Int64) Add(delta int64) int64 {
	return atomic.AddInt64(&i.value, delta)
}

func (i *Int64) Increment() int64 {
	return atomic.AddInt64(&i.value, 1)
}

func (i *Int64) CompareAndSwap(old int64, new int64) bool {
	return atomic.CompareAndSwapInt64(&i.value, old, new)
}

func (i *Int64) Swap(new int64) int64 {
	return atomic.SwapInt64(&i.value, new)
}
