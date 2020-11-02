package clt

import (
	"entity-call-entity/src/services/servicetype"

	log "github.com/cihub/seelog"
)

// ServiceAUserRpcProc LobbyUser的消息处理函数
type ServiceAUserRpcProc struct {
	Cli *Client
}

type Modifystruct struct {
	Index uint32
	Val   uint32
	//MyFriendsDbid [5]uint64
	Friends map[uint32]uint32
}

// RPCHello rpc hello
func (p *ServiceAUserRpcProc) RPCHello(name string, id uint32) {
	log.Debug("RPCHello, name: ", name, ", id: ", id)

	//往serviceA发送玩家属性改变消息
	var modifys Modifystruct
	modifys.Index = 222
	modifys.Val = 3222
	//modifys.MyFriendsDbid[0] = 111
	//modifys.MyFriendsDbid[1] = 222

	modifys.Friends = make(map[uint32]uint32)
	modifys.Friends[1] = 2
	modifys.Friends[111] = 2111
	p.Cli.AsyncCall(servicetype.ServiceTypeGateway, "ModifyAttr", "change attr", int32(1), int32(2), &modifys)
}

// RPCCreateTeamResult 创建队伍结果
func (p *ServiceAUserRpcProc) RPCCreateTeamResult(teamID uint64) {
	log.Debug("RPCCreateTeamResult, teamID: ", teamID)

	//离开队伍
	//p.Cli.AsyncCall(servicetype.ServiceTypeTeam, "LeaveTeam")
}
