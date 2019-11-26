package bll

import (
	"battleservice/src/services/base/ape"

	"github.com/giant-tech/go-service/framework/iserver"
)

type IScene interface {
	iserver.IEntityGroup
	//GetEntityID() uint64

	AddFeedPhysic(feed ape.IAbstractParticle)
	AddPlayerPhysic(player ape.IAbstractParticle)
	GetRandPos() (x, y float32)
	SceneSize() float32
	UpdateSkillBallCell(ball *BallSkill, oldCellID int)
}
