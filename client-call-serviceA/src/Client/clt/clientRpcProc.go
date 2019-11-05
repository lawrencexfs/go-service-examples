package clt

import (
	log "github.com/cihub/seelog"
)

// ServiceAUserRpcProc LobbyUser的消息处理函数
type ServiceAUserRpcProc struct {
	Cli *Client
}

func (p *ServiceAUserRpcProc) RPCHello(name string, id uint32) {
	log.Debug("RPCHello, name: ", name, ", id: ", id)
}

func (p *ServiceAUserRpcProc) RPCTestRoom(name string, id uint32) {
	log.Debug("RPCTestRoom, name: ", name, ", id: ", id)
}
