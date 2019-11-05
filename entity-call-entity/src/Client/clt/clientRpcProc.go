package clt

import (
	"entity-call-entity/src/services/servicetype"

	log "github.com/cihub/seelog"
)

// ServiceAUserRpcProc LobbyUser的消息处理函数
type ServiceAUserRpcProc struct {
	Cli *Client
}

// RPCHello rpc hello
func (p *ServiceAUserRpcProc) RPCHello(name string, id uint32) {
	log.Debug("RPCHello, name: ", name, ", id: ", id)

	//往serviceA发送玩家属性改变消息
	p.Cli.AsyncCall(servicetype.ServiceTypeGateway, "ModifyAttr", "change attr", int32(1), int32(2))
}

// RPCModifyAttr RPCModifyAttr
func (p *ServiceAUserRpcProc) RPCModifyAttr(name string, id uint32) {
	log.Debug("RPCModifyAttr, name: ", name, ", id: ", id)
}

// RPCCreateTeamResult 创建队伍结果
func (p *ServiceAUserRpcProc) RPCCreateTeamResult(teamID uint64) {
	log.Debug("RPCCreateTeamResult, teamID: ", teamID)

	//离开队伍
	//p.Cli.AsyncCall(servicetype.ServiceTypeTeam, "LeaveTeam")
}
