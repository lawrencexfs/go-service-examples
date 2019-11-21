package bll

// 技能球

import (
	bmath "battleservice/src/services/base/math"
	"battleservice/src/services/base/util"
	"battleservice/src/services/battle/scene/interfaces"
	"battleservice/src/services/battle/usercmd"

	"github.com/cihub/seelog"
)

type BallSkill struct {
	BallMove
	player IScenePlayer
	Skill  interfaces.ISkillBall
}

// OnInit 初始化
func (s *BallSkill) OnInit(initData interface{}) error {
	seelog.Info("BallFood.OnInit, id:", s.GetEntityID())

	return nil
}

// OnLoop 每帧调用
func (s *BallSkill) OnLoop() {
	seelog.Debug("BallFood.OnLoop")
}

// OnDestroy 销毁
func (s *BallSkill) OnDestroy() {
	seelog.Debug("BallFood.OnDestroy")
}

func NewBallSkill(_ballType usercmd.BallType, id uint64, x, y, radius float64, player IScenePlayer) *BallSkill {
	ball := BallSkill{
		BallMove: BallMove{
			BallFood: BallFood{
				id:       id,
				typeID:   uint16(_ballType),
				BallType: _ballType,
				Pos:      util.Vector2{float64(x), float64(y)},
				radius:   float64(radius),
			},
		},
		player: player,
	}

	// ball.Skill = MyProvider.NewSkillBall(player, &ball)
	ball.Skill = player.NewSkillBall(&ball) // XXX 临时实现，应该有更好的方法
	return &ball
}

func (s *BallSkill) Move(pertime float64, scene IScene) bool {
	if s.speed.IsEmpty() == false {
		if s.player.Frame() >= s.Skill.GetBeginFrame() {
			if false {
				pos := s.PhysicObj.GetPostion()
				s.Pos = util.Vector2{float64(pos.X), float64(pos.Y)}
				s.PhysicObj.SetVelocity(&bmath.Vector2{float32(s.speed.X), float32(s.speed.Y)})
			}

			s.Pos.X = s.Pos.X + s.speed.X/2
			s.Pos.Y = s.Pos.Y + s.speed.Y/2
		}
	}
	return true
}
