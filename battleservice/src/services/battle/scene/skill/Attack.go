package skill

// 用于技能伤害计算。注意这里主要是对目标做伤害计算，任何技能行为树最终会调到该文件内的方法。

import (
	b3core "battleservice/src/services/base/behavior3go/core"
	"battleservice/src/services/battle/scene/bll"
	"battleservice/src/services/battle/scene/consts"
	"battleservice/src/services/battle/scene/interfaces"
	"battleservice/src/services/battle/scene/plr"
	"battleservice/src/services/battle/usercmd"
)

type AttackType uint8

const (
	ATTACK_TYPE_SINGLE          AttackType = 1
	ATTACK_TYPE_AOE_SEMI_CIRCLE AttackType = 2
	ATTACK_TYPE_AOE_CIRCLE      AttackType = 3
)

// 技能攻击伤害计算入口
func SkillAttack(tick *b3core.Tick, player *plr.ScenePlayer, attackType int) {
	if player.IsLive == false {
		return
	}
	at := AttackType(attackType)
	switch at {
	case ATTACK_TYPE_SINGLE:
		SingleAttack(tick, player)
	case ATTACK_TYPE_AOE_SEMI_CIRCLE:
		AOEAttack_SemiCircle(tick, player)
	case ATTACK_TYPE_AOE_CIRCLE:
		AOEAttack_Circle(tick, player)
	default:
	}
}

// 单体攻击
func SingleAttack(tick *b3core.Tick, player *plr.ScenePlayer) {
	ball, balltype := FindNearTarget(tick, player)
	NormalAttack(tick, player, ball, balltype)
}

// 群体攻击 - 半圆范围
func AOEAttack_SemiCircle(tick *b3core.Tick, player *plr.ScenePlayer) {
	balls, balltypes := FindTarget_SemiCircle(tick, player)
	for i := 0; i < len(balls) && i < len(balltypes); i++ {
		NormalAttack(tick, player, balls[i], balltypes[i])
	}
}

// 群体攻击 - 圆形范围
func AOEAttack_Circle(tick *b3core.Tick, player *plr.ScenePlayer) {
	balls, balltypes := FindTarget_Circle(tick, player)
	for i := 0; i < len(balls) && i < len(balltypes); i++ {
		NormalAttack(tick, player, balls[i], balltypes[i])
	}
}

func NormalAttack(tick *b3core.Tick, player *plr.ScenePlayer, targetBall interfaces.IBall, targetType consts.BallKind) {
	var bHitFlag bool
	switch targetType {
	case consts.BallKind_Feed:
		bHitFlag = NormalAttack_Feed(tick, player, targetBall.(*bll.BallFeed))
	case consts.BallKind_Skill:
		bHitFlag = NormalAttack_BallSkill(tick, player, targetBall.(*bll.BallSkill))
	case consts.BallKind_Player:
		bHitFlag = NormalAttack_Player(tick, player, targetBall.(*bll.BallPlayer), true)
	default:
		bHitFlag = false
	}
	if bHitFlag {
		hits := tick.Blackboard.Get("hits", "", "").(map[uint64]int)
		if _, ok := hits[targetBall.GetID()]; ok {
			hits[targetBall.GetID()] = hits[targetBall.GetID()] + 1
		} else {
			hits[targetBall.GetID()] = 1
		}
	}
}

func NormalAttack_Feed(tick *b3core.Tick, player *plr.ScenePlayer, ball *bll.BallFeed) bool {
	feed := ball
	if IsCanAttack(tick, player, feed) {
		return playerHitFeed(player, feed)
	}
	return false
}

func NormalAttack_BallSkill(tick *b3core.Tick, player *plr.ScenePlayer, ball *bll.BallSkill) bool {
	ballskill := ball
	skill := ballskill.Skill.(*SkillBall)
	if skill.TryGetHit(player) {
		if IsCanAttack(tick, player, ballskill) {
			skill.GetHit(player)
			return true
		}
	}
	return false
}

func NormalAttack_Player(tick *b3core.Tick, player *plr.ScenePlayer, ball *bll.BallPlayer, bGetHit bool) bool {
	otherball := ball
	otherplayer := otherball.GetPlayer().(*plr.ScenePlayer)
	if otherplayer.CanBeEat() {
		if IsCanAttackPlayer(tick, player, otherball) {
			bHit := playerHitPlayer(player, otherplayer)
			if bHit && bGetHit {
				skillid := uint32(tick.Blackboard.Get("castskill", "", "").(float64))
				otherplayer.Skill.(*SkillPlayer).GetHit(player, skillid)
			}
			return bHit
		}
	}
	return false
}

func BallSkillAttack(tick *b3core.Tick, player *plr.ScenePlayer, ballskill *bll.BallSkill, attackScale float64, iball interfaces.IBall) bool {
	distance := iball.GetPosPtr().SqrMagnitudeTo(ballskill.GetPosPtr())
	tmp := iball.GetRect().Radius + float64(ballskill.GetRadius()) + attackScale
	if float64(distance) <= tmp*tmp {
		if iball.GetBallType() == usercmd.BallType_Player {
			targetball := iball.(*bll.BallPlayer)
			target := targetball.GetPlayer().(*plr.ScenePlayer)
			return playerHitPlayer(player, target)
		} else if iball.GetBallType() > usercmd.BallType_FeedBegin && iball.GetBallType() < usercmd.BallType_FeedEnd {
			return playerHitFeed(player, iball.(*bll.BallFeed))
		}
	}
	return false
}

func playerHitFeed(player *plr.ScenePlayer, feed *bll.BallFeed) bool {
	//TODO wei:
	// x, y := feed.GetPos()
	// _, ok := player.GetScene().GetCell(x, y)
	// if ok {
	// 	player.Eat(&feed.BallFood)
	// 	feedid := feed.GetID()

	// 	if feed.GetBirthPoint() != nil {
	// 		feed.GetBirthPoint().OnChildRemove(feed)
	// 	}

	// 	player.GetScene().RemoveFeed(feed)

	// 	//cell.Remove(feedid, feed.GetBallType())
	// 	player.AddEatMsg(player.GetID(), feedid)

	// 	return true
	// }
	return false
}

func playerHitPlayer(player *plr.ScenePlayer, target *plr.ScenePlayer) bool {
	targetball := &target.BallPlayer
	damage, bHit := player.Hit(targetball)
	if bHit {
		target.AddHitMsg(player.GetID(), targetball.GetID(), -damage, uint32(targetball.GetAttr(bll.AttrHP)), player.GetScene())
		return true
	}
	return false
}
