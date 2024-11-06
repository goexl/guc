package queue

import (
	"sync"

	"github.com/goexl/collection"
	"github.com/goexl/guc/internal/internal/param"
)

var _ collection.Queue[int] = (*Blocking[int])(nil)

type Blocking[T any] struct {
	data  []T
	mutex *sync.Mutex
	cond  *sync.Cond

	params *param.Blocking
}

func NewBlocking[T any](params *param.Blocking) (blocking *Blocking[T]) {
	blocking = new(Blocking[T])
	blocking.data = make([]T, 0, params.Capacity)
	blocking.mutex = new(sync.Mutex)
	blocking.cond = sync.NewCond(blocking.mutex)

	blocking.params = params

	return
}

func (b *Blocking[T]) Enqueue(required T, optionals ...T) {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	items := append([]T{required}, optionals...)
	for len(b.data)+len(items) > b.params.Capacity {
		b.cond.Wait()
	}
	b.data = append(b.data, items...)
	b.cond.Broadcast()
}

func (b *Blocking[T]) Dequeue() (items []T) {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	for len(b.data) == 0 {
		b.cond.Wait()
	}
	items = b.data[:]
	b.data = make([]T, 0, b.params.Capacity)
	b.cond.Broadcast()

	return
}

func (b *Blocking[T]) Size() int {
	return len(b.data)
}
