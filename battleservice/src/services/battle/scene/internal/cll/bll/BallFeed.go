package bll

// 蘑菇球

import (
	"battleservice/src/services/base/ape"
	"battleservice/src/services/base/util"
	"battleservice/src/services/battle/conf"
)

type BallFeed struct {
	BallMove
}

func NewBallFeed(scene IScene, typeId uint16, id uint32, x, y float64) *BallFeed {
	radius := float64(conf.ConfigMgr_GetMe().GetFoodSize(scene.GetEntityID(), typeId))
	ballType := conf.ConfigMgr_GetMe().GetFoodBallType(scene.GetEntityID(), typeId)
	ball := &BallFeed{
		BallMove: BallMove{
			BallFood: BallFood{
				id:       id,
				typeID:   typeId,
				BallType: ballType,
				Pos:      util.Vector2{float64(x), float64(y)},
				radius:   float64(radius),
			},
			PhysicObj: ape.NewCircleParticle(float32(x), float32(y), float32(radius)),
		},
	}
	ball.ResetRect()
	ball.PhysicObj.SetFixed(true)
	scene.AddBall(ball)
	scene.AddFeedPhysic(ball.PhysicObj)
	return ball
}
