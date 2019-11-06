package token

import (
	"battleservice/src/services/battle/types"
	"sync"
)

// 管理所有token.
// Thread-safe.
type TokenMgr struct {
	tokenToUserMap *sync.Map // map from token to *userInfo
}

type Token = types.Token
type userInfo struct {
	roomID   uint64
	playerID uint64
}

func (t *TokenMgr) Init() {
	t.tokenToUserMap = &sync.Map{}
}

func (t *TokenMgr) InsertToken(token Token, roomID uint64, playerID uint64) {
	userInfo := &userInfo{
		roomID:   roomID,
		playerID: playerID,
	}
	t.tokenToUserMap.Store(token, userInfo)
}

func (t *TokenMgr) LookupToken(token Token) (bool, uint64, uint64) {
	info, ok := t.tokenToUserMap.Load(token)
	if !ok {
		return ok, 0, 0
	}
	u := info.(*userInfo)
	return ok, u.roomID, u.playerID
}

func (t *TokenMgr) RemoveToken(token Token) {
	t.tokenToUserMap.Delete(token)
}
