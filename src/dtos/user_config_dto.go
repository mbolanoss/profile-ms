package dtos

type UserConfigDto struct {

	Username string `json:"username" bson:"username"`
	AutoplayOn bool `json:"autoplayOn" bson:"autoplayOn"`
	DownloadRoute string `json:"downloadRoute" bson:"downloadRoute"`
	PreferredColor string `json:"preferredColor" bson:"preferredColor"`
}

func NewUserConfig(username string ,autoplay bool, newDownloadRoute string, newPreferredColor string) UserConfigDto{
	return UserConfigDto{
		Username: username,
		AutoplayOn: autoplay,
		DownloadRoute: newDownloadRoute,
		PreferredColor: newPreferredColor,
	}
}