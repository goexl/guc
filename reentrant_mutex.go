package guc

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var (
	_             = NewReentrantMutex
	_ sync.Locker = (*reentrantMutex)(nil)
)

type reentrantMutex struct {
	recursion int32
	owner     uint64
	mutex     *sync.Mutex
}

// NewReentrantMutex 创建新的可重入锁
func NewReentrantMutex() sync.Locker {
	return &reentrantMutex{
		recursion: 0,
		owner:     0,
		mutex:     new(sync.Mutex),
	}
}

func (rm *reentrantMutex) Lock() {
	id := gid()
	if id == atomic.LoadUint64(&rm.owner) {
		rm.recursion++
	} else {
		rm.mutex.Lock()
		atomic.StoreUint64(&rm.owner, id)
		rm.recursion = 1
	}
}

func (rm *reentrantMutex) Unlock() {
	id := gid()
	if id != atomic.LoadUint64(&rm.owner) {
		panic(fmt.Sprintf("错误的协程持有者（%d）：%d！", rm.owner, id))
	}

	rm.recursion--
	if 0 < rm.recursion {
		return
	}

	atomic.StoreUint64(&rm.owner, 0)
	rm.mutex.Unlock()
}
