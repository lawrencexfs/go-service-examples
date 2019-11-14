package plr

import (
	"battleservice/src/services/base/ape"
	"battleservice/src/services/base/util"
	"battleservice/src/services/battle/scene/internal/cll"
	"battleservice/src/services/battle/scene/internal/cll/bll"
	"battleservice/src/services/battle/scene/internal/interfaces"

	"github.com/giant-tech/go-service/base/net/inet"
)

type IScene interface {
	GenBallID() uint32 // 生成一个场景内唯一的BallID
	AddBall(ball interfaces.IBall)
	RemoveBall(ball interfaces.IBall)
	RemovePlayerPhysic(player ape.IAbstractParticle)
	Frame() uint32
	GetAreaCells(s *util.Square) (cells []*cll.Cell)
	GetPlayer(playerID uint64) *ScenePlayer
	TravsalPlayers(f func(*ScenePlayer))
	SceneSize() float64
	GetEntityID() uint64
	GetCell(px, py float64) (*cll.Cell, bool)
	RemoveFeed(feed *bll.BallFeed)

	CellNumX() int
	CellNumY() int

	BroadcastMsg(msg inet.IMsg)
	BroadcastMsgExcept(msg inet.IMsg, uid uint64)
}
