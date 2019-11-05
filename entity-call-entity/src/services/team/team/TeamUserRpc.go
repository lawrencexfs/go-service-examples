package team

import (
	log "github.com/cihub/seelog"
)

// RPCLeaveTeam 离开队伍
func (tu *TeamUser) RPCLeaveTeam() {
	log.Debug("RPCLeaveTeam, EntityID: ", tu.GetEntityID(), ", GroupID: ", tu.GetGroupID())

	groupID := tu.GetGroupID()

	//删除队伍成员
	tu.GetIEntities().DestroyEntity(tu.GetEntityID())

	//删除队伍
	tu.GetIEntities().GetLocalService().DestroyEntity(groupID)
}
