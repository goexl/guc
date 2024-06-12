package guc

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var (
	_          = NewReentrantRWMutex
	_ RWLocker = (*reentrantRWMutex)(nil)
)

type reentrantRWMutex struct {
	recursion int32
	owner     uint64
	mutex     *sync.RWMutex
}

// NewReentrantRWMutex 创建新的可重入锁
func NewReentrantRWMutex() (locker RWLocker) {
	rm := &reentrantRWMutex{
		recursion: 0,
		owner:     0,
		mutex:     new(sync.RWMutex),
	}
	locker = rm

	return
}

func (rrm *reentrantRWMutex) Lock() {
	gid := gid()
	if gid == atomic.LoadUint64(&rrm.owner) {
		rrm.recursion++
	} else {
		rrm.mutex.Lock()
		atomic.StoreUint64(&rrm.owner, gid)
		rrm.recursion = 1
	}
}

func (rrm *reentrantRWMutex) Unlock() {
	gid := gid()
	if gid != atomic.LoadUint64(&rrm.owner) {
		panic(fmt.Sprintf("错误的协程持有者（%d）：%d！", rrm.owner, gid))
	}

	rrm.recursion--
	if 0 != rrm.recursion {
		return
	}

	atomic.StoreUint64(&rrm.owner, 0)
	rrm.mutex.Unlock()
}

func (rrm *reentrantRWMutex) RLock() {
	gid := gid()
	if gid == atomic.LoadUint64(&rrm.owner) {
		rrm.recursion++
	} else {
		rrm.mutex.RLock()
		atomic.StoreUint64(&rrm.owner, gid)
		rrm.recursion = 1
	}
}

func (rrm *reentrantRWMutex) RUnlock() {
	gid := gid()
	if gid != atomic.LoadUint64(&rrm.owner) {
		panic(fmt.Sprintf("错误的协程持有者（%d）：%d！", rrm.owner, gid))
	}

	rrm.recursion--
	if 0 != rrm.recursion {
		return
	}

	atomic.StoreUint64(&rrm.owner, 0)
	rrm.mutex.RUnlock()
}
