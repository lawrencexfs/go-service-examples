package plr

import (
	"battleservice/src/services/battle/types"
	"github.com/giant-tech/go-service/base/net/inet"
)

type IRoom interface {
	ID() types.RoomID
	BroadcastMsg(msg inet.IMsg)
	BroadcastMsgExcept(msg inet.IMsg, uid types.PlayerID)
	Frame() uint32
}
