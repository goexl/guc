package core

type Waiter interface {
	Add(int)
	Done()
	Wait()
	Completed() bool
}
