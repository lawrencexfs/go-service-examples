package scene

// 房间类

import (
	"battleservice/src/services/battle/scene/consts"
	"battleservice/src/services/battle/scene/plr"
	"battleservice/src/services/battle/scene/rank"
	"battleservice/src/services/battle/types"
	"battleservice/src/services/battle/usercmd"
	"runtime/debug"
	"time"

	"github.com/cihub/seelog"

	assert "github.com/aurelien-rainone/assertgo"
)

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
	seelog.Infof("[房间] 关闭, ID=%d, players=%d", s.GetEntityID(), len(s.scn.Players))

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
		case act := <-s.actC:
			act()
		case <-s.doneC:
			assert.True(s.IsClosed())
			return
		}
	}
}

// 从房间里删除玩家
func (s *Scene) removePlayerById(playerId types.PlayerID) {
	//退出房间处理
	s.scn.RemovePlayer(playerId)

	// 通知其它人删除玩家
	rmCmd := &usercmd.MsgRemovePlayer{
		Id: uint64(playerId),
	}
	s.BroadcastMsg(rmCmd)

	seelog.Info("[房间] 删除玩家 [", s.GetEntityID(), "] ", playerId, ",", len(s.scn.Players))
}

// 广播结束消息
func (s *Scene) broadcastEndMsg() {
	// 游戏结束，可以对 MsgEndRoom 字段做修改，发送具体游戏结果
	msg := &usercmd.MsgEndRoom{}
	s.BroadcastMsg(msg)
}

// 发送排行榜
func (s *Scene) broadcastTopList() {
	rank.BroadcastTopList(s.endTime, s.scn.Players)
}

// 房间帧定时器
func (s *Scene) render() {
	s.frame++
	s.scn.Render()

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
	if s.endTime < timeNow.Unix() || len(s.scn.Players) == 0 {
		s.close()
		return
	}

	for _, player := range s.scn.Players {
		player.TimeAction(timeNow)
	}
}

// PostAction 发送动作到 ActC. 动作在房间协程中执行。
func (s *Scene) PostAction(action func()) {
	s.actC <- action
}

// Call 发送函数到 ActC 并等待返回值. 函数在房间协程中执行。
func (s *Scene) Call(fn func() interface{}) interface{} {
	ch := make(chan interface{}, 1) // 用于返回结果
	s.actC <- func() {
		ch <- fn()
	}
	return <-ch // 等待直到返回结果
}

func (s *Scene) ID() uint64 {
	return s.GetEntityID()
}

// PostToRemovePlayerById 计划移除玩家，将在房间协程中执行动作。
func (s *Scene) PostToRemovePlayerById(playerID types.PlayerID) {
	s.PostAction(func() {
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

//获取场景玩家
func (s *Scene) GetScenePlayer(id types.PlayerID) *plr.ScenePlayer {
	return s.scn.GetPlayer(id)
}

func (s *Scene) Frame() uint32 {
	return s.frame
}
