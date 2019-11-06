package proc

import (
	"battleservice/src/services/battle/itf"
	"battleservice/src/services/battle/sess"
	"battleservice/src/services/battle/types"

	assert "github.com/aurelien-rainone/assertgo"
	"github.com/cihub/seelog"

	"github.com/giant-tech/go-service/base/net/inet"
	"github.com/giant-tech/go-service/framework/errormsg"
	"github.com/giant-tech/go-service/framework/iserver"
	"github.com/giant-tech/go-service/framework/msgdef"
)

// ProcBattle 是消息处理类(Processor).
type ProcBattle struct {
	sess               inet.ISession       // 一般都需要包含session对象
	itf.IBattleService                     // 本服务
	sessPlayer         *sess.SessionPlayer // 网络会话对应的玩家对象

	// 进入房间才会有
	scene  itf.IScene      // 场景实体
	player iserver.IEntity // 玩家实体
}

// RegisterMsgProcFunctions 克隆自身并注册消息处理函数.
func (p *ProcBattle) RegisterMsgProcFunctions(sess inet.ISession) interface{} {
	assert.True(sess != nil, "session is nil")
	result := &ProcBattle{
		sess:           sess,
		IBattleService: p.IBattleService,
	}

	sess.RegMsgProc(result)
	sess.AddOnClosed(result.OnClosed)

	//seelog.Debugf("ProcBattle.RegisterMsgProcFunctions, p: %p, sess: %p ", result, sess)

	return result
}

// MsgProcLoginReq  MsgProcLoginReq
func (p *ProcBattle) MsgProcLoginReq(msg *msgdef.LoginReq) {
	seelog.Debugf("Begin MsgProcLoginReq, UID: %d, sess: %p", msg.UID, p.sess)

	token := types.Token(msg.Token)

	// 验证
	respMsg := &msgdef.LoginResp{}
	ok, roomID, playerID := p.IBattleService.LookupToken(token)
	if !ok {
		seelog.Error("MsgProcEnterReq got illegal token ", token)
		//respMsg.Error = "illegal token"
		p.sess.Send(respMsg)

		return
	}

	eroom := p.IBattleService.GetEntity(roomID)
	if eroom == nil {
		seelog.Error("MsgProcEnterReq roomID not found: ", roomID)
		//respMsg.Error = "illegal token"
		p.sess.Send(respMsg)

		return
	}

	p.scene, ok = eroom.(itf.IScene)
	if !ok {
		seelog.Error("MsgProcEnterReq player is not scene")
		//respMsg.Error = "illegal token"
		p.sess.Send(respMsg)

		return
	}

	//创建房间成员
	player, err := p.scene.CreateEntityWithID("Player", playerID, p.scene.GetEntityID(), nil, true, 0)
	if err != nil {
		seelog.Error("Create scene player failed, err: ", err)
		p.sess.Send(respMsg)

		return
	}

	p.player = player

	p.sess.SetVerified()

	//发送送登录验证成功消息
	respMsg.Result = uint32(errormsg.ReturnTypeSUCCESS)
	p.sess.Send(respMsg)

	//seelog.Debugf("Finish MsgProcLoginReq, UID: %d, p: %p, player: %p", msg.UID, p, player)
}

// OnClosed 关闭回调
func (p *ProcBattle) OnClosed() {
	// 会话断开时动作...

	//seelog.Infof("ProcBattle OnClosed: %d %s, p: %p, player: %p", p.sess.GetID(), p.sess.RemoteAddr(), p, p.player)

	seelog.Debugf("ProcBattle OnClosed, sess: %p", p.sess)

	// 会话断开时动作...
	seelog.Infof("Closed %d %s", p.sess.GetID(), p.sess.RemoteAddr())

	if p.scene == nil {
		return
	}

	// 下线从房间删除
	//p.scene.PostToRemovePlayerById(p.sessPlayer.PlayerID())
}

