package mutex

import (
	"fmt"
	"sync"

	"github.com/goexl/guc/internal/sync/atomic"
	"github.com/goexl/guc/internal/sync/core"
)

var _ sync.Locker = (*Reentrant)(nil)

type Reentrant struct {
	owner     core.Atomic[uint64]
	recursion int32
	mutex     *sync.Mutex
}

// NewReentrant 创建新的可重入锁
func NewReentrant() sync.Locker {
	return &Reentrant{
		owner:     atomic.NewUint64(0),
		recursion: 0,
		mutex:     new(sync.Mutex),
	}
}

func (r *Reentrant) Lock() {
	id := core.Gid()
	if id == r.owner.Load() {
		r.recursion++
	} else {
		r.mutex.Lock()
		r.owner.Store(id)
		r.recursion = 1
	}
}

func (r *Reentrant) Unlock() {
	id := core.Gid()
	if id != r.owner.Load() {
		panic(fmt.Sprintf("错误的协程持有者（%d）：%d！", r.owner, id))
	}

	r.recursion--
	if 0 < r.recursion {
		return
	}

	r.owner.Store(0)
	r.mutex.Unlock()
}
