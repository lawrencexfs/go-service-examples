package bll

// 属性系统
//   TODO: 优化 MsgRefreshPlayer 协议，只发送 发生变化的属性值

//属性类型
type AttrType uint8

const (
	AttrHP        AttrType = iota // 红条
	AttrHpMax                     // 红条最大值
	AttrMP                        // 蓝条
	AttrExp                       // 经验
	AttrBombNum                   // 炸弹数
	AttrHammerNum                 // 锤子数
	AttrMax
)

type Attribute struct {
	attrs [AttrMax]float64
}

func (this *Attribute) GetAttr(index AttrType) float64 {
	return this.attrs[index]
}

func (this *Attribute) SetAttr(index AttrType, val float64) {
	this.attrs[index] = val
}

func (this *Attribute) GetHP() int32 {
	return int32(this.GetAttr(AttrHP))
}

func (this *Attribute) SetHP(val int32) {
	this.SetAttr(AttrHP, float64(val))
}

func (this *Attribute) GetHpMax() int32 {
	return int32(this.GetAttr(AttrHpMax))
}

func (this *Attribute) SetHpMax(val int32) {
	this.SetAttr(AttrHpMax, float64(val))
}

func (this *Attribute) GetMP() float64 {
	return this.GetAttr(AttrMP)
}

func (this *Attribute) SetMP(val float64) {
	this.SetAttr(AttrMP, val)
}
