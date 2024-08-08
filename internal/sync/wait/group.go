package wait

import (
	"sync"

	"github.com/goexl/guc/internal/sync/mutex"
)

// Group 等待组，是系统sync.WaitGroup的增强版
type Group struct {
	*sync.WaitGroup

	delta int
	mutex sync.Locker
}

func NewGroup(delta int) *Group {
	wg := new(sync.WaitGroup)
	wg.Add(delta)

	return &Group{
		WaitGroup: wg,

		delta: delta,
		mutex: mutex.NewReentrant(),
	}
}

func (g *Group) Add(delta int) {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	g.WaitGroup.Add(delta)
	g.delta = delta
}

func (g *Group) Done() {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	// 不允许出现计数为负的情况
	if 0 >= g.delta {
		return
	}

	g.WaitGroup.Done()
	g.delta--
}

func (g *Group) Wait() {
	g.WaitGroup.Wait()

	g.mutex.Lock()
	defer g.mutex.Unlock()
	g.delta = 0
}

func (g *Group) Completed() bool {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	return 0 >= g.delta
}
