package guc

import (
	"github.com/goexl/guc/internal"
)

func New() *internal.Factory {
	return internal.NewFactory()
}
