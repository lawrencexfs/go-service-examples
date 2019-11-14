package rank

import (
	"battleservice/src/services/battle/scene/plr"
	"battleservice/src/services/battle/usercmd"
	"sort"
	"time"
)

const (
	MaxTopPlayer int = 9 //排行人数
)

// BroadcastTopList 发送排行榜.
// endTime 结束时间, players 玩家对象
func BroadcastTopList(endTime int64, iscene plr.IScene) {
	msgTop := &usercmd.MsgTop{}
	ltime := endTime - time.Now().Unix()
	if ltime > 0 {
		msgTop.EndTime = uint32(ltime)
	} else {
		msgTop.EndTime = 0
	}
	msgTop.Players = []*usercmd.MsgPlayer{}
	tmpList := _FreeRankList{}

	iscene.TravsalPlayers(func(p *plr.ScenePlayer) {
		tmpList = append(tmpList, _FreeRank{PlayerID: p.GetEntityID(), Score: float64(p.GetExp())})
	})

	sort.Sort(tmpList)

	var topplayers [MaxTopPlayer]usercmd.MsgPlayer // 排行榜
	for index, v := range tmpList {
		p := iscene.GetPlayer(v.PlayerID)
		if p == nil {
			continue
		}
		if len(msgTop.Players) < MaxTopPlayer {
			topplayers[index].Id = uint64(v.PlayerID)
			topplayers[index].Name = p.Name
			topplayers[index].Curexp = uint32(v.Score)
			msgTop.Players = append(msgTop.Players, &topplayers[index])
		}
	}
	for k, v := range tmpList {
		p := iscene.GetPlayer(v.PlayerID)
		if p == nil {
			continue
		}

		p.Rank = uint32(k + 1)
		msgTop.Rank = p.Rank
		msgTop.KillNum = uint32(v.Score)
		p.Send(msgTop)
	}
}
