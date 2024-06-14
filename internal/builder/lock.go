package builder

import (
	"sync"

	"github.com/goexl/guc/internal/sync/core"
	"github.com/goexl/guc/internal/sync/mutex"
)

type Lock struct {
	// 方法
}

func NewLock() *Lock {
	return new(Lock)
}

func (l *Lock) Reentrant() sync.Locker {
	return mutex.NewReentrant()
}

func (l *Lock) ReentrantRW() core.RWLocker {
	return mutex.NewReentrantRW()
}
