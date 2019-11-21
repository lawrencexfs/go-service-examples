package skill

import (
	b3 "battleservice/src/services/base/behavior3go"
	b3config "battleservice/src/services/base/behavior3go/config"
	b3core "battleservice/src/services/base/behavior3go/core"
	bmath "battleservice/src/services/base/math"
	"battleservice/src/services/base/util"
	"battleservice/src/services/battle/conf"
	"battleservice/src/services/battle/scene/bll"
	"battleservice/src/services/battle/scene/plr"
	"battleservice/src/services/battle/usercmd"

	_ "github.com/cihub/seelog"
)

type ActionThrowBall struct {
	b3core.Action
	ball_type  uint16
	speed      float32
	ball_skill uint32
}

func (this *ActionThrowBall) Initialize(setting *b3config.BTNodeCfg) {
	this.Action.Initialize(setting)
	this.ball_type = uint16(setting.GetPropertyAsInt("ball_type"))
	this.speed = float32(setting.GetProperty("speed"))
	this.ball_skill = uint32(setting.GetPropertyAsInt("ball_skill"))
}

func (this *ActionThrowBall) OnTick(tick *b3core.Tick) b3.Status {
	player := tick.Blackboard.Get("player", "", "").(*plr.ScenePlayer)
	if player.IsLive == false {
		return b3.FAILURE
	}

	scene := player.GetScene()
	ballid := scene.GenBallID()
	posx, posy := player.GetPos()
	radius := float64(conf.ConfigMgr_GetMe().GetFoodSize(scene.GetEntityID(), this.ball_type))

	angleVel := &bmath.Vector2{}
	usedefault := true
	targetId := tick.Blackboard.GetUInt32("skillTargetId", "", "")
	if 0 != targetId {
		tball := player.FindViewPlayer(uint64(targetId))
		if tball != nil {
			x, y := tball.GetPos()
			angleVel.X = float32(x - posx)
			angleVel.Y = float32(y - posy)
			angleVel.NormalizeSelf()
			usedefault = false
		}
	}
	if usedefault {
		angleVel.X = float32(player.GetAngleVel().X)
		angleVel.Y = float32(player.GetAngleVel().Y)
	}

	pos := bmath.Vector2{float32(posx), float32(posy)}
	pos.IncreaseBy(angleVel.Mult(float32(player.GetRadius() + radius)))

	newBall := bll.NewBallSkill(usercmd.BallType(this.ball_type), ballid, float64(pos.X), float64(pos.Y), radius, player)
	newBall.ResetRect()

	scene.AddBall(newBall)
	//	scene.scenePhysic.AddSkill(newBall)

	angleVel.ScaleBy(this.speed)
	newBall.SetSpeed(&util.Vector2{float64(angleVel.X), float64(angleVel.Y)})
	//	if newBall.PhysicObj != nil {
	//		newBall.PhysicObj.SetVelocity(angleVel)
	//		newBall.PhysicObj.SetCollidable(false)
	//	}
	newBall.Skill.CastSkill(this.ball_skill)

	return b3.SUCCESS
}
