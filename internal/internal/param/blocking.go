package param

type Blocking struct {
	Capacity int
}

func NewBlocking() *Blocking {
	return &Blocking{
		Capacity: 1024,
	}
}
