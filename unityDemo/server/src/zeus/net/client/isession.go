package client

// XXX 可以删除了。只是有些用户还在用 client.ISession. 应该直接用 *client.Session
type ISession interface {
	RegMsgProcFunc(msgID MsgID, msgCreator func() IMsg, procFunc func(IMsg))

	Send(IMsg)
	SendRaw([]byte)
	EncodeMsg(IMsg) ([]byte, error)

	Start()
	Close()
	IsClosed() bool
	SetOnClosed(func())
}
