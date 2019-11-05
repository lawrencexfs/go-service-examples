package main

import (
	_ "net/http/pprof"

	"gitlab.ztgame.com/tech/public/go-service/zeus/framework/app"
)

func main() {
	// 注册所有服务
	regAllServices()

	app.Run("")
}
