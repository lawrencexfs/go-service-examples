package user

import (
	log "github.com/cihub/seelog"
	"gitlab.ztgame.com/tech/public/go-service/zeus/framework/idata"
)

// RPCHello rpc hello
func (lu *ServiceAUser) RPCHello(name string, id uint32) {
	log.Debug("RPCHello, name: ", name, ", id: ", id)

	err := lu.AsyncCall(idata.ServiceClient, "Hello", name, id)
	if err != nil {
		log.Error("RPCHello err: ", err)
	}
}
