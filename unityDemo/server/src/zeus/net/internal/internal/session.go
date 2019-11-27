package internal

import (
	"context"
	"fmt"
	"net"
	"reflect"
	"time"
	"zeus/net/internal/internal/msgenc"
	"zeus/net/internal/internal/msghdl"
	"zeus/net/internal/internal/sflist"
	"zeus/net/internal/types"

	assert "github.com/aurelien-rainone/assertgo"
	log "github.com/cihub/seelog"
	"github.com/spf13/viper"
	"go.uber.org/atomic"
	"golang.org/x/time/rate"
)

func NewSession(conn net.Conn, encryEnabled bool, msgCreator types.IMsgCreator) *session {
	assert.True(msgCreator != nil, "msgCreator is nil")
	sess := &session{
		conn:    conn,
		sendBuf: sflist.NewSafeList(),

		msgCreator: msgCreator,
		msgHdlr:    msghdl.New(msgCreator),

		hbTimerInterval: time.Duration(viper.GetInt("Config.HBTimer")) * time.Second,

		_encryEnabled: encryEnabled,

		bpsLimiter: rate.NewLimiter(rate.Inf, 0),
		qpsLimiter: rate.NewLimiter(rate.Inf, 0),
	}

	sess.hbTimer = time.NewTimer(sess.hbTimerInterval)
	sess.ctx, sess.ctxCancel = context.WithCancel(context.Background())
	return sess
}

// TODO: thread-safe

// session 代表一个网络连接
type session struct {
	conn    net.Conn         // 底层连接对象
	sendBuf *sflist.SafeList // 缓存待发送的数据

	ctx       context.Context
	ctxCancel context.CancelFunc

	msg2ID     types.IMsg2ID
	msgCreator types.IMsgCreator
	msgHdlr    *msghdl.MessageHandler

	hbTimer         *time.Timer
	hbTimerInterval time.Duration

	_isClosed     atomic.Bool
	_isError      atomic.Bool
	_encryEnabled bool

	// 关闭事件处理器
	onClosed func()

	// 会话验证功能，仅服务器有用。可以设置会话需要验证，第1个消息必须为验证消息。
	needVerify  bool        // 是否需要验证
	verifyMsgID types.MsgID // 验证消息ID
	isVerified  atomic.Bool // 是否已通过验证。由消息处理器设置。

	// 限流
	bpsLimiter *rate.Limiter // 限制每秒接收字节数
	qpsLimiter *rate.Limiter // 限制每秒接收请求数
}

// Start 验证完成
func (sess *session) Start() {
	assert.True(sess.msg2ID != nil)
	assert.True(sess.msgCreator != nil)

	go sess.recvLoop() // 协程中会调用消息处理器
	go sess.sendLoop()
	if viper.GetBool("Config.HeartBeat") {
		go sess.hbLoop()
	}
}

// Send 发送消息，立即返回.
func (sess *session) Send(msg types.IMsg) {
	if msg == nil {
		return
	}

	if sess.IsClosed() {
		log.Warnf("Send after sess close %s %s %s",
			sess.conn.RemoteAddr(), reflect.TypeOf(msg),
			fmt.Sprintf("%s", msg)) // log是异步的，所以 msg 必须复制下。
		return
	}

	// 队列已满, 表示客户端处理太慢，断开。
	// XXX
	//	if len(sess.sendBufC) >= cap(sess.sendBufC) {
	//		log.Error("Close slow session: ", sess.conn.RemoteAddr())
	//		sess.Close()
	//		return
	//	}

	msgBuf, err := sess.EncodeMsg(msg)
	if err != nil {
		log.Error("Encode message error in Send(): ", err)
		return
	}

	sess.sendBuf.Put(msgBuf)
}

// Send 发送数据，立即返回.
func (sess *session) SendRaw(buff []byte) {
	if sess.IsClosed() {
		log.Warnf("Send after sess close %s", sess.conn.RemoteAddr())
		return
	}

	sess.sendBuf.Put(buff)
}

// EncodeMsg
func (sess *session) EncodeMsg(msg types.IMsg) ([]byte, error) {
	msgID := sess.msg2ID.GetMsgID(msg)
	if msgID == 0 {
		// 应该用 msg2id.RegMsg2ID()注册才行
		log.Errorf("message '%s' is not registered", reflect.TypeOf(msg))
		return nil, fmt.Errorf("message '%s' is not registered", reflect.TypeOf(msg))
	}

	msgBuf, err := msgenc.EncodeMsg(msg, msgID)
	if err != nil {
		log.Error("Encode message error in EncodeMsg(): ", err)
		return nil, err
	}

	return msgBuf, nil
}

// Touch 记录心跳状态
func (sess *session) ResetHb() {
	sess.hbTimer.Reset(sess.hbTimerInterval)
}

func (sess *session) hbLoop() {
	for {
		select {
		case <-sess.ctx.Done():
			sess.hbTimer.Stop()
			return
		case <-sess.hbTimer.C:
			log.Error("sess heart tick expired ", sess.conn.RemoteAddr())

			sess._isError.Store(true)
			sess.hbTimer.Stop()
			sess.Close()
			return
		}
	}
}

func (sess *session) recvLoop() {
	var err error
	assert.True(err == nil, "init error is not nil")

	for {
		select {
		case <-sess.ctx.Done():
			return
		default:
		} // select

		// readAndHandleOneMsg 有限流
		if err = sess.readAndHandleOneMsg(); err != nil {
			break
		}
	} // for

	assert.True(err != nil, "must be error")
	sess._isError.Store(true)
	if sess.IsClosed() {
		return
	}

	// 底层检测到连接断开，可能是客户端主动Close或客户端断网超时
	log.Errorf("read and handle message error: %s, %s", err, sess.conn.RemoteAddr())
	sess.Close()
	return
}

