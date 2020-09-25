package proc

import (
	"github.com/giant-tech/go-service/framework/idata"
	"github.com/giant-tech/go-service/framework/msghandler"
	"github.com/giant-tech/go-service/framework/net/inet"

	assert "github.com/aurelien-rainone/assertgo"
	log "github.com/cihub/seelog"
	"github.com/giant-tech/go-service/framework/msgdef"
)

// IClient Interface to clt.Client.
type IClient interface {
	msghandler.IRPCHandlers
	Post(func())
	AsyncCall(st idata.ServiceType, methodName string, args ...interface{}) error
	SyncCall(st idata.ServiceType, retPtr interface{}, methodName string, args ...interface{}) error
}

// ProcGateway 是消息处理类(Processor).
// 必须实现 MsgProc_*() 接口。
type ProcGateway struct {
	Sess   inet.ISession // 一般都需要包含session对象
	Client IClient       // 指向 clt.Client
	Uid    uint64
}

// RegisterMsgProcFunctions 克隆自身并注册消息处理函数.
func (p *ProcGateway) RegisterMsgProcFunctions(sess inet.ISession) interface{} {
	assert.True(sess != nil, "session is nil")
	proc := &ProcGateway{
		Sess: sess,
	}

	sess.RegMsgProc(proc)

	sess.AddOnClosed(proc.OnClosed)

	return proc
}

func (p *ProcGateway) MsgProcLoginResp(msg *msgdef.LoginResp) {

}

// OnClosed 关闭处理
func (p *ProcGateway) OnClosed() {
	// 会话断开时动作...
}

// MsgProcCallMsg callmsg处理
func (p *ProcGateway) MsgProcCallMsg(msg *msgdef.CallMsg) {
	//log.Infof("MsgProcCallMsg MethodName:%s", msg.MethodName)

	p.Client.Post(func() {
		log.Debug("methodname = ", msg.MethodName)
		p.Client.DoRPCMsg(msg.MethodName, msg.Params, nil)
	})
}
