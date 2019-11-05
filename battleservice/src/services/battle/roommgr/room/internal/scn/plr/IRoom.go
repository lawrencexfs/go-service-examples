package plr

import (
	"battleservice/src/services/battle/types"
	"gitlab.ztgame.com/tech/public/go-service/zeus/base/net/inet"
)

type IRoom interface {
	ID() types.RoomID
	BroadcastMsg(msg inet.IMsg)
	BroadcastMsgExcept(msg inet.IMsg, uid types.PlayerID)
	Frame() uint32
}
