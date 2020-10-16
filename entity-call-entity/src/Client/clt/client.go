package clt

import (
	"entity-call-entity/src/Client/proc"
	"fmt"
	"sync"
	"time"

	"github.com/giant-tech/go-service/base/serializer"
	"github.com/giant-tech/go-service/framework/idata"
	"github.com/giant-tech/go-service/framework/msgdef"
	"github.com/giant-tech/go-service/framework/msghandler"
	"github.com/giant-tech/go-service/framework/net/client"
	"github.com/giant-tech/go-service/framework/net/inet"

	"go.uber.org/atomic"

	log "github.com/cihub/seelog"
)

// Client 客户端结构体
type Client struct {
	msghandler.RPCHandlers

	lobbySession inet.ISession

	pendingMap   sync.Map
	seq          atomic.Uint64
	pendingClose chan bool

	// ActC 是其他协程向自己Run协程发送动作的Channel，Run()中将依次执行动作。
	// MsgProc 网络消息处理协程会添加动作，在Run()中执行。
	// Action 动作, 是无参数无返回值的函数.
	ActC chan func()
}

// Init 初始化
func (cli *Client) Init(addr string) {
	if addr == "" {
		panic(fmt.Errorf("Client.init addr ni empty"))
	}

	// 与 LobbyServer 建立链接
	var err error
	sess, err := client.Dial("tcp", addr)
	if err != nil {
		panic(fmt.Errorf("dial lobby error: %s", err))
	}

	lp := &proc.ProcGateway{}
	lobbyProc := lp.RegisterMsgProcFunctions(sess).(*proc.ProcGateway)

	cli.lobbySession = sess
	lobbyProc.Client = cli
	//proc.Uid = cli.uid

	sess.Start()

	cli.pendingClose = make(chan bool, 1)
	go cli.loopCheckPendingCall(cli.pendingClose)
}

// AsyncCall 消息异步调用
func (cli *Client) AsyncCall(sType idata.ServiceType, methodName string, args ...interface{}) error {
	data := serializer.SerializeNew(args...)
	msg := &msgdef.CallMsg{
		SType:      uint8(sType),
		MethodName: methodName,
		Params:     data,
	}

	return cli.lobbySession.Send(msg)
}

// SyncCall 同步调用，等待返回
func (cli *Client) SyncCall(sType idata.ServiceType, retPtr interface{}, methodName string, args ...interface{}) error {
	msg := &msgdef.CallMsg{}

	msg.SType = uint8(sType)
	msg.MethodName = methodName
	msg.IsSync = true
	msg.Params = serializer.SerializeNew(args...)

	msg.Seq = cli.GetSeq()
	cli.lobbySession.Send(msg)

	//加入到pending中
	call := &idata.PendingCall{}
	call.RetChan = make(chan *idata.RetData, 1)
	call.Seq = msg.Seq
	call.MethodName = methodName
	call.Reply = retPtr
	call.StartTime = time.Now().Unix()
	cli.AddPendingCall(call)

	retData := <-call.RetChan
	if retData.Err != nil {
		log.Error("Client.SyncCall, remote retData.Err: ", retData.Err)
		return retData.Err
	}

	if retPtr != nil {
		if err := serializer.UnSerializeNew(retPtr, retData.Ret); err != nil {
			return err
		}
	}

	return nil
}

// AddPendingCall 加入等待调用
func (cli *Client) AddPendingCall(call *idata.PendingCall) {
	//seelog.Debug("AddPendingCall, seq: ", call.Seq, ", startTime: ", call.StartTime)
	cli.pendingMap.Store(call.Seq, call)
}

// DelPendingCall 删除等待调用
func (cli *Client) DelPendingCall(seq uint64) {
	//seelog.Debug("delPendingCall, seq: ", seq)
	cli.pendingMap.Delete(seq)
}

// GetPendingCall 获得等待调用
func (cli *Client) GetPendingCall(seq uint64) *idata.PendingCall {
	call, ok := cli.pendingMap.Load(seq)
	if ok {
		return call.(*idata.PendingCall)
	}

	return nil
}

// loopCheckPendingCall 循环检测等待调用
func (cli *Client) loopCheckPendingCall(closeSig chan bool) {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-closeSig:
			return

		case <-ticker.C:
			cli.delTimeoutPendingCall()
		}
	}
}

// delTimeoutPendingCall 删除调用
func (cli *Client) delTimeoutPendingCall() {
	//seelog.Debug("delTimeoutPendingCall")

	cli.pendingMap.Range(
		func(key, value interface{}) bool {
			call := value.(*idata.PendingCall)
			//seelog.Debug("call.StartTime: ", call.StartTime, ", now: ", time.Now().Unix())

			if call.StartTime+10 < time.Now().Unix() {
				cli.DelPendingCall(key.(uint64))
				retData := &idata.RetData{}
				retData.Err = fmt.Errorf("call timeout")
				call.RetChan <- retData
			}

			return true
		})
}

// Post Post
func (cli *Client) Post(act func()) {
	cli.ActC <- act
}

// IsFinal 是否final
func (cli *Client) IsFinal() bool {
	if cli.lobbySession != nil && cli.lobbySession.IsClosed() {
		return true
	}

	return false
}

// Run 运行
func (cli *Client) Run() {
	// log.Info("run...")
	ticker := time.NewTicker(time.Duration(50) * time.Millisecond)
	defer ticker.Stop()
	for !cli.IsFinal() {
		select {
		case act := <-cli.ActC:
			act()
		case <-ticker.C:
			cli.tick()
		}
	}
	log.Debug("Client exit. final:", cli.IsFinal())
}

// tick tick
func (cli *Client) tick() {
}

// GetSeq GetSeq
func (cli *Client) GetSeq() uint64 {
	return cli.seq.Inc()
}

// Login login
func (cli *Client) Login() {

	log.Debug("Client login")

	msg := &msgdef.LoginReq{
		Token: "abc",
		UID:   1,
	}

	if cli.lobbySession == nil {
		panic("lobbySession is nil")
	}

	cli.lobbySession.Send(msg)
}
