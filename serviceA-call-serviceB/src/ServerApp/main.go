package main

import (
	_ "net/http/pprof"

	"github.com/GA-TECH-SERVER/zeus/framework/app"
)

func main() {
	// 注册所有服务
	regAllServices()

	app.Run("")
}
