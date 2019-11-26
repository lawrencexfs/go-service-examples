package bll

// 球信息 到 球网络信息

import (
	"battleservice/src/services/battle/scene/consts"
	"battleservice/src/services/battle/usercmd"
)

func FoodToMsgBall(ball *BallFood) *usercmd.MsgBall {
	return &usercmd.MsgBall{
		Id:   ball.id,
		Type: int32(ball.typeID),
		X:    int32(ball.Pos.X * consts.MsgPosScaleRate),
		Z:    int32(ball.Pos.Z * consts.MsgPosScaleRate),
	}
}

func FeedToMsgBall(ball *BallFeed) *usercmd.MsgBall {
	cmd := &usercmd.MsgBall{
		Id:   ball.id,
		Type: int32(ball.typeID),
		X:    int32(ball.Pos.X * consts.MsgPosScaleRate),
		Z:    int32(ball.Pos.Z * consts.MsgPosScaleRate),
	}
	return cmd
}

func SkillToMsgBall(ball *BallSkill) *usercmd.MsgBall {
	cmd := &usercmd.MsgBall{
		Id:   ball.id,
		Type: int32(ball.BallType),
		X:    int32(ball.Pos.X * consts.MsgPosScaleRate),
		Z:    int32(ball.Pos.Z * consts.MsgPosScaleRate),
	}
	return cmd
}
