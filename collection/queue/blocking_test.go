package queue_test

import (
	"testing"

	"github.com/goexl/guc/collection/queue"
	"github.com/stretchr/testify/require"
)

func TestBlocking(t *testing.T) {
	blocking := queue.NewBlocking[int]().Build()
	require.NotNil(t, blocking, "阻塞队列创建出错")
}
