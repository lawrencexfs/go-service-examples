package interfaces

// 球的基本接口

import (
	"battleservice/src/services/base/util"
	"battleservice/src/services/battle/usercmd"
)

type IBall interface {
	GetID() uint32
	SetID(uint32)
	GetTypeId() uint16
	GetType() usercmd.BallType
	GetPos() (float64, float64)
	SetPos(float64, float64)
	SetBirthPoint(_birthPoint IBirthPoint)
	GetRect() *util.Square
	OnReset()
}
