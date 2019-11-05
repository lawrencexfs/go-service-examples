package user

import (
	"github.com/cihub/seelog"
	"gitlab.ztgame.com/tech/public/go-service/zeus/base/net/inet"
	"gitlab.ztgame.com/tech/public/go-service/zeus/logic/gatewaybase/igateway"
	"gitlab.ztgame.com/tech/public/go-service/zeus/logic/gatewaybase/userbase"

	"entity-call-entity/src/entitydef"
)

// GatewayUser 玩家
type GatewayUser struct {
	entitydef.PlayerDef
	userbase.GateUserBase
}

// OnUserInit 初始化
func (gu *GatewayUser) OnUserInit() error {
	seelog.Debug("GatewayUser.OnUserInit, dbid: ", gu.GetEntityID())

	return nil
}

// OnUserTick 每帧调用
func (gu *GatewayUser) OnUserTick() {

}

// OnUserFini 析构
func (gu *GatewayUser) OnUserFini() {
	seelog.Debug("GatewayUser.OnUserFini, dbid: ", gu.GetEntityID())
}

// OnReconnect 断线重连处理
func (gu *GatewayUser) OnReconnect(sess inet.ISession) *igateway.ReconnectData {
	//踢人
	seelog.Debug("OnReconnect, UID: ", gu.GetEntityID())

	//暂时没有断线重连，直接踢人
	gu.Logout()

	return &igateway.ReconnectData{IsCreateEntity: true}
}

// OnClose 网络连接断开
func (gu *GatewayUser) OnClose() {
	//直接踢掉
	gu.SetClientSess(nil)
	gu.Logout()

	//关闭逻辑协程
	gu.CloseRoutine()
}
