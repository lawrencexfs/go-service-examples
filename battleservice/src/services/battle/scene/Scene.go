package scene

// 场景类

import (
	"battleservice/src/services/base/ape"
	"battleservice/src/services/base/util"
	"battleservice/src/services/battle/conf"
	"battleservice/src/services/battle/scene/consts"
	"battleservice/src/services/battle/scene/internal"
	"battleservice/src/services/battle/scene/internal/birth"
	"battleservice/src/services/battle/scene/internal/cll"
	"battleservice/src/services/battle/scene/internal/cll/bll"
	"battleservice/src/services/battle/scene/internal/interfaces"
	"battleservice/src/services/battle/scene/internal/physic"
	"battleservice/src/services/battle/scene/plr"
	"battleservice/src/services/battle/scene/rank"
	"battleservice/src/services/battle/types"
	"battleservice/src/services/battle/usercmd"
	"math"
	"math/rand"
	"runtime/debug"
	"time"

	assert "github.com/aurelien-rainone/assertgo"
	"github.com/cihub/seelog"
	"github.com/giant-tech/go-service/base/net/inet"
	"github.com/giant-tech/go-service/framework/entity"
	"github.com/giant-tech/go-service/framework/iserver"

	"go.uber.org/atomic"
)

type Scene struct {
	entity.GroupEntity

	endTime int64  // 结束时间
	frame   uint32 // 当前帧数

	doneC chan bool // 用于结束房间协程

	isClosed atomic.Bool // 是否关闭标识

	id   types.SceneID
	room _IRoom // 所在房间

	genBallID   uint32             // 用于生成唯一BallID
	birthPoints *birth.BirthPoints // 出生点
	mapConfig   *conf.MapConfig    // 地图配置

	sceneSize float64     // 地图大小（长、宽相等）
	cellNumX  int         // 格子最大X
	cellNumY  int         // 格子最大Y
	cells     []*cll.Cell // 所有格子

	scenePhysic *physic.ScenePhysic // 场景物理
}

func NewScene(sceneID types.SceneID) *Scene {
	return &Scene{
		id:          sceneID,
		birthPoints: &birth.BirthPoints{},
	}
}

//场景初始化
func (s *Scene) Init() {
	seelog.Info("Scene Init")

	s.endTime = time.Now().Unix() + consts.DefaultPlayTime // 10min为一局
	s.doneC = make(chan bool)

	s.mapConfig = conf.GetMapConfigById(s.SceneID())
	s.scenePhysic = physic.NewScenePhysic()

	s.loadMap()
	for i := 0; i < s.cellNumX*s.cellNumY; i++ {
		s.cells = append(s.cells, cll.NewCell(i))
	}
	s.birthPoints.CreateAllBirthPoint(s)
}

func (s *Scene) loadMap() {
	s.sceneSize = s.mapConfig.Size
	s.cellNumX = int(math.Ceil(s.sceneSize / cll.CellWidth))
	s.cellNumY = int(math.Ceil(s.sceneSize / cll.CellHeight))
	s.scenePhysic.CreateBoard(float32(s.mapConfig.Size))
	for _, v := range s.mapConfig.Nodes {
		internal.LoadMapObjectByConfig(v, s.scenePhysic)
	}
}

//5帧更新
func (s *Scene) render5() {
	s.sendRoomMsg()
}

//时间片渲染
func (s *Scene) Render() {
	now := time.Now()
	nowNano := now.UnixNano()
	var d float64 = consts.FrameTime

	frame := s.Frame()
	if frame%2 == 0 {
		s.scenePhysic.Tick()
	}

	s.TravsalPlayers(func(player *plr.ScenePlayer) {
		player.Update(d, nowNano, s)
	})

	for _, cell := range s.cells {
		cell.Render(s, d, nowNano)
	}
	s.birthPoints.RefreshBirthPoint(d, s)

	if frame%consts.FrameCountBy100MS == 0 {
		//100ms, 5帧更新
		s.render5()
	}
}

// 发送消息, 100ms发送一次
func (s *Scene) sendRoomMsg() {

	s.TravsalPlayers(func(player *plr.ScenePlayer) {
		player.SendSceneMsg()
	})

	for _, cell := range s.cells {
		cell.ResetMsg()
	}

	s.TravsalPlayers(func(player *plr.ScenePlayer) {
		player.ResetMsg()
	})

}

//添加球到场景
func (s *Scene) AddBall(ball interfaces.IBall) {
	x, y := ball.GetPos()
	cell, ok := s.GetCell(x, y)
	if ok {
		cell.Add(ball)
	}
}

