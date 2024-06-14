package atomic

import (
	"sync/atomic"

	"github.com/goexl/guc/internal/sync/core"
)

var _ core.Atomic[int32] = (*Int32)(nil)

type Int32 struct {
	value int32
}

func NewInt32(value int32) *Int32 {
	return &Int32{
		value: value,
	}
}

func (i *Int32) Load() int32 {
	return atomic.LoadInt32(&i.value)
}

func (i *Int32) Store(val int32) {
	atomic.StoreInt32(&i.value, val)
}

func (i *Int32) Add(delta int32) int32 {
	return atomic.AddInt32(&i.value, delta)
}

func (i *Int32) Increment() int32 {
	return atomic.AddInt32(&i.value, 1)
}

func (i *Int32) CompareAndSwap(old int32, new int32) bool {
	return atomic.CompareAndSwapInt32(&i.value, old, new)
}

func (i *Int32) Swap(new int32) int32 {
	return atomic.SwapInt32(&i.value, new)
}
