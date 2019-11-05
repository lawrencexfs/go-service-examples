package user

import (
	"github.com/cihub/seelog"
	"github.com/giant-tech/go-service/base/net/inet"
	"github.com/giant-tech/go-service/logic/gatewaybase/igateway"
	"github.com/giant-tech/go-service/logic/gatewaybase/userbase"
)

// ServiceAUser 玩家
type ServiceAUser struct {
	userbase.GateUserBase
}

// OnUserInit 初始化
func (lu *ServiceAUser) OnUserInit() error {
	seelog.Debug("ServiceAUser.OnUserInit, entityID: ", lu.GetEntityID())

	return nil
}

// OnUserTick 每帧调用
func (lu *ServiceAUser) OnUserTick() {

}

// OnUserFini 析构
func (lu *ServiceAUser) OnUserFini() {
	seelog.Debug("ServiceAUser.OnUserFini, dbid: ", lu.GetEntityID())
}

// OnReconnect 断线重连处理
func (lu *ServiceAUser) OnReconnect(sess inet.ISession) *igateway.ReconnectData {
	//踢人
	seelog.Debug("OnReconnect, UID: ", lu.GetEntityID())

	//暂时没有断线重连，直接踢人
	lu.Logout()

	return &igateway.ReconnectData{IsCreateEntity: true}
}

// OnClose 网络连接断开
func (lu *ServiceAUser) OnClose() {
	//直接踢掉
	lu.SetClientSess(nil)
	lu.Logout()

	//关闭逻辑协程
	lu.CloseRoutine()
}
