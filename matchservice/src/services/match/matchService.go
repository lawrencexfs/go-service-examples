package match

import (
	"matchservice/src/services/match/mf"

	"github.com/cihub/seelog"
	"github.com/giant-tech/go-service/framework/service"
	"github.com/giant-tech/go-service/logic/matchbase"
)

// MatchService 队伍服务
type MatchService struct {
	service.BaseService
	matchbase.MatchBase
}

// OnInit 初始化
func (ms *MatchService) OnInit() error {
	seelog.Debug("MatchService.OnInit")

	//MatchBase初始化，设置回调函数
	ms.MatchBase.Init(&mf.MatchFunction{}, &mf.MatchNotify{})

	return nil
}

// OnTick tick
func (ms *MatchService) OnTick() {
	//ms.TryToMatch()
}

// OnDestroy 析构
func (ms *MatchService) OnDestroy() {
	seelog.Debug("MatchService.OnDestroy")

}