func (sess *session) sendLoop() {
	for {
		select {
		case <-sess.ctx.Done():
			return
		case <-sess.sendBuf.HasDataC: // 可能有数据了
			if sess.sendAllBufferedData() {
				continue // 正常，继续
			}
			assert.True(sess.IsClosed())
			return // 出错了，退出
		}
	}
}

// sendAllBufferedData 发送所有缓存数据，成功则返回true, 失败则关闭会话并返回false.
func (sess *session) sendAllBufferedData() bool {
	for {
		data, err := sess.sendBuf.Pop()
		if err != nil {
			return true // 已取完
		}

		if !sess.sendData(data.([]byte)) {
			assert.True(sess.IsClosed())
			return false // 出错退出
		}
	}
}

// sendData 发送数据，成功返回true, 失败则关闭会话并返回false.
func (sess *session) sendData(buf []byte) bool {
	msgBuf, err := msgenc.CompressAndEncrypt(buf, true, sess._encryEnabled)
	if err != nil {
		log.Error("compress and encrypt message error: ", err)
		return true // Todo: 是否应该出错关闭？
	}

	_, err = sess.conn.Write(msgBuf)
	if err == nil {
		return true
	}

	sess._isError.Store(true)
	if sess.IsClosed() {
		return false
	}

	log.Error("send message error ", err)
	sess.Close()
	return false
}

// Close 关闭.
// 所有发送完成后才关闭。或2s后强制关闭。
func (sess *session) Close() {
	if !sess._isClosed.CAS(false, true) {
		return
	}

	sess.hbTimer.Stop()

	if sess.onClosed != nil {
		sess.onClosed()
	}

	if !sess._isError.Load() {
		go func() {
			closeTicker := time.NewTicker(100 * time.Millisecond)
			defer closeTicker.Stop()
			closeTimer := time.NewTimer(2 * time.Second)
			defer closeTimer.Stop()
			for {
				select {
				case <-closeTimer.C:
					sess.ctxCancel()
					sess.conn.Close()
					return
				case <-closeTicker.C:
					if sess.sendBuf.IsEmpty() {
						sess.ctxCancel()
						sess.conn.Close()
						return
					}
				}
			}
		}()
	} else {
		sess.ctxCancel()
		sess.conn.Close()
	}
}

// RemoteAddr 远程地址
func (sess *session) RemoteAddr() string {
	return sess.conn.RemoteAddr().String()
}

// IsClosed 返回sess是否已经关闭
func (sess *session) IsClosed() bool {
	return sess._isClosed.Load()
}

// RegMsgProcFunc 注册消息处理函数.
func (sess *session) RegMsgProcFunc(msgID types.MsgID, procFun func(types.IMsg)) {
	sess.msgHdlr.RegMsgProcFunc(msgID, procFun)
}

func (sess *session) readAndHandleOneMsg() error {
	msgID, rawMsgBuf, err := readARQMsg(sess.conn)
	if err != nil {
		return err
	}

	// 流量限制, 等待直到允许接收
	if err = sess.qpsLimiter.Wait(context.Background()); err != nil {
		return err
	}
	if err = sess.bpsLimiter.WaitN(context.Background(), len(rawMsgBuf)); err != nil {
		return err
	}

	assert.True(rawMsgBuf != nil, "rawMsg is nil")
	if !sess.needVerify {
		sess.msgHdlr.HandleRawMsg(msgID, rawMsgBuf)
		return nil
	}

	// 会话需要验证，第1个消息为验证请求消息
	if msgID != sess.verifyMsgID {
		msg := sess.msgCreator.NewMsg(msgID)
		vrf := sess.msgCreator.NewMsg(sess.verifyMsgID)
		return fmt.Errorf("need verify message ID %d(%s), but got %d(%s)",
			sess.verifyMsgID, reflect.TypeOf(vrf), msgID, reflect.TypeOf(msg))
	}
	sess.ResetHb()
	sess.msgHdlr.HandleRawMsg(msgID, rawMsgBuf)

	if sess.isVerified.Load() {
		sess.needVerify = false // 已通过验证，不再需要了，
		return nil
	}

	return fmt.Errorf("session verification failed")
}

// SetVerified 设置会话已通过验证.
// thread-safe.
func (sess *session) SetVerified() {
	sess.isVerified.Store(true)
}

// SetVerifyMsgID 设置会话验证消息ID.
// 非线程安全，在Session.Start()之前设置。
func (sess *session) SetVerifyMsgID(verifyMsgID types.MsgID) {
	sess.needVerify = true
	sess.verifyMsgID = verifyMsgID
}

func (sess *session) SetOnClosed(onClosed func()) {
	sess.onClosed = onClosed
}

// SetBytePerSecLimiter 设置每秒接收字节数限制.
// r(rate) 为每秒字节数。
// b(burst) 为峰值字节数。
// 必须在 Start() 之前设置，避免 DataRace.
func (sess *session) SetBytePerSecLimiter(r rate.Limit, b int) {
	sess.bpsLimiter = rate.NewLimiter(r, b)
}

// SetQueryPerSecLimiter 设置每秒接收请求数限制.
// r(rate) 为每秒请求数。
// b(burst) 为峰值请求数。
// 必须在 Start() 之前设置，避免 DataRace.
func (sess *session) SetQueryPerSecLimiter(r rate.Limit, b int) {
	sess.qpsLimiter = rate.NewLimiter(r, b)
}

// SetMsg2ID 设置 Msg2ID, 消息类型到消息ID的表.
// 非线程安全。必须在 Start() 之前设置。
func (sess *session) SetMsg2ID(msg2ID types.IMsg2ID) {
	assert.True(msg2ID != nil)
	sess.msg2ID = msg2ID
}
