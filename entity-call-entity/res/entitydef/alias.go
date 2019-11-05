package entitydef

type FRIENDS struct {
	MyFriendsName	string `bson:"MyFriendsName"` 
	MyFriendsDbid	[]uint64 `bson:"MyFriendsDbid"` 
	ApplyFriendsDbid	[]int32 `bson:"ApplyFriendsDbid"` 
}
type HEROINFO struct {
	HeroName	string `bson:"HeroName"` 
	HeroID	int32 `bson:"HeroID"` 
	HeroProficiencyLv	int32 `bson:"HeroProficiencyLv"` 
	HeroProficiencyExp	int32 `bson:"HeroProficiencyExp"` 
	HeroProficiencyUnlockSkill	int32 `bson:"HeroProficiencyUnlockSkill"` 
	HeroSkinId	int32 `bson:"HeroSkinId"` 
}
type HEROS = map[string]HEROINFO

