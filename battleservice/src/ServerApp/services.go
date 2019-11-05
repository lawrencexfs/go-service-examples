package main

import (
	"battleservice/src/services/battle"
	"battleservice/src/services/servicetype"

	"gitlab.ztgame.com/tech/public/go-service/zeus/framework/service"
)

// regAllServices 注册所有的逻辑服务
func regAllServices() {
	service.RegService(servicetype.ServiceTypeBattle, "battle", &battle.BattleService{})
	//service.RegService(servicetype.ServiceTypeMatchClient, "matchclient", &matchclient.MatchClientService{})
}
