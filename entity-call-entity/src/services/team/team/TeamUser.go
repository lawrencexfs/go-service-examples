package team

import (
	"fmt"

	log "github.com/cihub/seelog"
	"github.com/giant-tech/go-service/framework/entity"

	"entity-call-entity/src/entitydef"
	"entity-call-entity/src/services/team/teamdata"
)

// TeamUser 玩家
type TeamUser struct {
	entity.Entity
	entitydef.PlayerDef
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
	log.Debug("TeamUser.OnLoop,level=", tu.Getlevel())
}

func (tu *TeamUser) OnDestroy() {
	log.Debug("OnDestroy")
}
