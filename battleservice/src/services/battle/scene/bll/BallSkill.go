package bll

// 技能球

import (
	bmath "battleservice/src/services/base/math"
	"battleservice/src/services/battle/scene/interfaces"
	"battleservice/src/services/battle/usercmd"
	"fmt"

	"github.com/giant-tech/go-service/base/linmath"

	"github.com/cihub/seelog"
)

// SkillInitData 初始化数据
type SkillInitData struct {
	BallType usercmd.BallType
	ID       uint64
	Pos      linmath.Vector3
	Radius   float32
	Player   IScenePlayer
}

// BallSkill 技能球
type BallSkill struct {
	BallMove
	player IScenePlayer
	Skill  interfaces.ISkillBall
}

// OnInit 初始化
func (s *BallSkill) OnInit(initData interface{}) error {
	seelog.Info("BallSkill.OnInit, id:", s.GetEntityID())
	SkillInitData, ok := initData.(*SkillInitData)
	if !ok {
		return fmt.Errorf("init data error")
	}

	s.BallMove = BallMove{
		BallFood: BallFood{
			id:       SkillInitData.ID,
			typeID:   uint16(SkillInitData.BallType),
			BallType: SkillInitData.BallType,
			Pos:      SkillInitData.Pos,
			radius:   SkillInitData.Radius,
		},
	}

	s.player = SkillInitData.Player

	// ball.Skill = MyProvider.NewSkillBall(player, &ball)
	s.Skill = SkillInitData.Player.NewSkillBall(s) // XXX 临时实现，应该有更好的方法
	s.ResetRect()

	return nil
}

// OnLoop 每帧调用
func (s *BallSkill) OnLoop() {
	seelog.Debug("BallSkill.OnLoop")

	if s.Skill.IsFinish() {
		//TODO: skill删除
	}

	// 检查移动是否出格子
	if s.Move(0, nil) {
		//cell.AddMsgMove(ball)
		// 如果技能球已移动新的格子，则更新，删除旧格中的球，添加到新格。
		//scene.UpdateSkillBallCell(ball, cell.id)

		s.ResetRect()
	}
	s.Skill.Update()
}

// OnDestroy 销毁
func (s *BallSkill) OnDestroy() {
	seelog.Debug("BallSkill.OnDestroy")
}

func (s *BallSkill) Move(pertime float64, scene IScene) bool {
	if !s.speed.IsEqual(linmath.Vector3{0, 0, 0}) {
		if s.player.Frame() >= s.Skill.GetBeginFrame() {
			if false {
				pos := s.PhysicObj.GetPostion()
				s.Pos = linmath.Vector3{pos.X, 0, pos.Y}
				s.PhysicObj.SetVelocity(&bmath.Vector2{float32(s.speed.X), float32(s.speed.Z)})
			}

			s.Pos.X = s.Pos.X + s.speed.X/2
			s.Pos.Z = s.Pos.Z + s.speed.Z/2
		}
	}
	return true
}
