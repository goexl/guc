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
	owner     uint64
	recursion int32
	padding   [4]byte // !字节对齐，填充对齐间隙
	mutex     *sync.Mutex
}

// NewReentrantMutex 创建新的可重入锁
func NewReentrantMutex() sync.Locker {
	return &reentrantMutex{
		owner:     0,
		recursion: 0,
		mutex:     new(sync.Mutex),
	}
}

func (m *reentrantMutex) Lock() {
	id := gid()
	if id == atomic.LoadUint64(&m.owner) {
		m.recursion++
	} else {
		m.mutex.Lock()
		atomic.StoreUint64(&m.owner, id)
		m.recursion = 1
	}
}

func (m *reentrantMutex) Unlock() {
	id := gid()
	if id != atomic.LoadUint64(&m.owner) {
		panic(fmt.Sprintf("错误的协程持有者（%d）：%d！", m.owner, id))
	}

	m.recursion--
	if 0 < m.recursion {
		return
	}

	atomic.StoreUint64(&m.owner, 0)
	m.mutex.Unlock()
}
