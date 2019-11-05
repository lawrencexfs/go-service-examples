package matchclient

import (
	"matchservice/src/services/servicetype"
	"strconv"

	"gitlab.ztgame.com/tech/public/go-service/zeus/framework/iserver"
	"gitlab.ztgame.com/tech/public/go-service/zeus/framework/service"
	"gitlab.ztgame.com/tech/public/go-service/zeus/logic/matchbase/matchdata"

	log "github.com/cihub/seelog"
)

// MatchClientService 匹配客户端
type MatchClientService struct {
	service.BaseService
	count int64
}

// OnInit 初始化
func (mc *MatchClientService) OnInit() error {
	log.Debug("MatchClientService.OnInit")

	return nil
}

// OnTick tick
func (mc *MatchClientService) OnTick() {
	//通过自定义函数删选所需的服务
	proxy := iserver.GetServiceProxyMgr().GetRandService(servicetype.ServiceTypeMatch)
	if proxy.IsValid() {

		var result string

		//m1 := &matchdata.Matcher{Key: mc.getKey(), MatchMode: matchdata.MatchModeResult, MatchType: "3v3", Num: 1}
		//proxy.SyncCall(&result, "MatchReq", m1)
		//log.Debug("SyncCall result: ", result)

		key := mc.getKey()
		m2 := &matchdata.Matcher{Key: key, MatchMode: matchdata.MatchModeProgress, MatchType: "3v3", Num: 1}
		proxy.SyncCall(&result, "MatchReq", m2)
		proxy.AsyncCall("CancleMatchReq", key)
	}
}

// OnDestroy 析构
func (mc *MatchClientService) OnDestroy() {
	log.Debug("MatchClientService.OnDestroy")
}

func (mc *MatchClientService) getKey() string {
	mc.count++

	return strconv.FormatInt(mc.count, 10)
}
