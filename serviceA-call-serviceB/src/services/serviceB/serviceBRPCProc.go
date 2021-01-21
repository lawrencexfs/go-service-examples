package serviceB

import (
	"serviceA-call-serviceB/src/services/servicetype"

	"github.com/giant-tech/go-service/framework/iserver"

	log "github.com/cihub/seelog"
)

// ServiceBRPCProc serviceB proc
type ServiceBRPCProc struct {
	srv *ServiceB
}

// RPCHello hello
func (p *ServiceBRPCProc) RPCHello(str string) {
	log.Debug("RPCHello: ", str)

	randProxy := iserver.GetServiceProxyMgr().GetRandService(servicetype.ServiceTypeA)
	err := randProxy.AsyncCall("Hello", "hello World from ServiceB")
	if err != nil {
		log.Error("AsyncCall: ", err)
	}
}

// RPCSyncCall synccall
func (p *ServiceBRPCProc) RPCSyncCall() string {
	log.Debug("RPCSyncCall: from ServiceA")

	return "ServiceA call ServiceB success"
}
