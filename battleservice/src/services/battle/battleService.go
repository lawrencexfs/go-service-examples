package battle

import (
	"battleservice/src/services/battle/conf"
	"battleservice/src/services/battle/itf"
	"battleservice/src/services/battle/scene"
	"battleservice/src/services/battle/scene/plr"
	"battleservice/src/services/battle/token"
	"battleservice/src/services/battle/types"
	"fmt"
	"math/rand"
	"time"

	"github.com/cihub/seelog"
	"github.com/giant-tech/go-service/base/net/inet"
	"github.com/giant-tech/go-service/framework/errormsg"
	"github.com/giant-tech/go-service/framework/msgdef"
	"github.com/giant-tech/go-service/logic/gatewaybase"
	"github.com/giant-tech/go-service/logic/gatewaybase/igateway"
)

// BattleService 战斗服务
type BattleService struct {
	gatewaybase.GatewayBase
	token.TokenMgr
}

// OnInit 初始化
func (bs *BattleService) OnInit() error {
	seelog.Debug("BattleService.OnInit")

	rand.Seed(time.Now().Unix())

	bs.TokenMgr.Init()

	var err error
	err = bs.GatewayBase.OnInit(bs)
	if err != nil {
		return err
	}

	// 注册proto
	bs.RegProtoType("Scene", &scene.Scene{}, false)
	bs.RegProtoType("Player", &plr.ScenePlayer{}, false)

	// 全局配置
	if !conf.ConfigMgr_GetMe().Init() {
		seelog.Error("[启动] 读取全局配置失败")
		return fmt.Errorf("conf init failed")
	}

	if !scene.Init() {
		return fmt.Errorf("scene.Init failed")
	}

	if !conf.InitMapConfig() {
		seelog.Error("[启动]InitMapConfig fail! ")
		return fmt.Errorf("InitMapConfig failed")
	}

	seelog.Info("[启动] 完成初始化")

	return nil
}

// OnTick tick
func (bs *BattleService) OnTick() {

}

// OnDestroy 析构
func (bs *BattleService) OnDestroy() {
	seelog.Debug("BattleService.OnDestroy")

	bs.GatewayBase.OnDestroy()
}

// OnLoginHandler 登录处理
func (bs *BattleService) OnLoginHandler(sess inet.ISession, msg *msgdef.LoginReq) *igateway.LoginRetData {
	//自己有登录方面的处理就放在这里
	seelog.Info("OnLoginHandler, msg.UID = ", msg.UID)

	loginRetData := &igateway.LoginRetData{Msg: &msgdef.LoginResp{}}

	token := types.Token(msg.Token)
	ok, roomID, playerID := bs.LookupToken(token)
	if !ok {
		seelog.Error("OnLoginHandler got illegal token ", token)
		loginRetData.Msg.ErrStr = "illegal token"
		loginRetData.Msg.Result = uint32(errormsg.ReturnTypeTOKENINVALID)

		return loginRetData
	}

	eroom := bs.GetEntity(roomID)
	if eroom == nil {
		seelog.Error("OnLoginHandler roomID not found: ", roomID)
		loginRetData.Msg.ErrStr = "illegal token"
		loginRetData.Msg.Result = uint32(errormsg.ReturnTypeTOKENINVALID)

		return loginRetData
	}

	scene, ok := eroom.(itf.IScene)
	if !ok {
		seelog.Error("OnLoginHandler player is not scene")
		loginRetData.Msg.ErrStr = "illegal token"
		loginRetData.Msg.Result = uint32(errormsg.ReturnTypeTOKENINVALID)

		return loginRetData
	}

	//TODO: 判断是否已经存在

	//创建房间成员
	player, err := scene.CreateEntityWithID("Player", playerID, scene.GetEntityID(), nil, true, 0)
	if err != nil {
		seelog.Error("Create scene player failed, err: ", err)

		loginRetData.Msg.Result = uint32(errormsg.ReturnTypeTOKENINVALID)

		return loginRetData
	}

	loginRetData.Entity = player
	loginRetData.Group = scene

	return loginRetData
}
