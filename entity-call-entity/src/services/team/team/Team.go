package team

import (
	"entity-call-entity/src/services/team/teamdata"
	"fmt"

	log "github.com/cihub/seelog"
	"github.com/giant-tech/go-service/framework/entity"
)

// Team 队伍
type Team struct {
	entity.GroupEntity
}

// OnInit 初始化
func (t *Team) OnInit(initData interface{}) error {
	teamData, ok := initData.(*teamdata.CreateTeamData)
	if !ok {
		log.Error("initData is not teamdata.CreateTeamData")
		return fmt.Errorf("initData is not teamdata.CreateTeamData")
	}

	log.Debug("TeamName: ", teamData.TeamName)
	t.GroupEntity.OnGroupInit()
	return nil
}

// OnLoop 每帧调用
func (t *Team) OnLoop() {
	log.Debug("Team.OnLoop")
	t.GroupEntity.OnGroupLoop()
}

// OnDestroy 销毁
func (t *Team) OnDestroy() {
	log.Debug("OnDestroy")
	t.GroupEntity.OnGroupDestroy()
}
