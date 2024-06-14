package guc

import (
	"github.com/goexl/guc/internal/builder"
)

func Lock() *builder.Lock {
	return builder.NewLock()
}

func Wait() *builder.Wait {
	return builder.NewWait()
}
