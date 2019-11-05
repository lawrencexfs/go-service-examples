package entitydef

import (
	"gitlab.ztgame.com/tech/public/go-service/zeus/framework/iserver"
	"github.com/globalsign/mgo/bson"
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

// SetLevel 设置 Level
func (p *PlayerDef) SetLevel(v uint32) {
	p.ip.SetProp("Level", v)
}

// SetLevelDirty 设置Level被修改
func (p *PlayerDef) SetLevelDirty() {
	p.ip.PropDirty("Level")
}

// GetLevel 获取 Level
func (p *PlayerDef) GetLevel() uint32 {
	v := p.ip.GetProp("Level")
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
	SetCoin(v uint32)
	SetCoinDirty()
	GetCoin() uint32
	SetFriends(v FRIENDS)
	SetFriendsDirty()
	GetFriends() FRIENDS
	SetHero(v HEROS)
	SetHeroDirty()
	GetHero() HEROS
	SetLevel(v uint32)
	SetLevelDirty()
	GetLevel() uint32
	Setrating(v int32)
	SetratingDirty()
	Getrating() int32
}
