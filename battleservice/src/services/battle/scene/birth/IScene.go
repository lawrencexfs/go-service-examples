package birth

type IScene interface {
	GetRandPos() (x, y float64)
	GetEntityID() uint64
	GenBallID() uint32
}
