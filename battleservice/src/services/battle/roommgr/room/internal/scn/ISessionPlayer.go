package scn

import (
	"battleservice/src/services/battle/roommgr/room/internal/scn/plr"
	"battleservice/src/services/battle/types"
)

type ISessionPlayer interface {
	plr.ISender
	Name() string
	PlayerID() types.PlayerID
}
