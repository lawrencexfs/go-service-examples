package bll

import (
	"battleservice/src/services/battle/scene/interfaces"
)

type IScenePlayer interface {
	GetBallScene() IScene
	GetID() uint64
	GetPower() float64
	IsRunning() bool
	GetIsLive() bool
	KilledByPlayer(killer IScenePlayer)
	RefreshPlayer()
	UpdateExp(addexp int32)
	NewSkillBall(sb *BallSkill) interfaces.ISkillBall
	Frame() uint32
	GetAngle() float64
	GetFace() uint32
}
