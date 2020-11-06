package user

import (
	"entity-call-entity/src/services/servicetype"
	"entity-call-entity/src/services/team/teamdata"

	"github.com/giant-tech/go-service/framework/idata"

	"github.com/giant-tech/go-service/framework/iserver"

	log "github.com/cihub/seelog"
	dbservice "github.com/giant-tech/go-service/base/mongodbservice"
	"github.com/globalsign/mgo/bson"

	"entity-call-entity/src/entitydef"

	protoMsg "entity-call-entity/src/pb"
)

// FriendsInfo 好友信息 临时数据结构
type FriendsInfo struct {
	MyFriendsDbid    []uint64 `bson:"MyFriendsDbid"`
	ApplyFriendsDbid []uint64 `bson:"ApplyFriendsDbid"`
}

type Modifystruct struct {
	Index uint32
	Val   uint32
	//MyFriendsDbid []uint64
	Friends map[uint32]uint32
}

// RPCHello hello
func (gu *GatewayUser) RPCHello(name string, id uint32) {
	log.Debug("GatewayUser::RPCHello, name: ", name, ", id: ", id)

	err := gu.AsyncCall(servicetype.ServiceTypeClient, "Hello", name, id)
	if err != nil {
		log.Error("err: = ", err)
	}
}

// RPCModifyAttr 修改rpc attr
func (gu *GatewayUser) RPCModifyAttr(name string, index uint32, level uint32, modifystru *Modifystruct, changebulletreq *protoMsg.ChangeBulletReq) {
	log.Debug("GatewayUser::RPCModifyAttr(send from client), name: ", name, ", index: ", index, " level: ", level, ", modifystru.index = ", modifystru.Index, ", modifysru.val = ", modifystru.Val, ",Friends[1]=", modifystru.Friends[1], ",Friends[2]=", modifystru.Friends[111])

	log.Debug("GatewayUser::RPCModifyAttr, protoMsg changebulletreq.Full= ", changebulletreq.GetFull(), ", protomsg changebulletreq.pos=", changebulletreq.GetPos())
	gu.SetLevel(level)

	selectProps := bson.M{}
	selectProps["Friends"] = 1

	ret := bson.M{}
	//todo: 根据dbtype去存储
	dbservice.MongoDBQueryOneWithSelect("game", "player", bson.M{"dbid": 1}, selectProps, ret)

	friends := gu.GetFriends()

	//friends.MyFriendsName = "yekoufeng"
	gu.SetFriends(friends)

	//modify heros

	var info entitydef.HEROINFO
	info.HeroName = "yekoufeng"
	info.HeroID = 2

	heros := gu.GetHero()
	//(*heros)["yekoufeng"] = info
	//gu.SetHero(heros)

	log.Debug("GatewayUser::RPCModifyAttr, name: ", name, ", index: ", index, " val: ", level, " ,selectProps: ", selectProps, " ,ret: ", ret, " ,friends: ", friends, " ,heros: ", heros)

}

// RPCCreateTeam 创建队伍
func (gu *GatewayUser) RPCCreateTeam(name string) {
	log.Debug("GatewayUser::RPCCreateTeam, name: ", name)

	//通过自定义函数删选所需的服务
	proxy := iserver.GetServiceProxyMgr().GetRandService(servicetype.ServiceTypeTeam)
	if proxy.IsValid() {
		teamData := &teamdata.CreateTeamData{}
		teamData.TeamName = "test"
		teamData.PlayerInfo = &teamdata.TeamPlayerInfo{PlayerID: gu.GetEntityID(), PlayerName: "Jim"}
		proxy.AsyncCall("CreateTeam", teamData)
	}
}

// RPCCreateTeamResult 创建结果队伍
func (gu *GatewayUser) RPCCreateTeamResult(teamID uint64) {
	log.Debug("GatewayUser::RPCCreateTeamResult, teamID: ", teamID)

	gu.AsyncCall(idata.ServiceClient, "CreateTeamResult", teamID)
}
