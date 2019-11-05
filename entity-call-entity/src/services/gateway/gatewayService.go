package gateway

import (
	"entity-call-entity/src/services/gateway/user"

	"github.com/cihub/seelog"
	log "github.com/cihub/seelog"
	"gitlab.ztgame.com/tech/public/go-service/zeus/framework/idata"
	"gitlab.ztgame.com/tech/public/go-service/zeus/framework/msgdef"
	"gitlab.ztgame.com/tech/public/go-service/zeus/logic/gatewaybase"
)

// GatewayService gateway
type GatewayService struct {
	gatewaybase.GatewayBase
}

// OnInit 初始化
func (sa *GatewayService) OnInit() error {
	log.Info("GatewayService OnInit")

	var err error
	err = sa.GatewayBase.OnInit(sa)
	if err != nil {
		return err
	}

	sa.RegProtoType("Player", &user.GatewayUser{}, false)

	return nil
}

// OnTick tick function
func (sa *GatewayService) OnTick() {

}

// OnDestroy 退出时调用
func (sa *GatewayService) OnDestroy() {
	log.Info("GatewayService OnDestroy")

	sa.GatewayBase.OnDestroy()
}

// OnDisconnected 服务断开链接
func (sa *GatewayService) OnDisconnected(infovec []*idata.ServiceInfo) {

}

// OnConnected 和别的服务建立链接
func (sa *GatewayService) OnConnected(infovec []*idata.ServiceInfo) {
	log.Info("GatewayService OnConnected, infovec = ", infovec)
}

// OnLoginHandler 登录
func (sa *GatewayService) OnLoginHandler(msg *msgdef.LoginReq) msgdef.ReturnType {
	//自己有登录方面的处理就放在这里
	seelog.Info("GatewayService OnLoginHandler, msg.UID = ", msg.UID)

	return msgdef.ReturnTypeSuccess
}
