package mf

import (
	"time"

	"github.com/cihub/seelog"
	"github.com/giant-tech/go-service/logic/matchbase/matchdata"
	"github.com/giant-tech/go-service/logic/matchbase/matchitf"
)

// MatchFunction 自定义匹配函数
type MatchFunction struct {
}

// MatchForResult 面向结果匹配
func (mfunc *MatchFunction) MatchForResult(pool matchitf.IMatchPool) []*matchdata.MatchResult {
	seelog.Debug("MatchForResult")

	var tempResultSlice []*matchdata.MatchResult

	var count uint32
	result := &matchdata.MatchResult{}

	pool.Range(func(m *matchdata.Matcher) bool {
		result.Matchers = append(result.Matchers, m)
		count += m.Num

		//满6个就匹配成功了
		if count >= 6 {
			tempResultSlice = append(tempResultSlice, result)
			result = &matchdata.MatchResult{}

			count = 0
		}

		return true
	})

	return tempResultSlice
}

// MatchForProgress 面向过程匹配
func (mfunc *MatchFunction) MatchForProgress(matcher *matchdata.Matcher, room matchitf.IMatchPool) (*matchdata.MatchResult, bool) {
	seelog.Debug("MatchForResult")

	// matcher如果为nil，判断room是否超时
	if matcher == nil {
		// 超过60秒直接开始
		if time.Now().Unix()-room.GetCreateTime() > 60 {
			result := &matchdata.MatchResult{}

			room.Range(func(m *matchdata.Matcher) bool {
				result.Matchers = append(result.Matchers, m)
				return true
			})

			return result, false
		}

		return nil, false
	}

	var count uint32
	room.Range(func(m *matchdata.Matcher) bool {
		count += m.Num
		return true
	})

	// 人数够了
	if count+matcher.Num >= 6 {
		result := &matchdata.MatchResult{}

		room.Range(func(m *matchdata.Matcher) bool {
			result.Matchers = append(result.Matchers, m)
			return true
		})

		result.Matchers = append(result.Matchers, matcher)

		return result, true
	}

	return nil, true
}
