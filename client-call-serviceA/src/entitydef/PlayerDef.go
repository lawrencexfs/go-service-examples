package entitydef

import "github.com/giant-tech/go-service/framework/iserver"

// PlayerDef 自动生成的属性包装代码
type PlayerDef struct {
	ip iserver.IEntityProps
}

// SetPropsSetter 设置接口
func (p *PlayerDef) SetPropsSetter(ip iserver.IEntityProps) {
	p.ip = ip
}

// Setcoin 设置 coin
func (p *PlayerDef) Setcoin(v uint32) {
	p.ip.SetProp("coin", v)
}

// SetcoinDirty 设置coin被修改
func (p *PlayerDef) SetcoinDirty() {
	p.ip.PropDirty("coin")
}

// Getcoin 获取 coin
func (p *PlayerDef) Getcoin() uint32 {
	v := p.ip.GetProp("coin")
	if v == nil {
		return uint32(0)
	}

	return v.(uint32)
}

// Setexp 设置 exp
func (p *PlayerDef) Setexp(v uint32) {
	p.ip.SetProp("exp", v)
}

// SetexpDirty 设置exp被修改
func (p *PlayerDef) SetexpDirty() {
	p.ip.PropDirty("exp")
}

// Getexp 获取 exp
func (p *PlayerDef) Getexp() uint32 {
	v := p.ip.GetProp("exp")
	if v == nil {
		return uint32(0)
	}

	return v.(uint32)
}

// Setlevel 设置 level
func (p *PlayerDef) Setlevel(v uint32) {
	p.ip.SetProp("level", v)
}

// SetlevelDirty 设置level被修改
func (p *PlayerDef) SetlevelDirty() {
	p.ip.PropDirty("level")
}

// Getlevel 获取 level
func (p *PlayerDef) Getlevel() uint32 {
	v := p.ip.GetProp("level")
	if v == nil {
		return uint32(0)
	}

	return v.(uint32)
}

// Setrating 设置 rating
func (p *PlayerDef) Setrating(v int32) {
	p.ip.SetProp("rating", v)
}

// SetratingDirty 设置rating被修改
func (p *PlayerDef) SetratingDirty() {
	p.ip.PropDirty("rating")
}

// Getrating 获取 rating
func (p *PlayerDef) Getrating() int32 {
	v := p.ip.GetProp("rating")
	if v == nil {
		return int32(0)
	}

	return v.(int32)
}

type IPlayerDef interface {
	Setcoin(v uint32)
	SetcoinDirty()
	Getcoin() uint32
	Setexp(v uint32)
	SetexpDirty()
	Getexp() uint32
	Setlevel(v uint32)
	SetlevelDirty()
	Getlevel() uint32
	Setrating(v int32)
	SetratingDirty()
	Getrating() int32
}
