package main

import (
	"entity-call-entity/src/Client/clt"
	"os"
	"os/signal"
	"syscall"
	"time"

	log "github.com/cihub/seelog"
	"gitlab.ztgame.com/tech/public/go-service/zeus/framework/idata"
)

func main() {

	cl := &clt.Client{}
	cl.ActC = make(chan func(), 1024)
	//cl.newClient()
	cl.Init("127.0.0.1:17000")
	cl.RegRPCMsg(&clt.ServiceAUserRpcProc{Cli: cl})

	cl.Login()
	log.Debug("send rpc hello")
	cl.AsyncCall(idata.ServiceGateway, "Hello", "hello world", int32(1))

	//创建房间
	cl.AsyncCall(idata.ServiceGateway, "CreateTeam", "TestTeam")

	//往serviceA发送玩家属性改变消息
	//cl.AsyncCall(idata.ServiceGateway, "ModifyAttr", "change attr", int32(1), int32(2))
	cl.Run()
	waitForExit() // Wait for exit

}

// waitForExit 等待退出
func waitForExit() {
	//tmode := viper.GetString("test-mode")
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)
	ticker := time.NewTicker(time.Duration(10) * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-c:
			return
		case <-ticker.C:
		}
	}
}
