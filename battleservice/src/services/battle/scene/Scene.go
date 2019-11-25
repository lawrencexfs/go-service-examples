package scene

import (
	"battleservice/src/services/base/ape"
	"battleservice/src/services/battle/conf"
	"battleservice/src/services/battle/scene/birth"
	"battleservice/src/services/battle/scene/bll"
	"battleservice/src/services/battle/scene/consts"
	"battleservice/src/services/battle/scene/physic"
	"battleservice/src/services/battle/scene/plr"
	"battleservice/src/services/battle/scene/rank"
	"battleservice/src/services/battle/usercmd"
	"math/rand"
	"runtime/debug"
	"time"

	assert "github.com/aurelien-rainone/assertgo"
	"github.com/cihub/seelog"
	"github.com/giant-tech/go-service/base/net/inet"
	"github.com/giant-tech/go-service/base/serializer"
	"github.com/giant-tech/go-service/framework/iserver"
	"github.com/giant-tech/go-service/framework/msgdef"
	"github.com/giant-tech/go-service/framework/space"

	"go.uber.org/atomic"
)

// Scene 场景
type Scene struct {
	space.Space

	endTime int64  // 结束时间
	frame   uint32 // 当前帧数

	doneC chan bool // 用于结束房间协程

	isClosed atomic.Bool // 是否关闭标识

	genBallID   uint64             // 用于生成唯一BallID
	birthPoints *birth.BirthPoints // 出生点
	mapConfig   *conf.MapConfig    // 地图配置

	sceneSize float64 // 地图大小（长、宽相等）

	scenePhysic *physic.ScenePhysic // 场景物理
}

// OnInit 场景初始化
func (s *Scene) OnInit(initData interface{}) error {
	seelog.Info("Scene Init")

	s.birthPoints = &birth.BirthPoints{}
	mapName, ok := initData.(string)
	if ok {
		s.SetMap(mapName)
	}

	s.endTime = time.Now().Unix() + consts.DefaultPlayTime // 10min为一局
	s.doneC = make(chan bool)

	s.mapConfig = conf.GetMapConfigById(s.GetEntityID())
	s.scenePhysic = physic.NewScenePhysic()

	s.loadMap()

	s.birthPoints.CreateAllBirthPoint(s)

	return nil
}

// OnLoop 每帧调用
func (s *Scene) OnLoop() {
	seelog.Debug("Scene.OnLoop")
}

// OnDestroy 销毁
func (s *Scene) OnDestroy() {
	seelog.Debug("Scene.OnDestroy")
}

func (s *Scene) loadMap() {
	s.sceneSize = s.mapConfig.Size

	s.scenePhysic.CreateBoard(float32(s.mapConfig.Size))
	for _, v := range s.mapConfig.Nodes {
		LoadMapObjectByConfig(v, s.scenePhysic)
	}
}

//5帧更新
func (s *Scene) render5() {
	s.sendRoomMsg()
}

//时间片渲染
func (s *Scene) Render() {
	frame := s.Frame()
	if frame%2 == 0 {
		s.scenePhysic.Tick()
	}

	s.birthPoints.RefreshBirthPoint(consts.FrameTime, s)

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

}

