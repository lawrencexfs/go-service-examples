package interfaces

// 球的基本接口

import (
	"battleservice/src/services/base/util"
	"battleservice/src/services/battle/usercmd"

	"github.com/giant-tech/go-service/base/linmath"
)

type IBall interface {
	GetID() uint64
	SetID(uint64)
	GetTypeId() uint16
	GetBallType() usercmd.BallType
	GetPos() linmath.Vector3
	GetPosPtr() *linmath.Vector3
	SetPos(linmath.Vector3)
	SetBirthPoint(_birthPoint IBirthPoint)
	GetRect() *util.Square
	OnReset()
}
