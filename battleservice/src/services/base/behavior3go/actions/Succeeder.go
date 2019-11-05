package actions

import (
	b3 "battleservice/src/services/base/behavior3go"
	. "battleservice/src/services/base/behavior3go/core"
)

type Succeeder struct {
	Action
}

func (this *Succeeder) OnTick(tick *Tick) b3.Status {
	return b3.SUCCESS
}
