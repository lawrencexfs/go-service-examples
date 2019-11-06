// 包 plr 处理玩家类相关功能。
package plr

import "battleservice/src/services/battle/usercmd"

func (s *ScenePlayer) RPCMsgMove(msg *usercmd.MsgMove) {
	s.OnNetMove(msg)
}

func (s *ScenePlayer) RPCMsgRun(msg *usercmd.MsgRun) {
	s.OnRun(msg)
}

func (s *ScenePlayer) RPCMsgRelife(msg *usercmd.MsgRelife) {
	s.OnNetReLife(msg)
}

func (s *ScenePlayer) RPCMsgSceneChat(msg *usercmd.MsgSceneChat) {
	s.OnSceneChat(msg)
}

func (s *ScenePlayer) RPCMsgCastSkill(msg *usercmd.MsgCastSkill) {
	s.OnCastSkill(msg)
}
