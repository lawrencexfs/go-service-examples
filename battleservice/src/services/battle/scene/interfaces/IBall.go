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
	GetPos() (float32, float32, float32)
	SetPos(float32, float32, float32)
	SetBirthPoint(_birthPoint IBirthPoint)
	GetRect() *util.Square
	OnReset()
}
