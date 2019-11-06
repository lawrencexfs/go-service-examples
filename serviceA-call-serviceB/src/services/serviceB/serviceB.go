package serviceB

import (
	log "github.com/cihub/seelog"
	"github.com/giant-tech/go-service/framework/idata"
	"github.com/giant-tech/go-service/framework/service"
)

// serviceB 服务
type ServiceB struct {
	service.BaseService
}

// OnInit 初始化
func (srv *ServiceB) OnInit() error {
	log.Debug("ServiceBService.OnInit")

	//设置服务的额外属性
	srv.SetMetadata("OS", "Linux")
	srv.SetMetadata("Version", "1.0.2")

	srv.RegRPCMsg(&ServiceBRPCProc{srv: srv})

	return nil
}

// OnTick tick
func (srv *ServiceB) OnTick() {

}

// OnDestroy destroy
func (srv *ServiceB) OnDestroy() {
	log.Debug("ServiceBService.OnDestroy")
}

// OnDisconnected 服务断开链接
func (srv *ServiceB) OnDisconnected(infovec []*idata.ServiceInfo) {

	for _, s := range infovec {
		log.Info("ServiceB OnDisconnected, infovec = ", *s)
	}

}

// OnConnected connected
func (srv *ServiceB) OnConnected(infovec []*idata.ServiceInfo) {

	for _, s := range infovec {

		log.Info("ServiceB OnConnected, infovec = ", *s)
	}
}
