package birth

import (
	"math"

	bmath "battleservice/src/services/base/math"
	"battleservice/src/services/battle/scene/bll"
	"battleservice/src/services/battle/scene/consts"
	"battleservice/src/services/battle/scene/interfaces"
	"battleservice/src/services/battle/scene/typekind"
	"battleservice/src/services/battle/usercmd"

	"github.com/cihub/seelog"
)

type BirthPoint struct {
	id             uint64
	scene          IScene
	pos            bmath.Vector2
	ballTypeId     uint16
	ballType       uint16
	birthTime      float64
	birthMax       uint32
	birthRadiusMin float32
	birthRadiusMax float32
	childrenCount  uint32
	birthTimer     float64
}

//创建动态出生点 食物、 动态障碍物 (BallFood、 BallFeed)
func NewBirthPoint(id uint64, x, y, rMin, rMax float32, ballTypeId uint16, ballType uint16, birthTime float64, birthMax uint32, scene IScene) *BirthPoint {

	point := &BirthPoint{
		id:         id,
		pos:        bmath.Vector2{x, y},
		ballTypeId: ballTypeId,
		ballType:   ballType,
		birthTime:  birthTime,
		birthMax:   birthMax,
		scene:      scene,
	}
	point.birthRadiusMin = rMin
	point.birthRadiusMax = rMax
	point.Init()
	return point
}

func (this *BirthPoint) Init() {
	var i uint32 = 0
	for ; i < this.birthMax; i++ {
		this.CreateUnit()
	}
}

func (this *BirthPoint) CreateUnit() interfaces.IBall {
	this.childrenCount++

	scene := this.scene.(bll.IScene)

	var ball interfaces.IBall
	ballType := typekind.BallTypeToKind(usercmd.BallType(this.ballType))

	switch ballType {
	case consts.BallKind_Food:
		posNew := BallFood_InitPos(&this.pos, usercmd.BallType(this.ballType), this.birthRadiusMin, this.birthRadiusMax)

		foodInitData := &bll.FoodInitData{
			ID:         this.id,
			TypeID:     this.ballTypeId,
			X:          float64(posNew.X),
			Y:          float64(posNew.Y),
			Scene:      scene,
			BirthPoint: this,
		}

		scene.CreateEntity("BallFood", scene.GetEntityID(), foodInitData, true, 0)
	case consts.BallKind_Feed:
		x := math.Floor(float64(this.pos.X)) + 0.25
		y := math.Floor(float64(this.pos.Y)) + 0.25

		initData := &bll.FeedInitData{
			Scene:      scene,
			TypeID:     this.ballTypeId,
			ID:         this.id,
			X:          x,
			Y:          y,
			BirthPoint: this,
		}

		scene.CreateEntity("BallFeed", scene.GetEntityID(), initData, true, 0)

	default:
		seelog.Error("CreateUnit unknow ballType:", ballType, "  typeid:", this.ballTypeId)
	}

	return ball
}

func (this *BirthPoint) Refresh(perTime float64, scene IScene) {
	if this.childrenCount >= this.birthMax {
		return
	}
	if this.birthTimer >= this.birthTime {
		this.birthTimer = 0
		this.CreateUnit()
	} else {
		this.birthTimer += perTime
	}
}

func (this *BirthPoint) OnChildRemove(ball interfaces.IBall) {
	this.childrenCount--
}
