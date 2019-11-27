package types

// msg2id.Msg2ID 的取ID接口.
type IMsg2ID interface {
	GetMsgID(msg IMsg) MsgID
}
