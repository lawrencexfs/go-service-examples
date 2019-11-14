package usercmd

// 给客户端提供的消息

type MsgTypeCmd int32

const (
	MsgTypeCmd_Login          MsgTypeCmd = 1001
	MsgTypeCmd_Top            MsgTypeCmd = 1002
	MsgTypeCmd_AddPlayer      MsgTypeCmd = 1003
	MsgTypeCmd_RemovePlayer   MsgTypeCmd = 1004
	MsgTypeCmd_Move           MsgTypeCmd = 1006
	MsgTypeCmd_Run            MsgTypeCmd = 1007
	MsgTypeCmd_ReLife         MsgTypeCmd = 1009
	MsgTypeCmd_Death          MsgTypeCmd = 1010
	MsgTypeCmd_EndRoom        MsgTypeCmd = 1011
	MsgTypeCmd_RefreshPlayer  MsgTypeCmd = 1013
	MsgTypeCmd_HeartBeat      MsgTypeCmd = 1016
	MsgTypeCmd_SceneChat      MsgTypeCmd = 1020
	MsgTypeCmd_ActCloseSocket MsgTypeCmd = 1021
	MsgTypeCmd_ErrorMsg       MsgTypeCmd = 1025
	MsgTypeCmd_SceneTCP       MsgTypeCmd = 1031
	MsgTypeCmd_SceneUDP       MsgTypeCmd = 1032
	MsgTypeCmd_CastSkill      MsgTypeCmd = 1050
)

type MapObjectConfigType int32

const (
	MapObjectConfigType_Empty MapObjectConfigType = 0
	MapObjectConfigType_Block MapObjectConfigType = 1
)

type BallType int32

const (
	BallType_Player      BallType = 1
	BallType_FoodBegin   BallType = 10
	BallType_FoodNormal  BallType = 11
	BallType_FoodHammer  BallType = 12
	BallType_FoodBomb    BallType = 13
	BallType_FoodEnd     BallType = 19
	BallType_FeedBegin   BallType = 20
	BallType_FeedNormal  BallType = 21
	BallType_FeedEnd     BallType = 29
	BallType_SkillBegin  BallType = 30
	BallType_SkillHammer BallType = 31
	BallType_SkillBomb   BallType = 32
	BallType_SkillEnd    BallType = 39
)

// 请求登录
type MsgLogin struct {
	Name string `protobuf:"bytes,1,req,name=name" json:"name"`
}

// 返回登录
type MsgLoginResult struct {
	Ok          bool             `protobuf:"varint,1,req,name=ok" json:"ok"`
	Id          uint64           `protobuf:"varint,2,req,name=id" json:"id"`
	Name        string           `protobuf:"bytes,3,opt,name=name" json:"name"`
	Others      []*MsgPlayer     `protobuf:"bytes,4,rep,name=others" json:"others,omitempty"`
	Frame       uint32           `protobuf:"varint,5,req,name=frame" json:"frame"`
	BallID      uint32           `protobuf:"varint,8,req,name=ballId" json:"ballId"`
	Balls       []*MsgBall       `protobuf:"bytes,9,rep,name=balls" json:"balls,omitempty"`
	Playerballs []*MsgPlayerBall `protobuf:"bytes,10,rep,name=playerballs" json:"playerballs,omitempty"`
	LeftTime    uint32           `protobuf:"varint,14,opt,name=leftTime" json:"leftTime"`
}

// 返回排行榜Top
type MsgTop struct {
	Players []*MsgPlayer `protobuf:"bytes,1,rep,name=players" json:"players,omitempty"`
	EndTime uint32       `protobuf:"varint,3,opt,name=EndTime,json=endTime" json:"EndTime"`
	Rank    uint32       `protobuf:"varint,4,opt,name=Rank,json=rank" json:"Rank"`
	KillNum uint32       `protobuf:"varint,5,opt,name=KillNum,json=killNum" json:"KillNum"`
}

type MsgSceneTCP struct {
	Eats          []*BallEat       `protobuf:"bytes,1,rep,name=eats" json:"eats,omitempty"`
	Adds          []*MsgBall       `protobuf:"bytes,2,rep,name=adds" json:"adds,omitempty"`
	Removes       []uint32         `protobuf:"varint,3,rep,name=removes" json:"removes,omitempty"`
	Hits          []*HitMsg        `protobuf:"bytes,4,rep,name=hits" json:"hits,omitempty"`
	AddPlayers    []*MsgPlayerBall `protobuf:"bytes,5,rep,name=addPlayers" json:"addPlayers,omitempty"`
	RemovePlayers []uint32         `protobuf:"varint,6,rep,name=removePlayers" json:"removePlayers,omitempty"`
}

type MsgSceneUDP struct {
	Moves []*BallMove `protobuf:"bytes,1,rep,name=moves" json:"moves,omitempty"`
	Frame uint32      `protobuf:"varint,2,req,name=frame" json:"frame"`
}

// 玩家数据
type MsgPlayer struct {
	Id        uint64         `protobuf:"varint,1,req,name=id" json:"id"`
	Name      string         `protobuf:"bytes,2,req,name=name" json:"name"`
	IsLive    bool           `protobuf:"varint,4,opt,name=IsLive,json=isLive" json:"IsLive"`
	SnapInfo  *MsgPlayerSnap `protobuf:"bytes,5,opt,name=SnapInfo,json=snapInfo" json:"SnapInfo,omitempty"`
	BallID    uint32         `protobuf:"varint,6,req,name=ballId" json:"ballId"`
	Curexp    uint32         `protobuf:"varint,7,opt,name=curexp" json:"curexp"`
	Curmp     uint32         `protobuf:"varint,8,opt,name=curmp" json:"curmp"`
	Curhp     uint32         `protobuf:"varint,10,opt,name=curhp" json:"curhp"`
	BombNum   int32          `protobuf:"varint,12,opt,name=bombNum" json:"bombNum"`
	HammerNum int32          `protobuf:"varint,13,opt,name=hammerNum" json:"hammerNum"`
}

