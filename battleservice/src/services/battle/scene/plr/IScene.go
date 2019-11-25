package plr

import (
	"battleservice/src/services/base/ape"
	"battleservice/src/services/battle/scene/bll"

	"github.com/giant-tech/go-service/base/net/inet"
	"github.com/giant-tech/go-service/framework/iserver"
)

type IScene interface {
	iserver.IEntities
	GenBallID() uint64 // 生成一个场景内唯一的BallID

	RemovePlayerPhysic(player ape.IAbstractParticle)
	Frame() uint32
	GetPlayer(playerID uint64) *ScenePlayer
	TravsalPlayers(f func(*ScenePlayer))
	SceneSize() float64
	GetEntityID() uint64
	RemoveFeed(feed *bll.BallFeed)

	BroadcastMsg(msg inet.IMsg)
	BroadcastMsgExcept(msg inet.IMsg, uid uint64)
}
