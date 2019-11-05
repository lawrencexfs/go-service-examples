package actions

import (
	b3 "battleservice/src/services/base/behavior3go"
	. "battleservice/src/services/base/behavior3go/core"
)

type Failer struct {
	Action
}

func (this *Failer) OnTick(tick *Tick) b3.Status {
	return b3.FAILURE
}
