package entitydef

type HEROINFO struct {
	HeroName	string `bson:"HeroName"` 
	HeroID	int32 `bson:"HeroID"` 
	HeroProficiencyLv	int32 `bson:"HeroProficiencyLv"` 
	HeroProficiencyExp	int32 `bson:"HeroProficiencyExp"` 
	HeroProficiencyUnlockSkill	int32 `bson:"HeroProficiencyUnlockSkill"` 
	HeroSkinId	int32 `bson:"HeroSkinId"` 
}
type HEROS = map[string]HEROINFO

type FRIENDS struct {
	MyFriendsDbid	[]uint64 `bson:"MyFriendsDbid"` 
	ApplyFriendsDbid	[]int32 `bson:"ApplyFriendsDbid"` 
	MyFriendsName	string `bson:"MyFriendsName"` 
}
