package interfaces

// 球的基本接口

import (
	"battleservice/src/services/base/util"
	"battleservice/src/services/battle/usercmd"
)

type IBall interface {
	GetID() uint64
	SetID(uint64)
	GetTypeId() uint16
	GetBallType() usercmd.BallType
	GetPos() (float64, float64)
	SetPos(float64, float64)
	SetBirthPoint(_birthPoint IBirthPoint)
	GetRect() *util.Square
	OnReset()
}
