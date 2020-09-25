package serviceA

import (
	"client-call-serviceA/src/services/serviceA/user"

	log "github.com/cihub/seelog"
	"github.com/giant-tech/go-service/framework/errormsg"
	"github.com/giant-tech/go-service/framework/idata"
	"github.com/giant-tech/go-service/framework/iserver"
	"github.com/giant-tech/go-service/framework/msgdef"
	"github.com/giant-tech/go-service/framework/net/inet"
	"github.com/giant-tech/go-service/logic/gatewaybase"
	"github.com/giant-tech/go-service/logic/gatewaybase/igateway"
	"github.com/giant-tech/go-service/logic/gatewaybase/userbase"
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
func (sa *ServiceA) OnDisconnected(infovec []*idata.ServiceInfo) {

	log.Info("ServiceA OnDisconnected, infovec = ", infovec)
	for _, s := range infovec {
		log.Info("ServiceA OnDisconnected, info= ", *s)
	}
}

// OnConnected 连接
func (sa *ServiceA) OnConnected(infovec []*idata.ServiceInfo) {
	log.Info("ServiceA OnConnected, infovec = ", infovec)
}

// OnLoginHandler 登录处理
func (sa *ServiceA) OnLoginHandler(sess inet.ISession, msg *msgdef.LoginReq) *igateway.LoginRetData {
	//自己有登录方面的处理就放在这里
	log.Info("ServiceA OnLoginHandler, msg.UID = ", msg.UID)

	loginRetData := &igateway.LoginRetData{Msg: &msgdef.LoginResp{}}

	var entity iserver.IEntity

	oldEntity := sa.GetEntity(msg.UID)
	if oldEntity != nil {
		log.Debugf("ServiceA.OnLoginHandler, OnReconnect, UID: %d", msg.UID)

		ireconnect, ok := oldEntity.(igateway.IReconnectHandler)
		if !ok {
			log.Errorf("ServiceA.OnLoginHandler not IReconnectHandler, UID: %d", msg.UID)
			loginRetData.Msg.Result = uint32(errormsg.ReturnTypeFAILRELOGIN)

			return loginRetData
		}

		ret := oldEntity.PostFunctionAndWait(func() interface{} { return ireconnect.OnReconnect(sess) })
		reData, ok := ret.(*igateway.ReconnectData)
		if !ok {
			log.Errorf("ServiceA.OnLoginHandler OnReconnect failed, UID: %d", msg.UID)

			loginRetData.Msg.Result = uint32(errormsg.ReturnTypeFAILRELOGIN)

			return loginRetData
		}

		if reData.Err != nil {
			log.Error("ServiceA.OnLoginHandler OnReconnect failed, UID: ", msg.UID, ", err: ", reData.Err)

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
		entity, err = sa.CreateEntityWithID("Player", msg.UID, 0, userBase, true, 0)
		if err != nil {
			log.Error("Create user failed, err: ", err, ", UID: ", msg.UID)
			loginRetData.Msg.Result = uint32(errormsg.ReturnTypeFAILRELOGIN)

			return loginRetData
		}
	}

	loginRetData.Entity = entity

	return loginRetData
}
