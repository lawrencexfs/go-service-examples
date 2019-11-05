package main

import (
	"matchservice/src/services/match"
	"matchservice/src/services/matchclient"
	"matchservice/src/services/servicetype"

	"gitlab.ztgame.com/tech/public/go-service/zeus/framework/service"
)

// regAllServices 注册所有的逻辑服务
func regAllServices() {
	service.RegService(servicetype.ServiceTypeMatch, "match", &match.MatchService{})
	service.RegService(servicetype.ServiceTypeMatchClient, "matchclient", &matchclient.MatchClientService{})
}
