package actions

import (
	b3 "battleservice/src/services/base/behavior3go"
	. "battleservice/src/services/base/behavior3go/core"
)

type Runner struct {
	Action
}

func (this *Runner) OnTick(tick *Tick) b3.Status {
	return b3.RUNNING
}
