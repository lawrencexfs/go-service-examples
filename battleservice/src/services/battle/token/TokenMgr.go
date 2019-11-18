package token

import (
	"sync"
)

// 管理所有token.
// Thread-safe.
type TokenMgr struct {
	tokenToUserMap *sync.Map // map from token to *userInfo
}

type userInfo struct {
	roomID   uint64
	playerID uint64
}

func (t *TokenMgr) Init() {
	t.tokenToUserMap = &sync.Map{}
}

func (t *TokenMgr) InsertToken(token string, roomID uint64, playerID uint64) {
	userInfo := &userInfo{
		roomID:   roomID,
		playerID: playerID,
	}
	t.tokenToUserMap.Store(token, userInfo)
}

func (t *TokenMgr) LookupToken(token string) (bool, uint64, uint64) {
	info, ok := t.tokenToUserMap.Load(token)
	if !ok {
		return ok, 0, 0
	}
	u := info.(*userInfo)
	return ok, u.roomID, u.playerID
}

func (t *TokenMgr) RemoveToken(token string) {
	t.tokenToUserMap.Delete(token)
}
