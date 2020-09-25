package user

import (
	"github.com/cihub/seelog"
	"github.com/giant-tech/go-service/framework/net/inet"
	"github.com/giant-tech/go-service/logic/gatewaybase/igateway"
	"github.com/giant-tech/go-service/logic/gatewaybase/userbase"

	"entity-call-entity/src/entitydef"
)

// GatewayUser 玩家
type GatewayUser struct {
	entitydef.PlayerDef
	userbase.GateUserBase
}

// OnInit 初始化
func (gu *GatewayUser) OnInit(interface{}) error {
	seelog.Debug("GatewayUser.OnInit, dbid: ", gu.GetEntityID())

	return nil
}

// OnLoop 每帧调用
func (gu *GatewayUser) OnLoop() {

}

// OnDestroy 析构
func (gu *GatewayUser) OnDestroy() {
	seelog.Debug("GatewayUser.OnDestroy, dbid: ", gu.GetEntityID())
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
