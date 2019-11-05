package main

import (
	"client-call-serviceA/src/services/serviceA"
	"client-call-serviceA/src/services/servicetype"

	"gitlab.ztgame.com/tech/public/go-service/zeus/framework/service"
)

// regAllServices 注册所有的逻辑服务
func regAllServices() {
	service.RegService(servicetype.ServiceTypeGateway, "serviceA", &serviceA.ServiceA{})
}
