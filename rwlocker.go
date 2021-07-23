package guc

import (
	`sync`
)

// RWLocker 读写锁
type RWLocker interface {
	sync.Locker

	// RLock 上读锁
	RLock()
	// RUnlock 解读锁
	RUnlock()
}
