package limit

type Atomic interface {
	int8 | int16 | int32 | int64 |
		uint | uint16 | uint32 | uint64
}
