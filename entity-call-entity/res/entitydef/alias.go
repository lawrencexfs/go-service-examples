package entitydef

type HEROINFO struct {
	HeroProficiencyExp	int32 `bson:"HeroProficiencyExp"` 
	HeroProficiencyUnlockSkill	int32 `bson:"HeroProficiencyUnlockSkill"` 
	HeroSkinId	int32 `bson:"HeroSkinId"` 
	HeroName	string `bson:"HeroName"` 
	HeroID	int32 `bson:"HeroID"` 
	HeroProficiencyLv	int32 `bson:"HeroProficiencyLv"` 
}
type HEROS = map[string]HEROINFO
type FRIENDS struct {
	ApplyFriendsDbid	[]int32 `bson:"ApplyFriendsDbid"` 
	MyFriendsName	string `bson:"MyFriendsName"` 
	MyFriendsDbid	[]uint64 `bson:"MyFriendsDbid"` 
}
