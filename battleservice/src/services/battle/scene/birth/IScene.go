package birth

type IScene interface {
	GetRandPos() (x, z float64)
	GetEntityID() uint64
	GenBallID() uint64
}
