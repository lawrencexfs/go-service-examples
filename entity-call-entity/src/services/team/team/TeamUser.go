package team

import (
	"fmt"

	log "github.com/cihub/seelog"
	"gitlab.ztgame.com/tech/public/go-service/zeus/framework/entity"

	"entity-call-entity/src/services/team/teamdata"
)

// TeamUser 玩家
type TeamUser struct {
	entity.Entity
}

func (tu *TeamUser) OnInit(initData interface{}) error {
	playerInfo, ok := initData.(*teamdata.TeamPlayerInfo)
	if !ok {
		return fmt.Errorf("initData is not TeamPlayerInfo")
	}

	log.Debug("PlayerName: ", playerInfo.PlayerName)

	return nil
}

// OnLoop 每帧调用
func (tu *TeamUser) OnLoop() {
	log.Debug("TeamUser.OnLoop")
}

func (tu *TeamUser) OnDestroy() {
	log.Debug("OnDestroy")
}
