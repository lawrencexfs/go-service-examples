package scene

import (
	"battleservice/src/services/battle/scene/plr"
)

type ISessionPlayer interface {
	plr.ISender
	Name() string
	PlayerID() uint64
}
