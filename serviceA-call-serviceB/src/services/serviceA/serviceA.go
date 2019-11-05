package serviceA

import (
	"serviceA-call-serviceB/src/services/servicetype"

	"github.com/giant-tech/go-service//framework/iserver"
	"github.com/giant-tech/go-service//framework/service"

	log "github.com/cihub/seelog"
	"github.com/giant-tech/go-service//framework/idata"
)

// ServiceA 大厅服务器
type ServiceA struct {
	service.BaseService
}

// OnInit 初始化
func (srv *ServiceA) OnInit() error {
	log.Info("ServiceA OnInit")

	//设置服务的额外属性
	srv.SetMetadata("OS", "Windows")
	srv.SetMetadata("Version", "1.0.1")

	srv.RegRPCMsg(&ServiceARPCProc{svc: srv})

	return nil
}

// osSelector操作系统选择器
func osSelector(proxySlice []iserver.IServiceProxy, data interface{}) iserver.IServiceProxy {
	os := data.(string)
	for _, proxy := range proxySlice {
		if proxy.GetMetadata("OS") == os {
			return proxy
		}
	}

	//没有合适的返回无效的proxy
	return &service.SNilProxy{}
}

// OnTick tick function
func (srv *ServiceA) OnTick() {
	//获取指定类型的随机服务
	randProxy := iserver.GetServiceProxyMgr().GetRandService(servicetype.ServiceTypeB)
	if randProxy.IsValid() {
		err := randProxy.AsyncCall("Hello", "hello World")
		if err != nil {
			log.Error("AsyncCall: ", err)
		}
	}

	//通过自定义函数删选所需的服务
	proxy := iserver.GetServiceProxyMgr().GetServiceByFunc(servicetype.ServiceTypeB, osSelector, "Linux")
	if proxy.IsValid() {
		var result string
		err := randProxy.SyncCall(&result, "SyncCall")
		if err != nil {
			log.Error("SyncCall:", err)
		} else {
			log.Debug("SyncCall result:", result)
		}
	}
}

// OnDestroy 退出时调用
func (srv *ServiceA) OnDestroy() {
	log.Info("ServiceA OnDestroy")
}

// OnDisconnected 服务断开链接
func (srv *ServiceA) OnDisconnected(infovec []*idata.ServiceInfo) {

	log.Info("ServiceA OnDisconnected, infovec = ", infovec)
	for _, s := range infovec {
		log.Info("ServiceA OnDisconnected, info= ", *s)
	}
}

// OnConnected 和别的服务建立链接
func (srv *ServiceA) OnConnected(infovec []*idata.ServiceInfo) {

	for _, s := range infovec {
		log.Info("ServiceA OnConnected, infovec = ", *s)

		randProxy := iserver.GetServiceProxyMgr().GetRandService(servicetype.ServiceTypeB)
		err := randProxy.AsyncCall("Hello", "hello World")
		if err != nil {
			log.Error("AsyncCall: ", err)
		}

		var result string
		err = randProxy.SyncCall(&result, "SyncCall")
		if err != nil {
			log.Error("SyncCall:", err)
		} else {
			log.Debug("SyncCall result:", result)
		}
	}
}
