package mf

import (
	"github.com/cihub/seelog"
	"gitlab.ztgame.com/tech/public/go-service/zeus/logic/matchbase/matchdata"
	"gitlab.ztgame.com/tech/public/go-service/zeus/logic/matchbase/matchitf"
)

// MatchNotify 自定义通知函数
type MatchNotify struct {
}

// MatchFinishNotify 匹配成功通知
func (mfunc *MatchNotify) MatchFinishNotify(result *matchdata.MatchResult) {
	seelog.Debug("MatchFinishNotify")

	//TODO: 通知房间服创建新房间
}

//MatcherJoinNotify 玩家进入，用于需要过程的匹配
func (mfunc *MatchNotify) MatcherJoinNotify(matcher *matchdata.Matcher, room matchitf.IMatchPool) {
	seelog.Debug("MatcherJoinNotify")
}

//MatcherLeaveNotify 玩家退出，用于需要过程的匹配
func (mfunc *MatchNotify) MatcherLeaveNotify(matcher *matchdata.Matcher, room matchitf.IMatchPool) {
	seelog.Debug("MatcherLeaveNotify")
}
