package guc

import (
	`fmt`
	`sync`
	`sync/atomic`
)

var (
	_             = NewReentrantMutex
	_ sync.Locker = (*reentrantMutex)(nil)
)

type reentrantMutex struct {
	mutex     sync.Mutex
	owner     uint64
	recursion int32
}

// NewReentrantMutex 创建新的可重入锁
func NewReentrantMutex() (locker sync.Locker) {
	return &reentrantMutex{
		mutex:     sync.Mutex{},
		owner:     0,
		recursion: 0,
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
