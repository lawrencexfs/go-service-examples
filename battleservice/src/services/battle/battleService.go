package battle

import (
	"battleservice/src/services/battle/conf"
	"battleservice/src/services/battle/scene"
	"battleservice/src/services/battle/scene/bll"
	"battleservice/src/services/battle/scene/plr"
	"battleservice/src/services/battle/token"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/cihub/seelog"
	"github.com/giant-tech/go-service/base/net/inet"
	"github.com/giant-tech/go-service/framework/errormsg"
	"github.com/giant-tech/go-service/framework/iserver"
	"github.com/giant-tech/go-service/framework/msgdef"
	"github.com/giant-tech/go-service/logic/gatewaybase"
	"github.com/giant-tech/go-service/logic/gatewaybase/igateway"
)

// BattleService 战斗服务
type BattleService struct {
	gatewaybase.GatewayBase
	token.TokenMgr

	mtx      sync.Mutex
	scene    iserver.ISpace // 测试专用，唯一场景
	playerID uint64         // 测试专用
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
	bs.RegProtoType("Room", &scene.Scene{}, false)
	bs.RegProtoType("Player", &plr.ScenePlayer{}, false)
	bs.RegProtoType("BallFood", &bll.BallFood{}, false)
	bs.RegProtoType("BallFeed", &bll.BallFeed{}, false)
	bs.RegProtoType("BallSkill", &bll.BallSkill{}, false)

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

// OnLoginHandler2 登录处理
func (bs *BattleService) OnLoginHandler2(sess inet.ISession, msg *msgdef.LoginReq) *igateway.LoginRetData {
	//自己有登录方面的处理就放在这里
	seelog.Info("OnLoginHandler, msg.UID = ", msg.UID)

	loginRetData := &igateway.LoginRetData{Msg: &msgdef.LoginResp{}}

	ok, roomID, playerID := bs.LookupToken(msg.Token)
	if !ok {
		seelog.Error("OnLoginHandler got illegal token ", msg.Token)
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

	scene, ok := eroom.(iserver.ISpace)
	if !ok {
		seelog.Error("OnLoginHandler player is not scene")
		loginRetData.Msg.ErrStr = "illegal token"
		loginRetData.Msg.Result = uint32(errormsg.ReturnTypeTOKENINVALID)

		return loginRetData
	}

	//加入场景
	player, err := scene.AddEntity("Player", playerID, nil, true)
	if err != nil {
		seelog.Error("Create scene player failed, err: ", err)

		loginRetData.Msg.Result = uint32(errormsg.ReturnTypeTOKENINVALID)

		return loginRetData
	}

	loginRetData.Entity = player
	loginRetData.Group = scene

	return loginRetData
}

// OnLoginHandler 登录处理，只用于测试
func (bs *BattleService) OnLoginHandler(sess inet.ISession, msg *msgdef.LoginReq) *igateway.LoginRetData {
	//自己有登录方面的处理就放在这里
	seelog.Info("OnLoginHandler, msg.UID = ", msg.UID)

	bs.mtx.Lock()
	defer bs.mtx.Unlock()

	loginRetData := &igateway.LoginRetData{Msg: &msgdef.LoginResp{}}

	if bs.scene == nil {
		scene, err := bs.CreateEntity("Room", 0, "1", true, 0)
		if err != nil {
			seelog.Error("create scene failed: ", err)

			loginRetData.Msg.Result = uint32(errormsg.ReturnTypeTOKENINVALID)
			return loginRetData
		}

		bs.scene = scene.(iserver.ISpace)
	}

	//加入场景
	//playerID递增
	bs.playerID++
	player, err := bs.scene.AddEntity("Player", bs.playerID, nil, true)
	if err != nil {
		seelog.Error("Create scene player failed, err: ", err)

		loginRetData.Msg.Result = uint32(errormsg.ReturnTypeTOKENINVALID)

		return loginRetData
	}

	loginRetData.Entity = player
	loginRetData.Group = bs.scene

	return loginRetData
}
