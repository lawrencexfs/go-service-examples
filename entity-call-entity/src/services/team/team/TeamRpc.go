package team

import (
	log "github.com/cihub/seelog"
)

// RPCHello hello
func (t *Team) RPCHello(name string, id uint32) {
	log.Debug("RPCHello, name: ", name, ", id: ", id)

}
