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
	owner     uint64
	recursion int32
	padding   [4]byte // !字节对齐，填充对齐间隙
	mutex     *sync.RWMutex
}

// NewReentrantRWMutex 创建新的可重入锁
func NewReentrantRWMutex() (locker RWLocker) {
	rm := &reentrantRWMutex{
		owner:     0,
		recursion: 0,
		mutex:     new(sync.RWMutex),
	}
	locker = rm

	return
}

func (m *reentrantRWMutex) Lock() {
	id := gid()
	if id == atomic.LoadUint64(&m.owner) {
		m.recursion++
	} else {
		m.mutex.Lock()
		atomic.StoreUint64(&m.owner, id)
		m.recursion = 1
	}
}

func (m *reentrantRWMutex) Unlock() {
	id := gid()
	if id != atomic.LoadUint64(&m.owner) {
		panic(fmt.Sprintf("错误的协程持有者（%d）：%d！", m.owner, id))
	}

	m.recursion--
	if 0 != m.recursion {
		return
	}

	atomic.StoreUint64(&m.owner, 0)
	m.mutex.Unlock()
}

func (m *reentrantRWMutex) RLock() {
	id := gid()
	if id == atomic.LoadUint64(&m.owner) {
		m.recursion++
	} else {
		m.mutex.RLock()
		atomic.StoreUint64(&m.owner, id)
		m.recursion = 1
	}
}

func (m *reentrantRWMutex) RUnlock() {
	id := gid()
	if id != atomic.LoadUint64(&m.owner) {
		panic(fmt.Sprintf("错误的协程持有者（%d）：%d！", m.owner, id))
	}

	m.recursion--
	if 0 != m.recursion {
		return
	}

	atomic.StoreUint64(&m.owner, 0)
	m.mutex.RUnlock()
}
