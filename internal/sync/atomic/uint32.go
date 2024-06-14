package atomic

import (
	"sync/atomic"

	"github.com/goexl/guc/internal/sync/core"
)

var _ core.Atomic[uint32] = (*Uint32)(nil)

type Uint32 struct {
	value uint32
}

func NewUint32(value uint32) *Uint32 {
	return &Uint32{
		value: value,
	}
}

func (ui *Uint32) Load() uint32 {
	return atomic.LoadUint32(&ui.value)
}

func (ui *Uint32) Store(val uint32) {
	atomic.StoreUint32(&ui.value, val)
}

func (ui *Uint32) Add(delta uint32) uint32 {
	return atomic.AddUint32(&ui.value, delta)
}

func (ui *Uint32) Increment() (new uint32) {
	return atomic.AddUint32(&ui.value, 1)
}

func (ui *Uint32) CompareAndSwap(old uint32, new uint32) bool {
	return atomic.CompareAndSwapUint32(&ui.value, old, new)
}

func (ui *Uint32) Swap(new uint32) uint32 {
	return atomic.SwapUint32(&ui.value, new)
}
