package bll

// 食物球

import (
	"battleservice/src/services/base/util"
	"battleservice/src/services/battle/conf"
	"battleservice/src/services/battle/scene/consts"
	"battleservice/src/services/battle/scene/interfaces"
	"battleservice/src/services/battle/usercmd"
	"fmt"

	"github.com/cihub/seelog"
	"github.com/giant-tech/go-service/framework/space"
)

// FoodInitData 初始化数据
type FoodInitData struct {
	ID         uint64
	TypeID     uint16
	X, Y       float64
	Scene      IScene
	BirthPoint interfaces.IBirthPoint
}

// BallFood 食物球
type BallFood struct {
	space.Entity
	id         uint64           //动态id
	typeID     uint16           //xml表里id
	BallType   usercmd.BallType //大类型
	Pos        util.Vector2
	radius     float64
	rect       util.Square
	birthPoint interfaces.IBirthPoint
	exp        int32
}

// OnInit 初始化
func (ball *BallFood) OnInit(initData interface{}) error {
	seelog.Info("BallFood.OnInit, id:", ball.GetEntityID())

	foodInitData, ok := initData.(*FoodInitData)
	if !ok {
		return fmt.Errorf("init data error")
	}

	var radius float32 = conf.ConfigMgr_GetMe().GetFoodSize(foodInitData.Scene.GetEntityID(), foodInitData.TypeID)
	ballType := conf.ConfigMgr_GetMe().GetFoodBallType(foodInitData.Scene.GetEntityID(), foodInitData.TypeID)

	ball.id = foodInitData.ID
	ball.typeID = foodInitData.TypeID
	ball.Pos = util.Vector2{foodInitData.X, foodInitData.Y}
	ball.BallType = ballType
	ball.radius = float64(radius)

	ball.ResetRect()
	ball.SetExp(consts.DefaultBallFoodExp)

	return nil
}

// OnLoop 每帧调用
func (ball *BallFood) OnLoop() {
	seelog.Debug("BallFood.OnLoop")
}

// OnDestroy 销毁
func (ball *BallFood) OnDestroy() {
	seelog.Debug("BallFood.OnDestroy")
}

func (ball *BallFood) GetRect() *util.Square {
	return &ball.rect
}

func (ball *BallFood) OnReset() {

}

func (ball *BallFood) GetID() uint64 {
	return ball.id
}

func (ball *BallFood) SetID(id uint64) {
	ball.id = id
}

func (ball *BallFood) GetTypeId() uint16 {
	return ball.typeID
}

func (ball *BallFood) GetBallType() usercmd.BallType {
	return ball.BallType
}

func (ball *BallFood) GetPos() (float64, float64) {
	return ball.Pos.X, ball.Pos.Y
}

func (ball *BallFood) SetPos(x, y float64) {
	ball.Pos.X = x
	ball.Pos.Y = y
}

func (ball *BallFood) GetPosV() *util.Vector2 {
	return &ball.Pos
}

func (this *BallFood) SetPosV(pos util.Vector2) {
	this.Pos = pos
}

func (ball *BallFood) SetExp(exp int32) {
	ball.exp = exp
}

func (ball *BallFood) ResetRect() {
	ball.rect.X = ball.Pos.X
	ball.rect.Y = ball.Pos.Y
	ball.rect.SetRadius(ball.radius)
}

func (ball *BallFood) SetBirthPoint(birthPoint interfaces.IBirthPoint) {
	ball.birthPoint = birthPoint
}

func (ball *BallFood) GetBirthPoint() interfaces.IBirthPoint {
	return ball.birthPoint
}

func (ball *BallFood) GetRadius() float64 {
	return ball.radius
}
