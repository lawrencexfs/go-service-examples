package skill

import (
	b3 "battleservice/src/services/base/behavior3go"
	b3config "battleservice/src/services/base/behavior3go/config"
	b3core "battleservice/src/services/base/behavior3go/core"
	"battleservice/src/services/battle/conf"
	"battleservice/src/services/battle/scene/bll"
	"battleservice/src/services/battle/scene/plr"
	"battleservice/src/services/battle/usercmd"

	_ "github.com/cihub/seelog"
	"github.com/giant-tech/go-service/base/linmath"
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
	posx, _, posz := player.GetPos()
	radius := float64(conf.ConfigMgr_GetMe().GetFoodSize(scene.GetEntityID(), this.ball_type))

	angleVel := &linmath.Vector3{}
	usedefault := true
	targetId := tick.Blackboard.GetUInt32("skillTargetId", "", "")
	if 0 != targetId {
		tball := player.FindViewPlayer(uint64(targetId))
		if tball != nil {
			x, _, z := tball.GetPos()
			angleVel.X = float32(x - posx)
			angleVel.Z = float32(z - posz)
			angleVel.Normalize()
			usedefault = false
		}
	}
	if usedefault {
		angleVel.X = float32(player.GetAngleVel().X)
		angleVel.Z = float32(player.GetAngleVel().Z)
	}

	pos := linmath.Vector3{X: float32(posx), Z: float32(posz)}
	pos.Add(angleVel.Mul(float32(player.GetRadius() + float32(radius))))

	initData := &bll.SkillInitData{
		BallType: usercmd.BallType(this.ball_type),
		ID:       ballid,
		Pos:      pos,
		Radius:   float32(radius),
		Player:   player,
	}

	ballEntity, err := scene.CreateEntity("BallSkill", scene.GetEntityID(), initData, true, 0)
	if err != nil {
		return b3.FAILURE
	}

	newBall := ballEntity.(*bll.BallSkill)

	angleVel.MulS(this.speed)
	newBall.SetSpeed(angleVel)
	//	if newBall.PhysicObj != nil {
	//		newBall.PhysicObj.SetVelocity(angleVel)
	//		newBall.PhysicObj.SetCollidable(false)
	//	}
	newBall.Skill.CastSkill(this.ball_skill)

	return b3.SUCCESS
}
