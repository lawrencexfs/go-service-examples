package bll

// 玩家球

import (
	"battleservice/src/services/base/ape"
	bmath "battleservice/src/services/base/math"
	"battleservice/src/services/base/util"
	"battleservice/src/services/battle/scene/consts"
	"battleservice/src/services/battle/usercmd"

	"github.com/giant-tech/go-service/base/linmath"
)

type BallPlayer struct {
	BallMove
	Attribute
	player IScenePlayer
}

func (this *BallPlayer) InitBallPlayer(player IScenePlayer, ballid uint64) {

	x, z := player.GetBallScene().GetRandPos()

	this.BallMove = BallMove{
		BallFood: BallFood{
			id: ballid,

			radius:   consts.DefaultBallSize,
			BallType: usercmd.BallType_Player,
		},
	}

	this.SetPos(linmath.Vector3{x, 0, z})

	this.player = player

	this.ResetRect()

	this.SetMP(consts.DefaultMaxMP)
	this.SetHpMax(consts.DefaultMaxHP)
	this.SetHP(consts.DefaultMaxHP)

	this.PhysicObj = ape.NewCircleParticle(float32(this.GetPos().X), float32(this.GetPos().Z), float32(this.radius))
	this.player.GetBallScene().AddPlayerPhysic(this.PhysicObj)
}

func (this *BallPlayer) GetPlayerId() uint64 {
	return this.player.GetID()
}

func (this *BallPlayer) GetPlayer() IScenePlayer {
	return this.player
}

//定时器1s
func (this *BallPlayer) TimeAction(nowsec int64) {

}

func (this *BallPlayer) Move(perTime float64, frameRate float64) bool {
	// 有推力情况下， 忽略原来速度方向
	if this.HasForce() == true {
		force := this.GetForce()
		pos := this.PhysicObj.GetPostion()
		this.SetPos(linmath.Vector3{pos.X, 0, pos.Y})
		this.PhysicObj.SetVelocity(&bmath.Vector2{float32(force.X), float32(force.Z)})
		return true
	}

	pos := this.PhysicObj.GetPostion()
	this.SetPos(linmath.Vector3{pos.X, 0, pos.Y})

	speed := consts.DefaultBallSpeed

	powerMul := util.Clamp(this.player.GetPower(), 0, 1)

	if this.player.IsRunning() {
		speed *= consts.DefaultRunRatio
		powerMul = 1
	}

	speed *= powerMul
	this.speed = this.angleVel.Mul(float32(speed))

	vel := this.speed
	vel.MulS(float32(frameRate)) //几帧执行一次物理tick
	if 0 == this.player.GetPower() {
		this.PhysicObj.SetVelocity(&bmath.Vector2{0, 0})
	} else {
		this.PhysicObj.SetVelocity(&bmath.Vector2{float32(vel.X) / 30, float32(vel.Z) / 30})
	}

	return true
}
func (this *BallPlayer) FixMapEdge() bool {
	SceneSize := this.player.GetBallScene().SceneSize()
	halfRadius := this.radius * 0.5
	pos := linmath.Vector3{}

	if this.GetPos().X < halfRadius {
		pos.X = halfRadius
		this.speed.X = -this.speed.X * 0.1
	} else if this.GetPos().X > SceneSize-halfRadius {
		pos.X = SceneSize - halfRadius
		this.speed.X = -this.speed.X * 0.1
	}
	if this.GetPos().Z < halfRadius {
		pos.Z = halfRadius
		this.speed.Z = -this.speed.Z * 0.1
	} else if this.GetPos().Z > SceneSize-halfRadius {
		pos.Z = SceneSize - halfRadius
		this.speed.Z = -this.speed.Z * 0.1
	}

	this.SetPos(pos)

	this.rect.X = float64(this.GetPos().X)
	this.rect.Z = float64(this.GetPos().Z)

	return true
}

//攻击预先判断
func (this *BallPlayer) PreTryHit(target *BallPlayer) bool {
	if !this.player.GetIsLive() || !target.player.GetIsLive() {
		return false
	}
	return true
}

func (this *BallPlayer) Hit(target *BallPlayer) (int32, bool) {
	damage := consts.DefaultAttack
	target.SetHP(target.GetHP() - int32(damage))
	if target.GetHP() <= 0 {
		target.player.KilledByPlayer(this.player)
	}
	return int32(damage), true
}

func (this *BallPlayer) Eat(food *BallFood) uint32 {
	if food.GetBallType() == usercmd.BallType_FoodBomb {
		this.SetAttr(AttrBombNum, 1)
		this.player.RefreshPlayer()
	} else if food.GetBallType() == usercmd.BallType_FoodHammer {
		this.SetAttr(AttrHammerNum, 1)
		this.player.RefreshPlayer()
	}

	player := this.player

	var addexp uint32 = 0
	if food.exp != 0 {
		addexp = uint32(food.exp)
	}
	if 0 != addexp {
		player.UpdateExp(int32(addexp))
	}
	return addexp
}

func (this *BallPlayer) PreCanEat(food *BallFood) bool {
	//是否已经有锤子或者炸弹了
	if food.GetBallType() == usercmd.BallType_FoodHammer {
		return this.GetAttr(AttrHammerNum) == 0
	} else if food.GetBallType() == usercmd.BallType_FoodBomb {
		return this.GetAttr(AttrBombNum) == 0
	}
	return true
}

func (this *BallPlayer) GetEatRange() float64 {
	r := consts.DefaultEatFoodRange
	if r == 0 {
		return float64(this.radius)
	}
	return r * this.GetSizeScale()
}

func (this *BallPlayer) isNear(target *BallFood) bool {
	distance := this.SqrMagnitudeTo(target)
	eatRange := float32(this.GetEatRange())
	return distance <= (eatRange+target.radius)*(eatRange+target.radius)
}

func (this *BallPlayer) CanEat(food *BallFood) bool {
	if !this.PreCanEat(food) {
		return false
	}
	return this.isNear(food)
}

func (this *BallPlayer) OnDead() {

}

func (this *BallPlayer) GetSizeScale() float64 {
	return 1.0
}

func (this *BallPlayer) SetAngleVelAndNormalize(x, z float64) {
	this.angleVel.X = float32(x)
	this.angleVel.Z = float32(z)
	this.angleVel.Normalize()
}
