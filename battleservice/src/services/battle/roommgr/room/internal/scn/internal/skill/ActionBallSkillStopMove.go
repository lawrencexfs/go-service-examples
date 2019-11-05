package skill

import (
	b3 "battleservice/src/services/base/behavior3go"
	b3core "battleservice/src/services/base/behavior3go/core"
	//	bmath "battleservice/src/services/base/math"
)

type ActionBallSkillStopMove struct {
	b3core.Action
}

func (this *ActionBallSkillStopMove) OnTick(tick *b3core.Tick) b3.Status {
	ballskill := tick.Blackboard.Get("ballskill", "", "").(*SkillBall).ball
	ballskill.GetSpeed().ScaleBy(0)
	//	if ballskill.PhysicObj != nil {
	//		ballskill.PhysicObj.SetVelocity(&bmath.Vector2{0, 0})
	//	}
	return b3.SUCCESS
}
