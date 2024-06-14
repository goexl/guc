package guc

import (
	"github.com/goexl/guc/internal/sync/core"
	"github.com/goexl/guc/internal/sync/limit"
)

// Atomic 原子操作
type Atomic[T limit.Atomic] interface {
	core.Atomic[T]
}
