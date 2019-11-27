package types

import (
	"golang.org/x/time/rate"
)

// ISession 代表一个网络连接
type ISession interface {
	RegMsgProcFunc(msgID MsgID, procFunc func(IMsg))
	SetMsg2ID(IMsg2ID)

	Send(IMsg)
	SendRaw([]byte)
	EncodeMsg(IMsg) ([]byte, error)

	Start()
	Close()
	IsClosed() bool

	RemoteAddr() string

	ResetHb()

	SetVerifyMsgID(verifyMsgID MsgID)
	SetVerified()

	SetOnClosed(func())

	SetBytePerSecLimiter(r rate.Limit, b int)
	SetQueryPerSecLimiter(r rate.Limit, b int)
}
