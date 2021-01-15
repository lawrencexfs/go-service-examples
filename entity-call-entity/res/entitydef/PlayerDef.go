package entitydef

import (
	"github.com/giant-tech/go-service/framework/iserver"
	"github.com/globalsign/mgo/bson"
	protoMsg "entity-call-entity/src/pb"
	proto "github.com/golang/protobuf/proto"
)

// PlayerDef 自动生成的属性包装代码
type PlayerDef struct {
	ip iserver.IEntityProps
}

// SetPropsSetter 设置接口
func (p *PlayerDef) SetPropsSetter(ip iserver.IEntityProps) {
	p.ip = ip
}

// SetCoin 设置 Coin
func (p *PlayerDef) SetCoin(v uint32) {
	p.ip.SetProp("Coin", v)
}

// SetCoinDirty 设置Coin被修改
func (p *PlayerDef) SetCoinDirty() {
	p.ip.PropDirty("Coin")
}

// GetCoin 获取 Coin
func (p *PlayerDef) GetCoin() uint32 {
	v := p.ip.GetProp("Coin")
	if v == nil {
		return uint32(0)
	}

	return v.(uint32)
}

// SetFriends 设置 Friends
func (p *PlayerDef) SetFriends(v *FRIENDS) {
	p.ip.SetProp("Friends", v)
}

// SetFriendsDirty 设置Friends被修改
func (p *PlayerDef) SetFriendsDirty() {
	p.ip.PropDirty("Friends")
}

// GetFriends 获取 Friends
func (p *PlayerDef) GetFriends() *FRIENDS {
	v := p.ip.GetProp("Friends")
	if v == nil {
		return nil
	}

	var tempV FRIENDS
	bson.Unmarshal(v.([]byte), &tempV)
	return &tempV
}

// SetHero 设置 Hero
func (p *PlayerDef) SetHero(v *HEROS) {
	p.ip.SetProp("Hero", v)
}

// SetHeroDirty 设置Hero被修改
func (p *PlayerDef) SetHeroDirty() {
	p.ip.PropDirty("Hero")
}

// GetHero 获取 Hero
func (p *PlayerDef) GetHero() *HEROS {
	v := p.ip.GetProp("Hero")
	if v == nil {
		return nil
	}

	var tempV HEROS
	bson.Unmarshal(v.([]byte), &tempV)
	return &tempV
}

// Setbullet 设置 bullet
func (p *PlayerDef) Setbullet(v *protoMsg.ChangeBulletReq) {
	p.ip.SetProp("bullet", v)
}

// SetbulletDirty 设置bullet被修改
func (p *PlayerDef) SetbulletDirty() {
	p.ip.PropDirty("bullet")
}

// Getbullet 获取 bullet
func (p *PlayerDef) Getbullet() *protoMsg.ChangeBulletReq {
	v := p.ip.GetProp("bullet")
	if v == nil {
		return nil
	}

	var tempV protoMsg.ChangeBulletReq
	proto.Unmarshal(v.([]byte), &tempV)
	return &tempV
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

// Setname 设置 name
func (p *PlayerDef) Setname(v string) {
	p.ip.SetProp("name", v)
}

// SetnameDirty 设置name被修改
func (p *PlayerDef) SetnameDirty() {
	p.ip.PropDirty("name")
}

// Getname 获取 name
func (p *PlayerDef) Getname() string {
	v := p.ip.GetProp("name")
	if v == nil {
		return ""
	}

	return v.(string)
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
	SetCoin(v uint32)
	SetCoinDirty()
	GetCoin() uint32
	SetFriends(v FRIENDS)
	SetFriendsDirty()
	GetFriends() FRIENDS
	SetHero(v HEROS)
	SetHeroDirty()
	GetHero() HEROS
	Setbullet(v protoMsg.ChangeBulletReq)
	SetbulletDirty()
	Getbullet() protoMsg.ChangeBulletReq
	Setexp(v uint32)
	SetexpDirty()
	Getexp() uint32
	Setlevel(v uint32)
	SetlevelDirty()
	Getlevel() uint32
	Setname(v string)
	SetnameDirty()
	Getname() string
	Setrating(v int32)
	SetratingDirty()
	Getrating() int32
}
