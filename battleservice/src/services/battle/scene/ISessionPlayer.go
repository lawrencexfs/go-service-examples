package scene

import (
	"battleservice/src/services/battle/scene/plr"
	"battleservice/src/services/battle/types"
)

type ISessionPlayer interface {
	plr.ISender
	Name() string
	PlayerID() types.PlayerID
}
