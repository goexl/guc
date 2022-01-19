package guc

import (
	`sync`
)

// WaitGroup 等待组，是系统sync.WaitGroup的增强版
type WaitGroup struct {
	sync.WaitGroup

	delta int
	mutex reentrantMutex
}

func (wg *WaitGroup) Add(delta int) {
	wg.WaitGroup.Add(delta)

	wg.mutex.Lock()
	defer wg.mutex.Unlock()
	wg.delta = delta
}

// Done 完成
func (wg *WaitGroup) Done() {
	wg.WaitGroup.Done()

	wg.mutex.Lock()
	defer wg.mutex.Unlock()
	wg.delta--
}

// Wait 等待
func (wg *WaitGroup) Wait() {
	wg.WaitGroup.Wait()

	wg.mutex.Lock()
	defer wg.mutex.Unlock()
	wg.delta = 0
}

// Completed 是否已经完成
func (wg *WaitGroup) Completed() bool {
	wg.mutex.Lock()
	defer wg.mutex.Unlock()

	return 0 >= wg.delta
}
