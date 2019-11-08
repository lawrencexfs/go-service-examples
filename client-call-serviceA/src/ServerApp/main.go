package main

import (
	"github.com/giant-tech/go-service/framework/app"
)

func main() {
	// 注册所有服务
	regAllServices()

	app.Run("")
}
