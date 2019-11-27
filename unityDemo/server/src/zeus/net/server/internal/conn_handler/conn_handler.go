package conn_handler

import (
	"net"
	"zeus/net/internal/types"
	"zeus/net/server/internal/msg_proc_set"
	"zeus/net/server/internal/session"

	assert "github.com/aurelien-rainone/assertgo"
	"golang.org/x/time/rate"
)

type ConnHandler struct {
	msgCreator types.IMsgCreator // 会话创建时需要
	msg2ID     types.IMsg2ID     // 会话创建后设置到会话
	msgProcSet *msg_proc_set.MsgProcSet

	isSessNeedVerify bool        // 是否会话需要验证
	sessVerifyMsgID  types.MsgID // 验证消息ID

	// 限流
	bpsLimiter *rate.Limiter // 限制每秒接收字节数
	qpsLimiter *rate.Limiter // 限制每秒接收请求数
}

func New(msgCreator types.IMsgCreator, msg2ID types.IMsg2ID) *ConnHandler {
	assert.True(msgCreator != nil)
	assert.True(msg2ID != nil)
	return &ConnHandler{
		msgCreator: msgCreator,
		msg2ID:     msg2ID,
		msgProcSet: msg_proc_set.New(),
	}
}

// 处理连接.
// 创建 Session, 并且开始会话协程。
func (h *ConnHandler) HandleConn(conn net.Conn) {
	var encrypt bool // Todo: enable encrypt
	sess := session.New(conn, encrypt, h.msgCreator)
	sess.SetMsg2ID(h.msg2ID)

	// 在会话上注册所有处理函数
	if h.msgProcSet != nil {
		h.msgProcSet.RegisterToSession(sess)
	}

	if h.isSessNeedVerify {
		sess.SetVerifyMsgID(h.sessVerifyMsgID)
	}

	// 设置限流
	h.setLimitersToSession(sess)

	sess.ResetHb() // Todo: Move it into Start()
	sess.Start()
}

func (h *ConnHandler) AddMsgProc(msgProc msg_proc_set.IMsgProc) {
	h.msgProcSet.AddMsgProc(msgProc)
}

func (h *ConnHandler) HasMsgProc() bool {
	return !h.msgProcSet.IsEmpty()
}

// SetVerifyMsg 设置会话的验证消息.
// 强制会话必须验证，会话的第1个消息将做为验证消息，消息类型必须为输入类型.
func (h *ConnHandler) SetVerifyMsgID(verifyMsgID types.MsgID) {
	h.isSessNeedVerify = true
	h.sessVerifyMsgID = verifyMsgID
}

// SetBytePerSecLimiter 设置每秒接收字节数限制.
// r(rate) 为每秒字节数。
// b(burst) 为峰值字节数。
func (h *ConnHandler) SetBytePerSecLimiter(r rate.Limit, b int) {
	h.bpsLimiter = rate.NewLimiter(r, b)
}

// SetQueryPerSecLimiter 设置每秒接收请求数限制.
// r(rate) 为每秒请求数。
// b(burst) 为峰值请求数。
// 必须在 Start() 之前设置，避免 DataRace.
func (h *ConnHandler) SetQueryPerSecLimiter(r rate.Limit, b int) {
	h.qpsLimiter = rate.NewLimiter(r, b)
}

func (h *ConnHandler) setLimitersToSession(sess *session.Session) {
	bps := h.bpsLimiter
	if bps != nil {
		sess.SetBytePerSecLimiter(bps.Limit(), bps.Burst())
	}

	qps := h.qpsLimiter
	if qps != nil {
		sess.SetQueryPerSecLimiter(qps.Limit(), qps.Burst())
	}
}
