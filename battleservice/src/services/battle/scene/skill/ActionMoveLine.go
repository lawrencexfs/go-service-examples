package skill

import (
	b3 "battleservice/src/services/base/behavior3go"
	b3config "battleservice/src/services/base/behavior3go/config"
	b3core "battleservice/src/services/base/behavior3go/core"
	"battleservice/src/services/battle/scene/plr"

	_ "github.com/cihub/seelog"
	"github.com/giant-tech/go-service/base/linmath"
)

type ActionMoveLine struct {
	b3core.Action
	d        float64
	n        uint64
	dir_type int
}

func (this *ActionMoveLine) Initialize(setting *b3config.BTNodeCfg) {
	this.Action.Initialize(setting)
	this.d = setting.GetProperty("d")
	this.n = uint64(setting.GetPropertyAsInt("n"))
	this.dir_type = setting.GetPropertyAsInt("dir_type")
}

func (this *ActionMoveLine) OnOpen(tick *b3core.Tick) {
	player := tick.Blackboard.Get("player", "", "").(*plr.ScenePlayer)

	var speed linmath.Vector3

	if this.dir_type == 2 {
		x0, y0, z0 := player.GetPos()
		x1 := tick.Blackboard.GetFloat64("source_pos_x", "", "")
		y1 := tick.Blackboard.GetFloat64("source_pos_y", "", "")

		v := &linmath.Vector3{x0, y0, z0}
		hv := v.Sub(linmath.Vector3{float32(x1), 0, float32(y1)})
		hv.Normalize()
		speed = hv
	} else if this.dir_type == 1 {
		speed = linmath.Vector3{player.GetAngleVel().X, 0, player.GetAngleVel().Y}
	} else {
		panic("error dir_type!")
	}

	player.ClearForce()

	force1 := speed
	force1.Mul(float32(this.d / float64(this.n) * 2))
	player.AddForce(force1, this.n)
}

func (this *ActionMoveLine) OnTick(tick *b3core.Tick) b3.Status {
	player := tick.Blackboard.Get("player", "", "").(*plr.ScenePlayer)
	if player.HasForce() == true {
		return b3.RUNNING
	}
	player.ClearForce()
	return b3.SUCCESS
}
