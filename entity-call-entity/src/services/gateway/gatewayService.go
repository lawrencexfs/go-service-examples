package gateway

import (
	"entity-call-entity/src/services/gateway/user"

	"github.com/cihub/seelog"
	log "github.com/cihub/seelog"
	"github.com/giant-tech/go-service/base/net/inet"
	"github.com/giant-tech/go-service/framework/errormsg"
	"github.com/giant-tech/go-service/framework/idata"
	"github.com/giant-tech/go-service/framework/iserver"
	"github.com/giant-tech/go-service/framework/msgdef"
	"github.com/giant-tech/go-service/logic/gatewaybase"
	"github.com/giant-tech/go-service/logic/gatewaybase/igateway"
	"github.com/giant-tech/go-service/logic/gatewaybase/userbase"
)

// GatewayService gateway
type GatewayService struct {
	gatewaybase.GatewayBase
}

// OnInit 初始化
func (gs *GatewayService) OnInit() error {
	log.Info("GatewayService OnInit")

	var err error
	err = gs.GatewayBase.OnInit(gs)
	if err != nil {
		return err
	}

	gs.RegProtoType("Player", &user.GatewayUser{}, false)

	return nil
}

// OnTick tick function
func (gs *GatewayService) OnTick() {

}

// OnDestroy 退出时调用
func (gs *GatewayService) OnDestroy() {
	log.Info("GatewayService OnDestroy")

	gs.GatewayBase.OnDestroy()
}

// OnDisconnected 服务断开链接
func (gs *GatewayService) OnDisconnected(infovec []*idata.ServiceInfo) {

}

// OnConnected 和别的服务建立链接
func (gs *GatewayService) OnConnected(infovec []*idata.ServiceInfo) {
	log.Info("GatewayService OnConnected, infovec = ", infovec)
}

// OnLoginHandler 登录
func (gs *GatewayService) OnLoginHandler(sess inet.ISession, msg *msgdef.LoginReq) *igateway.LoginRetData {
	//自己有登录方面的处理就放在这里
	seelog.Info("GatewayService OnLoginHandler, msg.UID = ", msg.UID)

	loginRetData := &igateway.LoginRetData{Msg: &msgdef.LoginResp{}}

	var entity iserver.IEntity

	oldEntity := gs.GetEntity(msg.UID)
	if oldEntity != nil {
		log.Debugf("GatewayService.OnLoginHandler, OnReconnect, UID: %d", msg.UID)

		ireconnect, ok := oldEntity.(igateway.IReconnectHandler)
		if !ok {
			log.Errorf("GatewayService.OnLoginHandler not IReconnectHandler, UID: %d", msg.UID)
			loginRetData.Msg.Result = uint32(errormsg.ReturnTypeFAILRELOGIN)

			return loginRetData
		}

		ret := oldEntity.PostFunctionAndWait(func() interface{} { return ireconnect.OnReconnect(sess) })
		reData, ok := ret.(*igateway.ReconnectData)
		if !ok {
			log.Errorf("GatewayService.OnLoginHandler OnReconnect failed, UID: %d", msg.UID)

			loginRetData.Msg.Result = uint32(errormsg.ReturnTypeFAILRELOGIN)

			return loginRetData
		}

		if reData.Err != nil {
			log.Error("GatewayService.OnLoginHandler OnReconnect failed, UID: ", msg.UID, ", err: ", reData.Err)

			loginRetData.Msg.Result = uint32(errormsg.ReturnTypeFAILRELOGIN)

			return loginRetData
		}

		if !reData.IsCreateEntity {
			entity = oldEntity
		}
	}

	if entity == nil {
		userBase := &userbase.UserInitData{
			Sess:    sess,
			Version: msg.Version,
		}

		var err error
		// 创建新玩家
		entity, err = gs.CreateEntityWithID("Player", msg.UID, 0, userBase, true, 0)
		if err != nil {
			log.Error("Create user failed, err: ", err, ", UID: ", msg.UID)
			loginRetData.Msg.Result = uint32(errormsg.ReturnTypeFAILRELOGIN)

			return loginRetData
		}
	}

	loginRetData.Entity = entity

	return loginRetData
}
