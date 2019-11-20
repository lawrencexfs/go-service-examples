package itf

import (
	"github.com/giant-tech/go-service/framework/iserver"
)

// IBattleService battle服接口
type IBattleService interface {
	iserver.IServiceBase

	//验证token
	LookupToken(token string) (bool, uint64, uint64)
}