//获取场景玩家
func (s *Scene) GetPlayer(playerID uint64) *plr.ScenePlayer {
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
	othermsg.Player.Id = uint64(scenePlayer.GetEntityID())
	othermsg.Player.Name = scenePlayer.Name
	othermsg.Player.IsLive = scenePlayer.IsLive
	othermsg.Player.SnapInfo = scenePlayer.GetSnapInfo()
	othermsg.Player.BallID = scenePlayer.GetID()
	othermsg.Player.Curmp = uint32(scenePlayer.GetMP())
	othermsg.Player.Curhp = uint32(scenePlayer.GetHP())

	scenePlayer.UpdateView(s)
	scenePlayer.UpdateViewPlayers(s)
	scenePlayer.ResetMsg()

	// 发送MsgTop消息给玩家(主要是更新EndTime)
	msgTop := &usercmd.MsgTop{}
	ltime := s.EndTime() - time.Now().Unix()
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
			Id:     uint64(player.GetEntityID()),
			BallID: player.GetID(),
			Name:   player.Name,
			IsLive: player.IsLive,

			SnapInfo: player.GetSnapInfo(),
			Curhp:    uint32(player.GetHP()),
			Curmp:    uint32(player.GetMP()),
			Curexp:   player.GetExp(),
		})
	})

	// 玩家视野中的所有球，发送给自己
	cells := scenePlayer.LookCells

	scenePlayer.LookFeeds = make(map[uint64]*bll.BallFeed)
	addfeeds, _ := scenePlayer.UpdateVeiwFeeds()
	balls = append(balls, addfeeds...)

	scenePlayer.LookBallSkill = make(map[uint64]*bll.BallSkill)
	adds, _ := scenePlayer.UpdateVeiwBallSkill()
	balls = append(balls, adds...)

	scenePlayer.LookBallFoods = make(map[uint64]*bll.BallFood)
	addfoods, _ := scenePlayer.UpdateVeiwFoods()
	balls = append(balls, addfoods...)

	//自己
	playerballs = append(playerballs, plr.PlayerBallToMsgBall(scenePlayer))
	//周围玩家
	for _, other := range scenePlayer.Others {
		if true == other.IsLive {
			playerballs = append(playerballs, plr.PlayerBallToMsgBall(other))
		}
	}

	msg := &usercmd.MsgLoginResult{}
	msg.Id = uint64(scenePlayer.GetEntityID())
	msg.BallID = scenePlayer.GetID()
	msg.Name = scenePlayer.Name
	msg.Ok = true
	msg.Frame = s.Frame()
	msg.Balls = balls
	msg.Playerballs = playerballs
	msg.Others = others
	msg.LeftTime = uint32(s.EndTime() - time.Now().Unix())

	scenePlayer.Sess.Send(msg)
	seelog.Info("[登录] 添加玩家成功addplayer [", s.GetEntityID(), ",", scenePlayer.Name, "],", scenePlayer.GetEntityID(), ",ballId:", msg.BallID, ",view:", scenePlayer.GetViewRect(), ",cell:", len(cells),
		",otplayer:", len(others), "so", len(scenePlayer.Others), ",ball:", len(balls))

	s.BroadcastMsg(othermsg)
}

// 删除玩家
func (s *Scene) RemovePlayer(playerId uint64) bool {
	e := s.GetEntity(uint64(playerId))
	if e == nil {
		return false
	}

	player := e.(*plr.ScenePlayer)

	s.DestroyEntity(uint64(playerId))

	oldstatus := player.IsLive

	s.scenePhysic.RemovePlayer(player.PhysicObj)

	player.IsLive = false
	player.SetDeadTime(time.Now().Unix())
	player.Dead(nil)
	player.Sess = nil

	seelog.Info("[注销] 删除玩家成功 [", s.GetEntityID(), "],", player.GetEntityID(), " players:", s.EntityCount(), ";", oldstatus, ",exp:", player.GetExp())
	return true
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

}

// 广播 msg
func (s *Scene) BroadcastMsg(msg inet.IMsg) {

	s.TravsalPlayers(func(player *plr.ScenePlayer) {
		player.Send(msg)
	})

}

func (s *Scene) BroadcastCall(methodName string, args ...interface{}) {
	msg := &msgdef.CallMsg{}
	msg.MethodName = methodName
	msg.Params = serializer.SerializeNew(args...)

	s.TravsalPlayers(func(player *plr.ScenePlayer) {
		player.Send(msg)
	})

}

//广播(剔除特定ID)
func (s *Scene) BroadcastMsgExcept(msg inet.IMsg, uid uint64) {
	s.TravsalPlayers(func(player *plr.ScenePlayer) {
		if player.GetEntityID() == uid {
			return
		}
		player.Send(msg)
	})

}

func (s *Scene) GetRandPos() (x, y float64) {
	return s.sceneSize * rand.Float64(), s.sceneSize * rand.Float64()
}

func (s *Scene) GenBallID() uint64 {
	s.genBallID++
	return s.genBallID
}

func Init() bool {
	return LoadSkillBevTree()
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

// Run 主循环
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
func (s *Scene) removePlayerById(playerId uint64) {
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
func (s *Scene) PostToRemovePlayerById(playerID uint64) {
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