//删除球球
func (s *Scene) RemoveBall(ball interfaces.IBall) {
	if nil == ball {
		return
	}
	for _, cell := range s.cells {
		cell.Remove(ball.GetID(), ball.GetType())
	}
}

//获取区域内的所有格子
func (s *Scene) GetAreaCells(sqr *util.Square) (cells []*cll.Cell) {
	minX := int(math.Max(math.Floor(sqr.Left/cll.CellWidth), 0))
	maxX := int(math.Min(math.Floor(sqr.Right/cll.CellWidth), float64(s.cellNumX-1)))
	minY := int(math.Max(math.Floor(sqr.Bottom/cll.CellHeight), 0))
	maxY := int(math.Min(math.Floor(sqr.Top/cll.CellHeight), float64(s.cellNumY-1)))
	for i := minY; i <= maxY; i++ {
		for j := minX; j <= maxX; j++ {
			cells = append(cells, s.cells[i*s.cellNumX+j])
		}
	}
	return
}

//根据坐标获取格子
func (s *Scene) GetCell(px, py float64) (*cll.Cell, bool) {
	idxX := int(math.Max(math.Floor(px/cll.CellWidth), 0))
	idxY := int(math.Max(math.Floor(py/cll.CellHeight), 0))
	//seelog.Info("[房间] cells:", px, "---", py, ", " , "-", idxX, "-", idxY)
	if idxX < s.cellNumX && idxY < s.cellNumY {
		return s.cells[idxY*s.cellNumX+idxX], true
	}
	return nil, false
}

//获取场景玩家
func (s *Scene) GetPlayer(playerID types.PlayerID) *plr.ScenePlayer {
	player := s.GetEntity(uint64(playerID))
	if player == nil {
		return nil
	}
	return player.(*plr.ScenePlayer)
}

//AddPlayer 添加玩家到场景玩家
func (s *Scene) AddPlayer(player ISessionPlayer) {
	playerID := player.PlayerID()

	if s.GetEntity(uint64(playerID)) != nil {
		return // 已存在
	}

	playerEntity, err := s.CreateEntityWithID("Player", uint64(playerID), s.GetEntityID(), nil, true, 0)
	if err != nil {
		return
	}

	scenePlayer := playerEntity.(*plr.ScenePlayer)

	scenePlayer.SetExp(0)
	scenePlayer.Sess = player

	// 把玩家发给其它人
	othermsg := &usercmd.MsgAddPlayer{}
	othermsg.Player = &usercmd.MsgPlayer{}
	othermsg.Player.Id = uint64(scenePlayer.ID)
	othermsg.Player.Name = scenePlayer.Name
	othermsg.Player.IsLive = scenePlayer.IsLive
	othermsg.Player.SnapInfo = scenePlayer.GetSnapInfo()
	othermsg.Player.BallId = scenePlayer.SelfBall.GetID()
	othermsg.Player.Curmp = uint32(scenePlayer.SelfBall.GetMP())
	othermsg.Player.Curhp = uint32(scenePlayer.SelfBall.GetHP())

	scenePlayer.UpdateView(s)
	scenePlayer.UpdateViewPlayers(s)
	scenePlayer.ResetMsg()

	// 发送MsgTop消息给玩家(主要是更新EndTime)
	msgTop := &usercmd.MsgTop{}
	ltime := s.room.EndTime() - time.Now().Unix()
	if ltime > 0 {
		msgTop.EndTime = uint32(ltime)
	} else {
		msgTop.EndTime = 0
	}
	scenePlayer.Sess.Send(msgTop)

	// 把当前场景的人、球都发给玩家
	var others []*usercmd.MsgPlayer
	var balls []*usercmd.MsgBall
	var playerballs []*usercmd.MsgPlayerBall

	s.TravsalPlayers(func(player *plr.ScenePlayer) {
		others = append(others, &usercmd.MsgPlayer{
			Id:     uint64(player.ID),
			BallId: player.SelfBall.GetID(),
			Name:   player.Name,
			IsLive: player.IsLive,

			SnapInfo: player.GetSnapInfo(),
			Curhp:    uint32(player.SelfBall.GetHP()),
			Curmp:    uint32(player.SelfBall.GetMP()),
			Curexp:   player.GetExp(),
		})
	})

	// 玩家视野中的所有球，发送给自己
	cells := scenePlayer.LookCells

	scenePlayer.LookFeeds = make(map[uint32]*bll.BallFeed)
	addfeeds, _ := scenePlayer.UpdateVeiwFeeds()
	balls = append(balls, addfeeds...)

	scenePlayer.LookBallSkill = make(map[uint32]*bll.BallSkill)
	adds, _ := scenePlayer.UpdateVeiwBallSkill()
	balls = append(balls, adds...)

	scenePlayer.LookBallFoods = make(map[uint32]*bll.BallFood)
	addfoods, _ := scenePlayer.UpdateVeiwFoods()
	balls = append(balls, addfoods...)

	//自己
	playerballs = append(playerballs, bll.PlayerBallToMsgBall(scenePlayer.SelfBall))
	//周围玩家
	for _, other := range scenePlayer.Others {
		if true == other.IsLive {
			playerballs = append(playerballs, bll.PlayerBallToMsgBall(other.SelfBall))
		}
	}

	msg := &usercmd.MsgLoginResult{}
	msg.Id = uint64(scenePlayer.ID)
	msg.BallId = scenePlayer.SelfBall.GetID()
	msg.Name = scenePlayer.Name
	msg.Ok = true
	msg.Frame = s.Frame()
	msg.Balls = balls
	msg.Playerballs = playerballs
	msg.Others = others
	msg.LeftTime = uint32(s.room.EndTime() - time.Now().Unix())

	scenePlayer.Sess.Send(msg)
	seelog.Info("[登录] 添加玩家成功addplayer [", s.room.GetEntityID(), ",", scenePlayer.Name, "],", scenePlayer.ID, ",ballId:", msg.BallId, ",view:", scenePlayer.GetViewRect(), ",cell:", len(cells),
		",otplayer:", len(others), "so", len(scenePlayer.Others), ",ball:", len(balls))

	s.BroadcastMsg(othermsg)
}

