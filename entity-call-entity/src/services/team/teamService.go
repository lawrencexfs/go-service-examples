package team

import (
	"github.com/giant-tech/go-service/framework/service"

	log "github.com/cihub/seelog"

	"entity-call-entity/src/services/team/team"

	"github.com/giant-tech/go-service/framework/idata"
)

// TeamService 队伍服务
type TeamService struct {
	service.BaseService
}

// OnInit 初始化
func (ts *TeamService) OnInit() error {
	log.Debug("ServiceBService.OnInit")

	ts.RegProtoType("Player", &team.TeamUser{}, false)
	ts.RegProtoType("Team", &team.Team{}, false)

	return nil
}

// OnTick tick
func (ts *TeamService) OnTick() {
	ts.Loop()
}

// OnDestroy 析构
func (ts *TeamService) OnDestroy() {
	log.Debug("ServiceBService.OnDestroy")
	ts.Destroy()
}

// OnDisconnected 服务断开链接
func (ts *TeamService) OnDisconnected(infovec []*idata.ServiceInfo) {

	log.Info("team OnDisconnected, infovec = ", infovec)
	for _, s := range infovec {
		log.Info("team OnDisconnected, info= ", *s)
	}
}

// OnConnected 和别的服务建立链接
func (ts *TeamService) OnConnected(infovec []*idata.ServiceInfo) {

}
