// 包 cll 处理场景中的格子(Cell).
package cll

// 场景中的Cell

import (
	"battleservice/src/services/base/util"
	"battleservice/src/services/battle/scene/bll"
	"battleservice/src/services/battle/scene/consts"
	"battleservice/src/services/battle/scene/interfaces"
	"battleservice/src/services/battle/usercmd"
	"fmt"

	"github.com/giant-tech/go-service/base/linmath"
)

var (
	CellWidth  float64 = 5 //格子像素宽
	CellHeight float64 = 5 //格子像素高
)

type Cell struct {
	id            int                        // Cell编号
	rect          util.Square                // Cell地图中的位置、区域
	Foods         map[uint64]*bll.BallFood   // 该Cell上的食物球
	playerballs   map[uint64]*bll.BallPlayer // 该Cell上的玩家球
	Feeds         map[uint64]*bll.BallFeed   // 该Cell上的动态障碍物
	Skills        map[uint64]*bll.BallSkill  // 该Cell上的技能球
	MsgMoves      []*usercmd.BallMove        // 语义上非必须，一些优化手段
	msgMovesMap   map[uint64]int             // 语义上非必须，一些优化手段
	msgRemovesMap map[uint64]bool            // 语义上非必须，一些优化手段
	msgAddsMap    map[uint64]bool            // 语义上非必须，一些优化手段
}

func NewCell(id int) *Cell {
	cell := Cell{id: id}
	cell.Clean()
	return &cell
}

func (cell *Cell) FindNearFood(player *bll.BallPlayer, pos *linmath.Vector3, ballType uint32, dir *linmath.Vector3) (*bll.BallFood, float32) {
	var min float32 = 10000
	var minball *bll.BallFood
	for _, ball := range cell.Foods {
		dis := ball.Pos.SqrMagnitudeTo(pos)
		if ballType == uint32(ball.BallType) || ballType == 0 {
			if dir != nil && linmath.IsSameDir(dir, ball.GetPosV(), player.GetPosV()) == false {
				continue
			}
			if player.PreCanEat(ball) {
				if minball == nil || dis < min {
					min = dis
					minball = ball
				}
			}
		}
	}
	return minball, min
}

func (cell *Cell) FindNearFeed(player *bll.BallPlayer, pos *linmath.Vector3, dir *linmath.Vector3) (*bll.BallFeed, float32) {
	var min float32
	var minball *bll.BallFeed
	for _, ball := range cell.Feeds {
		if dir != nil && linmath.IsSameDir(dir, ball.GetPosV(), player.GetPosV()) == false {
			continue
		}
		if player.PreCanEat(&ball.BallFood) {
			dis := ball.Pos.SqrMagnitudeTo(pos)
			if minball == nil || dis < min {
				min = dis
				minball = ball
			}
		}
	}
	return minball, min
}

func (cell *Cell) FindNearSkill(player *bll.BallPlayer, pos *linmath.Vector3, ballType uint32, dir *linmath.Vector3) (*bll.BallSkill, float32) {
	var min float32
	var minball *bll.BallSkill
	for _, ball := range cell.Skills {
		if (ballType == uint32(ball.BallType) || ballType == 0) && player.PreCanEat(&ball.BallFood) {
			if dir != nil && linmath.IsSameDir(dir, ball.GetPosV(), player.GetPosV()) == false {
				continue
			}
			dis := ball.Pos.SqrMagnitudeTo(pos)
			if minball == nil || dis < min {
				min = dis
				minball = ball
			}
		}
	}
	return minball, min
}

func (cell *Cell) FindNearBallByKind(player *bll.BallPlayer, pos *linmath.Vector3, kind consts.BallKind, dir *linmath.Vector3, ballType uint32) (interfaces.IBall, float32) {
	if kind == consts.BallKind_Food {
		if ball, dis := cell.FindNearFood(player, pos, ballType, dir); ball != nil {
			return ball, dis
		}
	} else if kind == consts.BallKind_Feed {
		if ball, dis := cell.FindNearFeed(player, pos, dir); ball != nil {
			return ball, dis
		}
	} else if kind == consts.BallKind_Skill {
		if ball, dis := cell.FindNearSkill(player, pos, ballType, dir); ball != nil {
			return ball, dis
		}
	}
	return nil, 10000
}

func (cell *Cell) Clean() {
	cell.Foods = make(map[uint64]*bll.BallFood)
	cell.playerballs = make(map[uint64]*bll.BallPlayer)
	cell.Feeds = make(map[uint64]*bll.BallFeed)
	cell.Skills = make(map[uint64]*bll.BallSkill)
	cell.ResetMsg()
}

func (cell *Cell) ResetMsg() {
	cell.MsgMoves = cell.MsgMoves[:0]
	cell.msgAddsMap = make(map[uint64]bool)
	cell.msgRemovesMap = make(map[uint64]bool)
	cell.msgMovesMap = make(map[uint64]int)
}

func (cell *Cell) AddMsgMove(ball interfaces.IBall) {
	x, y, _ := ball.GetPos()
	if msgIndex, ok := cell.msgMovesMap[ball.GetID()]; ok {
		msg := cell.MsgMoves[msgIndex]
		msg.X = int32(x * consts.MsgPosScaleRate)
		msg.Y = int32(y * consts.MsgPosScaleRate)
	} else {
		cell.MsgMoves = append(cell.MsgMoves,
			&usercmd.BallMove{
				Id: uint64(ball.GetID()),
				X:  int32(x * consts.MsgPosScaleRate),
				Y:  int32(y * consts.MsgPosScaleRate),
			})
		cell.msgMovesMap[ball.GetID()] = len(cell.MsgMoves) - 1
	}
}

func (cell *Cell) EatByPlayer(playerBall *bll.BallPlayer, player IScenePlayer) bool {
	isEat := false
	for _, food := range cell.Foods {
		if playerBall.CanEat(food) {
			playerBall.Eat(food)
			cell.OnFoodRemoved(food)
			//cell.Remove(food.GetID(), food.GetBallType())
			player.AddEatMsg(playerBall.GetID(), food.GetID())
			isEat = true
		}
	}
	return isEat
}

//移除食物事件
func (cell *Cell) OnFoodRemoved(food *bll.BallFood) {
	if food.GetBirthPoint() != nil {
		food.GetBirthPoint().OnChildRemove(food)
	}
}

//渲染
func (cell *Cell) Render(scene bll.IScene, per float64, now int64) {
	// var delSkills []*bll.BallSkill
	// for _, ball := range cell.Skills {
	// 	if ball.Skill.IsFinish() {
	// 		delSkills = append(delSkills, ball)
	// 	}
	// }
	// for _, skill := range delSkills {
	// 	cell.Remove(skill.GetID(), skill.GetBallType())
	// }
}

func (cell *Cell) string() string {
	return fmt.Sprintf("cell:%d%v", cell.id, cell.rect)
}

func (cell *Cell) ID() int {
	return cell.id
}
