package bll

import (
	"battleservice/src/services/base/ape"
	"battleservice/src/services/battle/scene/interfaces"

	"github.com/giant-tech/go-service/framework/iserver"
)

type IScene interface {
	iserver.IEntityGroup
	//GetEntityID() uint64
	AddBall(ball interfaces.IBall)
	AddFeedPhysic(feed ape.IAbstractParticle)
	AddPlayerPhysic(player ape.IAbstractParticle)
	GetRandPos() (x, y float64)
	SceneSize() float64
	UpdateSkillBallCell(ball *BallSkill, oldCellID int)
}
