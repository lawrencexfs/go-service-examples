package scene

import (
	"battleservice/src/services/battle/scene/internal/skill"
	"battleservice/src/services/battle/scene/plr"

	"github.com/cihub/seelog"
)

func init() {
	plr.NewISkillPlayer = skill.NewISkillPlayer
	plr.NewISkillBall = skill.NewISkillBall
}

// LoadSkillBevTree 初始化加载技能行为树.
// 需要在配置加载之后执行，所以不能在 init() 中。
func LoadSkillBevTree() bool {
	if skill.LoadSkillBevTree() {
		return true
	}
	seelog.Error("[启动]LoadSkillBevTree fail! ")
	return false
}
