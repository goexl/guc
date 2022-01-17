package guc

import (
	`sync`
)

// WaitGroup 等待组，是系统sync.WaitGroup的增强版
type WaitGroup struct {
	sync.WaitGroup

	completed bool
	mutex     sync.RWMutex
}

// Wait 等待
func (wg *WaitGroup) Wait() {
	wg.WaitGroup.Wait()

	// 结束后，置完成状态
	wg.mutex.Lock()
	defer wg.mutex.Unlock()
	wg.completed = true
}

// Completed 是否已经完成
func (wg *WaitGroup) Completed() bool {
	wg.mutex.RLock()
	defer wg.mutex.RUnlock()

	return wg.completed
}
