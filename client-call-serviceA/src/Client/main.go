package main

import (
	"client-call-serviceA/src/Client/clt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"client-call-serviceA/src/services/servicetype"

	log "github.com/cihub/seelog"
	"github.com/giant-tech/go-service/base/zlog"
)

func main() {

	// 设置Seelog
	zlog.InitDefault()
	cl := &clt.Client{}
	cl.ActC = make(chan func(), 1024)
	//cl.newClient()
	cl.Init("127.0.0.1:17000")
	cl.RegRPCMsg(&clt.ServiceAUserRpcProc{Cli: cl})

	cl.Login()
	log.Debug("send rpc hello")
	cl.AsyncCall(servicetype.ServiceTypeGateway, "Hello", "hello world", int32(1))

	cl.Run()
	waitForExit() // Wait for exit

}

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
