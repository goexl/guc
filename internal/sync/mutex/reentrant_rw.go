package mutex

import (
	"fmt"
	"sync"

	"github.com/goexl/guc/internal/sync/atomic"
	"github.com/goexl/guc/internal/sync/core"
)

var _ core.RWLocker = (*ReentrantRW)(nil)

type ReentrantRW struct {
	owner     core.Atomic[uint64]
	recursion int32
	mutex     sync.RWMutex
}

// NewReentrantRW 创建新的可重入锁
func NewReentrantRW() core.RWLocker {
	return &ReentrantRW{
		owner:     atomic.NewUint64(0),
		recursion: 0,
		mutex:     sync.RWMutex{},
	}
}

func (rw *ReentrantRW) Lock() {
	id := core.Gid()
	if id == rw.owner.Load() {
		rw.recursion++
	} else {
		rw.mutex.Lock()
		rw.owner.Store(id)
		rw.recursion = 1
	}
}

func (rw *ReentrantRW) Unlock() {
	id := core.Gid()
	if id != rw.owner.Load() {
		panic(fmt.Sprintf("错误的协程持有者（%d）：%d！", rw.owner, id))
	}

	rw.recursion--
	if 0 != rw.recursion {
		return
	}

	rw.owner.Store(0)
	rw.mutex.Unlock()
}

func (rw *ReentrantRW) RLock() {
	id := core.Gid()
	if id == rw.owner.Load() {
		rw.recursion++
	} else {
		rw.mutex.RLock()
		rw.owner.Store(id)
		rw.recursion = 1
	}
}

func (rw *ReentrantRW) RUnlock() {
	id := core.Gid()
	if id != rw.owner.Load() {
		panic(fmt.Sprintf("错误的协程持有者（%d）：%d！", rw.owner, id))
	}

	rw.recursion--
	if 0 != rw.recursion {
		return
	}

	rw.owner.Store(0)
	rw.mutex.RUnlock()
}
