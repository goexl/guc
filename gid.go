package guc

import (
	`bytes`
	`runtime`
	`strconv`
)

func gid() (id uint64) {
	buffer := make([]byte, 64)
	buffer = buffer[:runtime.Stack(buffer, false)]
	buffer = bytes.TrimPrefix(buffer, []byte("goroutine "))
	buffer = buffer[:bytes.IndexByte(buffer, ' ')]
	id, _ = strconv.ParseUint(string(buffer), 10, 64)

	return
}
