package matchclient

import (
	"github.com/cihub/seelog"
	"github.com/giant-tech/go-service/logic/matchbase/matchdata"
)

// RPCMatchResult 匹配结果
func (mc *MatchClientService) RPCMatchResult(result *matchdata.MatchResult) {
	seelog.Debug("RPCMatchResult")
}
