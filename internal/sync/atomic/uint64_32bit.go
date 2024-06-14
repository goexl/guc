//go:build 386 || arm || mips || mipsle
// +build 386 arm mips mipsle

package atomic

import (
	"sync"

	"github.com/goexl/guc/internal/sync/core"
)

var _ core.Atomic[uint64] = (*Uint64)(nil)

type Uint64 struct {
	mutex sync.Mutex
	value uint64
}

func NewUint64(value uint64) *Uint64 {
	return &Uint64{
		value: value,
	}
}

func (ui *Uint64) Load() (value uint64) {
	ui.mutex.Lock()
	value = ui.value
	ui.mutex.Unlock()

	return
}

func (ui *Uint64) Store(value uint64) {
	ui.mutex.Lock()
	ui.value = value
	ui.mutex.Unlock()
}

func (ui *Uint64) Add(delta uint64) (new uint64) {
	ui.mutex.Lock()
	ui.value += delta
	new = ui.value
	ui.mutex.Unlock()

	return
}

func (ui *Uint64) Increment() (new uint64) {
	ui.mutex.Lock()
	ui.value += 1
	new = ui.value
	ui.mutex.Unlock()

	return
}

func (ui *Uint64) CompareAndSwap(old uint64, new uint64) (swapped bool) {
	ui.mutex.Lock()
	if ui.value == old {
		ui.value = new
		ui.mutex.Unlock()
		swapped = true
	} else {
		swapped = false
	}
	ui.mutex.Unlock()

	return
}

func (ui *Uint64) Swap(new uint64) (old uint64) {
	ui.mutex.Lock()
	old = ui.value
	ui.value = new
	ui.mutex.Unlock()

	return
}
