//go:build !386 && !arm && !mips && !mipsle

package atomic

import (
	"sync/atomic"

	"github.com/goexl/guc/internal/sync/core"
)

var _ core.Atomic[uint64] = (*Uint64)(nil)

type Uint64 struct {
	value uint64
}

func NewUint64(value uint64) *Uint64 {
	return &Uint64{
		value: value,
	}
}

func (ui *Uint64) Load() uint64 {
	return atomic.LoadUint64(&ui.value)
}

func (ui *Uint64) Store(val uint64) {
	atomic.StoreUint64(&ui.value, val)
}

func (ui *Uint64) Add(delta uint64) uint64 {
	return atomic.AddUint64(&ui.value, delta)
}

func (ui *Uint64) Increment() (new uint64) {
	return atomic.AddUint64(&ui.value, 1)
}

func (ui *Uint64) CompareAndSwap(old uint64, new uint64) bool {
	return atomic.CompareAndSwapUint64(&ui.value, old, new)
}

func (ui *Uint64) Swap(new uint64) uint64 {
	return atomic.SwapUint64(&ui.value, new)
}
