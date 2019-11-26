package bll

// 球移动 基类

import (
	"battleservice/src/services/base/ape"
	"battleservice/src/services/battle/scene/bll/internal"
	"battleservice/src/services/battle/scene/interfaces"

	"github.com/giant-tech/go-service/base/linmath"
)

type BallMove struct {
	BallFood
	internal.Force
	speed     linmath.Vector3     //速度
	angleVel  linmath.Vector3     //单位速度向量
	PhysicObj *ape.CircleParticle //物理体
}

func (ball *BallMove) GetSpeed() *linmath.Vector3 {
	return &ball.speed
}

func (ball *BallMove) SetSpeed(v *linmath.Vector3) {
	ball.speed = *v
}

func (this *BallMove) GetAngleVel() *linmath.Vector3 {
	return &this.angleVel
}

func (ball *BallMove) SqrMagnitudeTo(target interfaces.IBall) float32 {
	x, y, z := target.GetPos()
	return (ball.Pos.X-x)*(ball.Pos.X-x) + (ball.Pos.Y-y)*(ball.Pos.Y-y) + (ball.Pos.Z-z)*(ball.Pos.Z-z)
}
