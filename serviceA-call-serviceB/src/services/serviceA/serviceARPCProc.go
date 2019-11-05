package serviceA

import (
	log "github.com/cihub/seelog"
)

// ServiceARPCProc serviceA RPC proc
type ServiceARPCProc struct {
	svc *ServiceA
}

// RPCHello rpc hello
func (p *ServiceARPCProc) RPCHello(str string) {
	log.Debug("RPCHello: ", str)
}
