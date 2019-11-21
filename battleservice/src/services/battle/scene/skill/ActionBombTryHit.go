package skill

import (
	b3 "battleservice/src/services/base/behavior3go"
	b3config "battleservice/src/services/base/behavior3go/config"
	b3core "battleservice/src/services/base/behavior3go/core"
	"battleservice/src/services/battle/scene/plr"

	_ "github.com/cihub/seelog"
)

type ActionBombTryHit struct {
	b3core.Action
	scale  float64
	gethit uint32
}

func (this *ActionBombTryHit) Initialize(setting *b3config.BTNodeCfg) {
	this.Action.Initialize(setting)
	this.scale = setting.GetProperty("scale")
	this.gethit = uint32(setting.GetPropertyAsInt("gethit"))
}

func (this *ActionBombTryHit) OnTick(tick *b3core.Tick) b3.Status {
	ballskill := tick.Blackboard.Get("ballskill", "", "").(*SkillBall).ball
	player := tick.Blackboard.Get("player", "", "").(*plr.ScenePlayer)
	scene := player.GetScene()

	attckRect := ballskill.GetRect()
	attckRect.SetRadius(ballskill.GetRadius() + this.scale)
	cells := scene.GetAreaCells(attckRect)

	scene.TravsalPlayers(func(other *plr.ScenePlayer) {
		if BallSkillAttack(tick, player, ballskill, this.scale, other.SelfBall) {
			x, y := ballskill.GetPos()
			other.Skill.GetHit2(x, y, this.gethit)
		}
	})

	for _, cell := range cells {
		for _, feed := range cell.Feeds {
			BallSkillAttack(tick, player, ballskill, this.scale, feed)
		}
	}

	return b3.SUCCESS
}
