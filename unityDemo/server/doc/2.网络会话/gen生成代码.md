# gen 生成代码

消息ID配置为 [`gen/wilds.toml`](../../src/roomserver/gen/wilds.toml),
运行 [`generate.bat`](../../src/roomserver/gen/generate.bat) 生成代码，
服务器代码为 `gen/gensvr/`.

## `generated_server.go`
```go
// 使用 gensvr.New() 创建服务器可以保证注册了所有消息和处理器。
func New(protocal string, addr string, maxConns int) (*Server, error) {
	svr, err := server.New(protocal, addr, maxConns)
	if err != nil {
		return nil, err
	}

	// 注册消息
	internal.RegMsg_wilds(svr)

	// 添加MsgProc, 这样新连接创建时会
	// CloneAndRegisterMsgProcFunctions() 注册所有处理函数。
	svr.AddMsgProc(&internal.Wrapper_wilds{})

	return svr, nil
}
```

## `proc/Proc_wilds.go`
``` go
// Proc_wilds 是消息处理类(Processor).
// 必须实现 NewProc_wilds(), OnClosed() 和 MsgProc_*() 接口。
type Proc_wilds struct {
	sess server.ISession // 一般都需要包含session对象

	// 可能还应该包含用户和房间对象
	// user *User
	// room *Room
}

func NewProc_wilds(sess server.ISession) *Proc_wilds {
	return &Proc_wilds{
		sess: sess,
		// user, room 暂时为空，待创建
	}
}

func (p *Proc_wilds) OnClosed() {
	// 会话断开时动作...
}

func (p *Proc_wilds) MsgProc_MsgLogin(msg *usercmd.MsgLogin) {
	panic("待实现")
}
...
func (p *Proc_wilds) MsgProc_MsgCastSkill(msg *usercmd.MsgCastSkill) {
	panic("待实现")
}
```