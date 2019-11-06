package main

import (
	"serviceA-call-serviceB/src/services/serviceA"
	"serviceA-call-serviceB/src/services/serviceB"
	"serviceA-call-serviceB/src/services/servicetype"

	"github.com/giant-tech/go-service/framework/service"
)

// regAllServices 注册所有的逻辑服务
func regAllServices() {
	service.RegService(servicetype.ServiceTypeA, "serviceA", &serviceA.ServiceA{})
	service.RegService(servicetype.ServiceTypeB, "serviceB", &serviceB.ServiceB{})
}
