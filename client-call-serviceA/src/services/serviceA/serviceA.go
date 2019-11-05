package serviceA

import (
	"client-call-serviceA/src/services/serviceA/user"

	log "github.com/cihub/seelog"
	"gitlab.ztgame.com/tech/public/go-service/zeus/framework/idata"
	"gitlab.ztgame.com/tech/public/go-service/zeus/framework/msgdef"
	"gitlab.ztgame.com/tech/public/go-service/zeus/logic/gatewaybase"
)

// ServiceA serviceA
type ServiceA struct {
	gatewaybase.GatewayBase
}

// OnInit 初始化
func (sa *ServiceA) OnInit() error {
	log.Info("ServiceA OnInit")

	var err error
	err = sa.GatewayBase.OnInit(sa)
	if err != nil {
		return err
	}

	sa.RegProtoType("Player", &user.ServiceAUser{}, false)

	return nil
}

// OnTick tick function
func (sa *ServiceA) OnTick() {

}

// OnDestroy 退出时调用
func (sa *ServiceA) OnDestroy() {
	log.Info("ServiceA OnDestroy")

	sa.GatewayBase.OnDestroy()
}

// OnDisConnected 服务断开链接
func (srv *ServiceA) OnDisconnected(infovec []*idata.ServiceInfo) {

	log.Info("ServiceA OnDisconnected, infovec = ", infovec)
	for _, s := range infovec {
		log.Info("ServiceA OnDisconnected, info= ", *s)
	}
}

// OnConnected 连接
func (srv *ServiceA) OnConnected(infovec []*idata.ServiceInfo) {
	log.Info("ServiceA OnConnected, infovec = ", infovec)
}

// OnLoginHandler 登录处理
func (sa *ServiceA) OnLoginHandler(msg *msgdef.LoginReq) msgdef.ReturnType {
	//自己有登录方面的处理就放在这里
	log.Info("ServiceA OnLoginHandler, msg.UID = ", msg.UID)

	return msgdef.ReturnTypeSuccess
}
