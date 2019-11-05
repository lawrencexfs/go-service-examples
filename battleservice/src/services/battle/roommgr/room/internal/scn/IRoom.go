package scn

import (
	"battleservice/src/services/battle/types"
)

type _IRoom interface {
	ID() types.RoomID
	EndTime() int64
	Frame() uint32
}