// 删除玩家
func (s *Scene) RemovePlayer(playerId types.PlayerID) bool {
	e := s.GetEntity(uint64(playerId))
	if e == nil {
		return false
	}

	player := e.(*plr.ScenePlayer)

	s.DestroyEntity(uint64(playerId))

	oldstatus := player.IsLive

	s.RemoveBall(player.SelfBall)
	s.scenePhysic.RemovePlayer(player.SelfBall.PhysicObj)

	player.IsLive = false
	player.SetDeadTime(time.Now().Unix())
	player.Dead(nil)
	player.Sess = nil

	seelog.Info("[注销] 删除玩家成功 [", s.room.GetEntityID(), "],", player.ID, " players:", s.EntityCount(), ";", oldstatus, ",exp:", player.GetExp())
	return true
}

func (s *Scene) SceneID() types.SceneID {
	return s.id
}

func (s *Scene) RemoveFeed(feed *bll.BallFeed) {
	if feed.PhysicObj != nil {
		s.scenePhysic.RemoveFeed(feed.PhysicObj)
	}
}

func (s *Scene) AddFeedPhysic(feed ape.IAbstractParticle) {
	s.scenePhysic.AddFeed(feed)
}

func (s *Scene) AddPlayerPhysic(player ape.IAbstractParticle) {
	s.scenePhysic.AddPlayer(player)
}

func (s *Scene) RemovePlayerPhysic(player ape.IAbstractParticle) {
	s.scenePhysic.RemovePlayer(player)
}

// 地图大小（长、宽相等）
func (s *Scene) SceneSize() float64 {
	return s.sceneSize
}

func (s *Scene) UpdateSkillBallCell(ball *bll.BallSkill, oldCellID int) {
	x, y := ball.GetPos()
	newCell, ok := s.GetCell(x, y)
	if !ok || newCell.ID() == oldCellID {
		return
	}
	oldCell := s.cells[oldCellID]
	if oldCell.ID() != oldCellID {
		panic("从Cell ID获取相应Cell算法有误")
	}

	oldCell.Remove(ball.GetID(), ball.GetType())
	newCell.Add(ball)
}

// 格子最大X
func (s *Scene) CellNumX() int {
	return s.cellNumX
}

// 格子最大Y
func (s *Scene) CellNumY() int {
	return s.cellNumY
}

// 广播 msg
func (s *Scene) BroadcastMsg(msg inet.IMsg) {

	s.TravsalPlayers(func(player *plr.ScenePlayer) {
		player.Send(msg)
	})

}

//广播(剔除特定ID)
func (s *Scene) BroadcastMsgExcept(msg inet.IMsg, uid types.PlayerID) {
	s.TravsalPlayers(func(player *plr.ScenePlayer) {
		if player.ID == uid {
			return
		}
		player.Send(msg)
	})

}