// 返回添加玩家 AddPlayer
type MsgAddPlayer struct {
	Player *MsgPlayer `protobuf:"bytes,1,req,name=player" json:"player,omitempty"`
}

// 返回刷新玩家数据
type MsgRefreshPlayer struct {
	Player *MsgPlayer `protobuf:"bytes,1,req,name=player" json:"player,omitempty"`
}

// 返回删除玩家 RemovePlayer
type MsgRemovePlayer struct {
	Id uint64 `protobuf:"varint,1,req,name=id" json:"id"`
}

// 请求移动 Move
type MsgMove struct {
	Angle int32  `protobuf:"varint,1,req,name=angle" json:"angle"`
	Power int32  `protobuf:"varint,2,req,name=power" json:"power"`
	Face  uint32 `protobuf:"varint,3,opt,name=face" json:"face"`
}

// 请求复活 ReLife
type MsgRelife struct {
}

// 返回复活 ReLife
type MsgS2CRelife struct {
	Name     string         `protobuf:"bytes,1,opt,name=name" json:"name"`
	Frame    uint32         `protobuf:"varint,2,opt,name=frame" json:"frame"`
	SnapInfo *MsgPlayerSnap `protobuf:"bytes,3,req,name=SnapInfo,json=snapInfo" json:"SnapInfo,omitempty"`
	Curhp    uint32         `protobuf:"varint,4,opt,name=curhp" json:"curhp"`
	Curmp    uint32         `protobuf:"varint,5,opt,name=curmp" json:"curmp"`
}

// 返回死亡 Death
type MsgDeath struct {
	MaxScore uint32 `protobuf:"varint,1,req,name=maxScore" json:"maxScore"`
	KillId   uint64 `protobuf:"varint,2,req,name=killId" json:"killId"`
	KillName string `protobuf:"bytes,3,req,name=killName" json:"killName"`
	Id       uint64 `protobuf:"varint,4,req,name=Id,json=id" json:"Id"`
}

// 返回结束 EndRoom
type MsgEndRoom struct {
}

// 客户端心跳包
type ClientHeartBeat struct {
}

type MsgActCloseSocket struct {
}

// 释放技能 CastSkill
type MsgCastSkill struct {
	Skillid uint32 `protobuf:"varint,1,opt,name=skillid" json:"skillid"`
}

// 奔跑
type MsgRun struct {
}

// 聊天命令命令
type MsgSceneChat struct {
	Msg string `protobuf:"bytes,1,req,name=Msg,json=msg" json:"Msg"`
	Id  uint64 `protobuf:"varint,2,req,name=Id,json=id" json:"Id"`
}

// 位置同步
type MsgPlayerSnap struct {
	Snapx float32 `protobuf:"fixed32,1,req,name=Snapx,json=snapx" json:"Snapx"`
	Snapy float32 `protobuf:"fixed32,2,req,name=Snapy,json=snapy" json:"Snapy"`
	Angle float32 `protobuf:"fixed32,3,req,name=Angle,json=angle" json:"Angle"`
	Id    uint64  `protobuf:"varint,4,req,name=Id,json=id" json:"Id"`
}

type MsgBall struct {
	Id   uint32 `protobuf:"varint,1,req,name=id" json:"id"`
	Type int32  `protobuf:"varint,2,req,name=type" json:"type"`
	X    int32  `protobuf:"varint,3,req,name=x" json:"x"`
	Y    int32  `protobuf:"varint,4,req,name=y" json:"y"`
}

// 玩家球
type MsgPlayerBall struct {
	Id    uint32 `protobuf:"varint,1,req,name=id" json:"id"`
	Hp    uint32 `protobuf:"varint,3,opt,name=hp" json:"hp"`
	Mp    uint32 `protobuf:"varint,4,opt,name=mp" json:"mp"`
	X     int32  `protobuf:"varint,5,req,name=x" json:"x"`
	Y     int32  `protobuf:"varint,6,req,name=y" json:"y"`
	Angle int32  `protobuf:"varint,7,opt,name=angle" json:"angle"`
	Face  uint32 `protobuf:"varint,8,opt,name=face" json:"face"`
}

// 移动数据
type BallMove struct {
	Id    uint32 `protobuf:"varint,1,req,name=id" json:"id"`
	X     int32  `protobuf:"varint,2,req,name=x" json:"x"`
	Y     int32  `protobuf:"varint,3,req,name=y" json:"y"`
	State uint32 `protobuf:"varint,4,opt,name=state" json:"state"`
	Angle int32  `protobuf:"varint,5,opt,name=angle" json:"angle"`
	Face  uint32 `protobuf:"varint,6,opt,name=face" json:"face"`
}

// 吃球
type BallEat struct {
	Source uint32 `protobuf:"varint,1,req,name=source" json:"source"`
	Target uint32 `protobuf:"varint,2,req,name=target" json:"target"`
}

// 攻击
type HitMsg struct {
	Source uint32 `protobuf:"varint,1,req,name=source" json:"source"`
	Target uint32 `protobuf:"varint,2,req,name=target" json:"target"`
	AddHp  int32  `protobuf:"varint,3,opt,name=addHp" json:"addHp"`
	CurHp  uint32 `protobuf:"varint,4,opt,name=curHp" json:"curHp"`
}
