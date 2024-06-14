//go:build 386 || arm || mips || mipsle
// +build 386 arm mips mipsle

package atomic

import (
	"sync"

	"github.com/goexl/guc/internal/sync/core"
)

var _ core.Atomic[int64] = (*Int64)(nil)

type Int64 struct {
	mutex *sync.Mutex
	value int64
}

func NewInt64(value int64) *Int64 {
	return &Int64{
		value: value,
		mutex: new(sync.Mutex),
	}
}

func (i *Int64) Load() (value int64) {
	i.mutex.Lock()
	value = i.value
	i.mutex.Unlock()

	return
}

func (i *Int64) Store(value int64) {
	i.mutex.Lock()
	i.value = value
	i.mutex.Unlock()
}

func (i *Int64) Add(delta int64) (new int64) {
	i.mutex.Lock()
	i.value += delta
	new = i.value
	i.mutex.Unlock()

	return
}

func (i *Int64) Increment() (new int64) {
	i.mutex.Lock()
	i.value += 1
	new = i.value
	i.mutex.Unlock()

	return
}

func (i *Int64) CompareAndSwap(old int64, new int64) (swapped bool) {
	i.mutex.Lock()
	if i.value == old {
		i.value = new
		i.mutex.Unlock()
		swapped = true
	} else {
		swapped = false
	}
	i.mutex.Unlock()

	return
}

func (i *Int64) Swap(new int64) (old int64) {
	i.mutex.Lock()
	old = i.value
	i.value = new
	i.mutex.Unlock()

	return
}
