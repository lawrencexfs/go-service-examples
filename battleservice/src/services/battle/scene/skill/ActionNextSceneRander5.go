package skill

import (
	b3 "battleservice/src/services/base/behavior3go"
	b3config "battleservice/src/services/base/behavior3go/config"
	b3core "battleservice/src/services/base/behavior3go/core"
	_ "github.com/cihub/seelog"
	"math"
	"battleservice/src/services/battle/scene/consts"
	"battleservice/src/services/battle/scene/plr"
)

type ActionNextSceneRander5 struct {
	b3core.Action
	n         uint32
	canAttack uint64
}

func (this *ActionNextSceneRander5) Initialize(setting *b3config.BTNodeCfg) {
	this.Action.Initialize(setting)
	this.n = uint32(setting.GetPropertyAsInt("n"))
	this.canAttack = uint64(setting.GetPropertyAsInt("canAttack"))
}

func (this *ActionNextSceneRander5) OnOpen(tick *b3core.Tick) {
	player := tick.Blackboard.Get("player", "", "").(*plr.ScenePlayer)
	endframe := GetEndFrame(player, this.n)
	tick.Blackboard.Set("endframe", uint64(endframe), tick.GetTree().GetID(), this.GetID())

}

func (this *ActionNextSceneRander5) OnTick(tick *b3core.Tick) b3.Status {
	player := tick.Blackboard.Get("player", "", "").(*plr.ScenePlayer)
	endframe := tick.Blackboard.Get("endframe", tick.GetTree().GetID(), this.GetID()).(uint64)
	if endframe <= uint64(player.Frame()) {
		return b3.SUCCESS
	} else {
		if this.canAttack != 0 {
			attackType := tick.Blackboard.GetInt("attackType", "", "")
			hits := tick.Blackboard.Get("hits", "", "").(map[uint32]int)
			if len(hits) == 0 {
				SkillAttack(tick, player, attackType)
			}
		}
		return b3.RUNNING
	}
}

func GetEndFrame(player *plr.ScenePlayer, n uint32) uint32 {
	return uint32(math.Floor(float64(player.Frame()-1)/float64(consts.FrameCountBy100MS)))*consts.FrameCountBy100MS + consts.FrameCountBy100MS*n + 1
}
