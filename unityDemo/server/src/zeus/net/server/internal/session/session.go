package session

import (
	"net"
	"zeus/net/internal"
	"zeus/net/internal/types"

	assert "github.com/aurelien-rainone/assertgo"
	"go.uber.org/atomic"
	"golang.org/x/time/rate"
)

var sessionIDGen atomic.Uint64

// Session 封装客户端服务器通用会话，并提供服务器专用功能。
type Session struct {
	_IInternalSession

	id uint64

	// 可以保存任意用户数据
	userData atomic.Value

	onClosedFuncs []func() // 关闭时依次调用
}

// _IInternalSession 代表一个客户端服务器通用会话.
type _IInternalSession interface {
	RegMsgProcFunc(msgID types.MsgID, procFunc func(types.IMsg))
	SetMsg2ID(msg2ID types.IMsg2ID)

	Send(types.IMsg)
	SendRaw([]byte)
	EncodeMsg(types.IMsg) ([]byte, error)

	Start()
	Close()
	IsClosed() bool

	RemoteAddr() string

	ResetHb()

	SetVerifyMsgID(verifyMsgID types.MsgID)
	SetVerified()

	SetOnClosed(func())

	SetBytePerSecLimiter(r rate.Limit, b int)
	SetQueryPerSecLimiter(r rate.Limit, b int)
}

func New(conn net.Conn, encryEnabled bool, msgCreator types.IMsgCreator) *Session {
	assert.True(msgCreator != nil, "msgCreator is nil")
	result := &Session{
		_IInternalSession: internal.NewSession(conn, encryEnabled, msgCreator),

		id: sessionIDGen.Inc(),
	}
	result.SetOnClosed(result.onClosed)
	return result
}

func (s *Session) onClosed() {
	for _, f := range s.onClosedFuncs {
		f()
	}
}

// GetID 获取ID
func (s *Session) GetID() uint64 {
	return s.id
}

func (s *Session) GetUserData() interface{} {
	return s.userData.Load()
}

func (s *Session) SetUserData(data interface{}) {
	s.userData.Store(data)
}

func (s *Session) AddOnClosed(f func()) {
	s.onClosedFuncs = append(s.onClosedFuncs, f)
}
