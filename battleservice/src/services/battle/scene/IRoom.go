package scene

type _IRoom interface {
	GetEntityID() uint64
	EndTime() int64
	Frame() uint32
}
