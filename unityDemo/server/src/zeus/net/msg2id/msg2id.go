// msg2id 包注册消息对应的ID, 用于发送消息时，查找消息名对应的ID.
// 生成的代码将注册所有消息。

package msg2id

import (
	"fmt"
	"reflect"
	"sync"
	"zeus/net/internal/types"
)

// 线程安全
type Msg2ID struct {
	// Msg type -> Msg ID
	// 用于发送消息
	// 允许不同消息注册为同一ID, 因为可能是属于不同的服务，接收者不同。
	msgTypeToID *sync.Map
}

func New() *Msg2ID {
	return &Msg2ID{
		msgTypeToID: &sync.Map{},
	}
}

// RegMsg2ID 注册消息的ID.
func (m *Msg2ID) RegMsg2ID(msg types.IMsg, msgID types.MsgID) {
	actual, loaded := m.msgTypeToID.LoadOrStore(reflect.TypeOf(msg), msgID)
	if !loaded {
		return
	}

	actualID := actual.(types.MsgID)
	panic(fmt.Sprintf("try to register message %s(ID=%d) to ID=%d", reflect.TypeOf(msg), actualID, msgID))
}

// GetMsgID 从消息类型获取ID.
func (m *Msg2ID) GetMsgID(msg types.IMsg) types.MsgID {
	ID, ok := m.msgTypeToID.Load(reflect.TypeOf(msg))
	if !ok {
		return 0
	}
	return ID.(types.MsgID)
}
