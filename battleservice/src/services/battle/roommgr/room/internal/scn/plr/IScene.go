package plr

import (
	"battleservice/src/services/base/ape"
	"battleservice/src/services/base/util"
	"battleservice/src/services/battle/roommgr/room/internal/scn/internal/cll"
	"battleservice/src/services/battle/roommgr/room/internal/scn/internal/cll/bll"
	"battleservice/src/services/battle/roommgr/room/internal/scn/internal/interfaces"
	"battleservice/src/services/battle/types"
	"gitlab.ztgame.com/tech/public/go-service/zeus/base/net/inet"
)

type IScene interface {
	GenBallID() uint32 // 生成一个场景内唯一的BallID
	AddBall(ball interfaces.IBall)
	RemoveBall(ball interfaces.IBall)
	RemovePlayerPhysic(player ape.IAbstractParticle)
	Frame() uint32
	GetAreaCells(s *util.Square) (cells []*cll.Cell)
	GetPlayers() map[types.PlayerID]*ScenePlayer
	SceneSize() float64
	SceneID() types.SceneID
	GetCell(px, py float64) (*cll.Cell, bool)
	RemoveFeed(feed *bll.BallFeed)

	CellNumX() int
	CellNumY() int

	BroadcastMsg(msg inet.IMsg)
	BroadcastMsgExcept(msg inet.IMsg, uid types.PlayerID)
}
