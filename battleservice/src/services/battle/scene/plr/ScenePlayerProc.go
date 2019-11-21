package plr

import (
	"battleservice/src/services/battle/usercmd"
)

func (s *ScenePlayer) RPCMove(msg *usercmd.MsgMove) {
	s.OnNetMove(msg)
}

func (s *ScenePlayer) RPCRun(msg *usercmd.MsgRun) {
	s.OnRun(msg)
}

func (s *ScenePlayer) RPCRelife(msg *usercmd.MsgRelife) {
	s.OnNetReLife(msg)
}

func (s *ScenePlayer) RPCSceneChat(msg *usercmd.MsgSceneChat) {
	s.OnSceneChat(msg)
}

func (s *ScenePlayer) RPCCastSkill(msg *usercmd.MsgCastSkill) {
	s.OnCastSkill(msg)
}
