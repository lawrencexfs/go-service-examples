package matchclient

import (
	"github.com/cihub/seelog"
	"gitlab.ztgame.com/tech/public/go-service/zeus/logic/matchbase/matchdata"
)

// RPCMatchResult 匹配结果
func (mc *MatchClientService) RPCMatchResult(result *matchdata.MatchResult) {
	seelog.Debug("RPCMatchResult")
}