func (s *Scene) GetRandPos() (x, y float64) {
	return s.sceneSize * rand.Float64(), s.sceneSize * rand.Float64()
}

func (s *Scene) GenBallID() uint32 {
	s.genBallID++
	return s.genBallID
}

var genRoomID types.RoomID

func Init() bool {
	return LoadSkillBevTree()
}

func generateRoomID() types.RoomID {
	genRoomID++
	return genRoomID
}

// OnInit 初始化
func (s *Scene) OnInit(initData interface{}) error {
	s.GroupEntity.OnGroupInit()

	return nil
}

// OnLoop 每帧调用
func (s *Scene) OnLoop() {
	seelog.Debug("Team.OnLoop")
	s.GroupEntity.OnGroupLoop()
}

// OnDestroy 销毁
func (s *Scene) OnDestroy() {
	seelog.Debug("OnDestroy")
	s.GroupEntity.OnGroupDestroy()
}

// 关闭房间
func (s *Scene) close() {
	if !s.isClosed.CAS(false /*old*/, true /*new*/) {
		return // CompareAndSet()失败说明已经关闭了
	}
	seelog.Infof("[房间] 关闭, ID=%d, players=%d", s.GetEntityID(), s.EntityCount())

	// 更新排行榜
	s.broadcastTopList()
	// 广播房间结束
	s.broadcastEndMsg()

	close(s.doneC) // 结束房间协程
}

func (s *Scene) IsClosed() bool {
	return s.isClosed.Load()
}

//主循环
func (s *Scene) Run() {
	timeTicker := time.NewTicker(time.Millisecond * consts.FrameTimeMS)

	defer func() {
		timeTicker.Stop()
		if err := recover(); err != nil {
			seelog.Error("[异常] 房间线程出错 [", s.GetEntityID(), "] ", err, "\n", string(debug.Stack()))
		}
	}()

	for {
		select {
		case <-timeTicker.C:
			s.render()

		case <-s.doneC:
			assert.True(s.IsClosed())
			return
		}
	}
}

// 从房间里删除玩家
func (s *Scene) removePlayerById(playerId types.PlayerID) {
	//退出房间处理
	s.RemovePlayer(playerId)

	// 通知其它人删除玩家
	rmCmd := &usercmd.MsgRemovePlayer{
		Id: uint64(playerId),
	}
	s.BroadcastMsg(rmCmd)

	seelog.Info("[房间] 删除玩家 [", s.GetEntityID(), "] ", playerId, ",", s.EntityCount())
}

// 广播结束消息
func (s *Scene) broadcastEndMsg() {
	// 游戏结束，可以对 MsgEndRoom 字段做修改，发送具体游戏结果
	msg := &usercmd.MsgEndRoom{}
	s.BroadcastMsg(msg)
}

// 发送排行榜
func (s *Scene) broadcastTopList() {
	rank.BroadcastTopList(s.endTime, s)
}

// 房间帧定时器
func (s *Scene) render() {
	s.frame++
	s.Render()

	//1s
	if s.frame%(consts.FrameCountBy100MS*10) != 0 {
		return
	}
	s.timeAction1s()

	if s.frame%(consts.FrameCountBy100MS*40) == 0 {
		s.broadcastTopList()
	}
}

// 房间定时器 (一秒一次)
func (s *Scene) timeAction1s() {
	timeNow := time.Now()
	if s.endTime < timeNow.Unix() || s.EntityCount() == 0 {
		s.close()
		return
	}

	s.TravsalPlayers(func(player *plr.ScenePlayer) {
		player.TimeAction(timeNow)
	})

}

func (s *Scene) ID() uint64 {
	return s.GetEntityID()
}

// PostToRemovePlayerById 计划移除玩家，将在房间协程中执行动作。
func (s *Scene) PostToRemovePlayerById(playerID types.PlayerID) {
	s.PostFunction(func() {
		s.removePlayerById(playerID)
	})
}

func (s *Scene) EndTime() int64 {
	return s.endTime
}

// 协程安全
func (s *Scene) GetPlayerCount() uint32 {
	return s.GetPlayerCount()
}

func (s *Scene) Frame() uint32 {
	return s.frame
}

func (s *Scene) TravsalPlayers(f func(*plr.ScenePlayer)) {
	s.TravsalEntity("Player", func(e iserver.IEntity) {
		sp := e.(*plr.ScenePlayer)
		f(sp)
	})
}
