package teamdata

// CreateTeamData 创建队伍的数据
type CreateTeamData struct {
	TeamName   string
	PlayerInfo *TeamPlayerInfo
}

// TeamPlayerInfo 队伍成员信息
type TeamPlayerInfo struct {
	PlayerID   uint64
	PlayerName string
}
