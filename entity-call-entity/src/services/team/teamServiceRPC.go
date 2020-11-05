package team

import (
	"entity-call-entity/src/services/servicetype"
	"entity-call-entity/src/services/team/teamdata"

	log "github.com/cihub/seelog"
	"github.com/giant-tech/go-service/framework/entity"
	"github.com/giant-tech/go-service/framework/iserver"
)

// RPCCreateTeam 创建队伍
func (ts *TeamService) RPCCreateTeam(teamData *teamdata.CreateTeamData) {
	log.Debug("TeamService::RPCCreateTeam: ")

	//创建队伍
	e, err := ts.CreateEntity("Team", 0, teamData, true, 0)
	if err != nil {
		log.Error("create team failed, err: ", err)
	}

	teamID := 0
	group, ok := e.GetRealPtr().(iserver.IEntityGroup)
	if !ok {
		log.Error("team is not GroupEntity")
		ts.DestroyEntity(e.GetEntityID())
		entity.NewEntityProxy(teamData.PlayerInfo.PlayerID).AsyncCall(servicetype.ServiceTypeGateway, "CreateTeamResult", teamID)
		return
	}

	//创建队伍成员
	teamPlayer, err := group.CreateEntityWithID("Player", teamData.PlayerInfo.PlayerID, group.GetEntityID(), teamData.PlayerInfo, true, 0)
	if err != nil {
		log.Error("Create team player failed, err: ", err)
		ts.DestroyEntity(e.GetEntityID())
		entity.NewEntityProxy(teamData.PlayerInfo.PlayerID).AsyncCall(servicetype.ServiceTypeGateway, "CreateTeamResult", teamID)

		return
	}

	//通知玩家队伍创建成功
	teamPlayer.AsyncCall(servicetype.ServiceTypeGateway, "CreateTeamResult", e.GetEntityID())
}
