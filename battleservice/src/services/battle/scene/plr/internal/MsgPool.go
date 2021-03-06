package internal

import "battleservice/src/services/battle/usercmd"

type MsgPool struct {
	MsgSceneTCP      usercmd.MsgSceneTCP
	MsgSceneUDP      usercmd.MsgSceneUDP
	MsgRelife        usercmd.MsgS2CRelife
	MsgRefreshPlayer usercmd.MsgRefreshPlayer
	MsgDeath         usercmd.MsgDeath
	MsgPlayerSnap    usercmd.MsgPlayerSnap
}

func NewMsgPool() *MsgPool {
	pool := MsgPool{}
	pool.MsgRefreshPlayer.Player = &usercmd.MsgPlayer{}
	return &pool
}
