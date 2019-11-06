package bll

import (
	"battleservice/src/services/base/ape"
	"battleservice/src/services/battle/scene/internal/interfaces"
	"battleservice/src/services/battle/types"
)

type IScene interface {
	SceneID() types.SceneID
	AddBall(ball interfaces.IBall)
	AddFeedPhysic(feed ape.IAbstractParticle)
	AddPlayerPhysic(player ape.IAbstractParticle)
	GetRandPos() (x, y float64)
	SceneSize() float64
	UpdateSkillBallCell(ball *BallSkill, oldCellID int)
}