// MsgProcCallMsg CallMsg消息处理
func (p *ProcBattle) MsgProcCallMsg(msg *msgdef.CallMsg) {
	//seelog.Infof("MsgProcCallMsg, Seq:%d, MethodName:%s, stype: %d", msg.Seq, msg.MethodName, msg.SType)

	msg.EntityID = p.player.GetEntityID()
	msg.IsFromClient = true

	//如果是投递本服务并且是多协程的，消息投递给实体协程
	if p.IBattleService.IsMultiThread() {
		if msg.IsSync {
			retData := p.player.PostCallMsgAndWait(msg)

			retMsg := &msgdef.CallRespMsg{}
			retMsg.Seq = msg.Seq
			retMsg.RetData = retData.Ret

			if retData.Err != nil {
				retMsg.ErrString = retData.Err.Error()
			}

			p.sess.Send(retMsg)
		} else {
			err := p.player.PostCallMsg(msg)
			if err != nil {
				seelog.Error("AsyncCall err: ", err)
			}
		}
	} else {
		//消息投递给本服务
		msg.GroupID = p.player.GetGroupID()
		p.postToLocalService(p.IBattleService.GetSID(), msg)
	}
}

func (p *ProcBattle) postToLocalService(srvID uint64, msg *msgdef.CallMsg) error {
	var err error
	localS := iserver.GetLocalServiceMgr().GetLocalService(srvID)
	if localS != nil {
		if msg.IsSync {
			retData := localS.PostCallMsgAndWait(msg)

			retMsg := &msgdef.CallRespMsg{}
			retMsg.Seq = msg.Seq
			retMsg.RetData = retData.Ret

			if retData.Err != nil {
				retMsg.ErrString = retData.Err.Error()
				err = retData.Err
			}

			p.sess.Send(retMsg)
		} else {
			err = localS.PostCallMsg(msg)
			if err != nil {
				seelog.Error("service proxy PostCallMsg err: ", err)
			}
		}
	} else {
		//TODO:
	}

	return err
}

// func (p *ProcBattle) MsgProc_MsgLogin(msg *usercmd.MsgLogin) {
// 	seelog.Infof("[登录] 收到登录请求 %s, %d, %s", p.sess.RemoteAddr(), p.sess.GetID(), msg.Name)
// 	p.sessPlayer.SetName(msg.Name)
// 	matchMaker.AddPlayer(p.sessPlayer)
// }

// func (p *ProcBattle) MsgProc_MsgMove(msg *usercmd.MsgMove) {
// 	if p.CheckPlaying() == false {
// 		return
// 	}

// 	p.scene.PostAction(func() {
// 		p.scenePlayer.OnNetMove(msg)
// 	})
// }

// func (p *ProcBattle) MsgProc_MsgRun(msg *usercmd.MsgRun) {
// 	if p.CheckPlaying() == false {
// 		return
// 	}

// 	p.scene.PostAction(func() {
// 		p.scenePlayer.OnRun(msg)
// 	})
// }

// func (p *ProcBattle) MsgProc_MsgRelife(msg *usercmd.MsgRelife) {
// 	if p.CheckPlaying() == false {
// 		return
// 	}
// 	p.scene.PostAction(func() {
// 		p.scenePlayer.OnNetReLife(msg)
// 	})
// }

// func (p *ProcBattle) MsgProc_MsgSceneChat(msg *usercmd.MsgSceneChat) {
// 	if p.CheckPlaying() == false {
// 		return
// 	}

// 	p.scene.PostAction(func() {
// 		p.scenePlayer.OnSceneChat(msg)
// 	})
// }

// func (p *ProcBattle) MsgProc_MsgCastSkill(msg *usercmd.MsgCastSkill) {
// 	if p.CheckPlaying() == false {
// 		return
// 	}
// 	p.scene.PostAction(func() {
// 		p.scenePlayer.OnCastSkill(msg)
// 	})
// }
