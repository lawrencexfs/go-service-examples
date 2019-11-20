package usercmd

import (
	"github.com/giant-tech/go-service/base/linmath"
)

// EnterAOI 玩家进入AOI
type EnterAOI struct {
	EntityID   uint64
	EntityType string
	State      []byte
	PropNum    uint16
	Properties []byte

	Pos  linmath.Vector3
	Rota linmath.Vector3
}

// LeaveAOI 离开AOI
type LeaveAOI struct {
	EntityID uint64
}

// NewEntityAOISMsg 新建进入AOI消息
func NewEntityAOISMsg() *EntityAOIS {
	return &EntityAOIS{
		0,
		make([][]byte, 0, 1),
	}
}

// EntityAOIS 进入AOI范围
type EntityAOIS struct {
	Num  uint32
	data [][]byte
}

// AddData 增加一个玩家数据
func (msg *EntityAOIS) AddData(data []byte) {
	msg.Num++
	msg.data = append(msg.data, data)

	// seelog.Debug("EntityAOIS AddData ", msg)
}

// Clear 清理
func (msg *EntityAOIS) Clear() {
	msg.Num = 0
	msg.data = msg.data[0:0]
}
