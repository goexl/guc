package internal

import (
	"github.com/goexl/guc/internal/factory"
)

type Factory struct {
	// 方法
}

func NewFactory() *Factory {
	return new(Factory)
}

func (f *Factory) Lock() *factory.Lock {
	return factory.NewLock()
}

func (f *Factory) Wait() *factory.Wait {
	return factory.NewWait()
}
