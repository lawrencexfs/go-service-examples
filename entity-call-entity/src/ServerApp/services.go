package main

import (
	"entity-call-entity/src/services/gateway"
	"entity-call-entity/src/services/servicetype"
	"entity-call-entity/src/services/team"

	"gitlab.ztgame.com/tech/public/go-service/zeus/framework/service"
)

// regAllServices 注册所有的逻辑服务
func regAllServices() {
	service.RegService(servicetype.ServiceTypeGateway, "gateway", &gateway.GatewayService{})
	service.RegService(servicetype.ServiceTypeTeam, "team", &team.TeamService{})
}
