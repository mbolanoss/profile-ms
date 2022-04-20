package models

import "github.com/kamva/mgm/v3"

type UserConfig struct {
	mgm.DefaultModel `bson:",inline"`

	Username string `json:"username" bson:"username"`
	AutoplayOn bool `json:"autoplayOn" bson:"autoplayOn"`
	DownloadRoute string `json:"downloadRoute" bson:"downloadRoute"`
	PreferredColor string `json:"preferredColor" bson:"preferredColor"`
}

func NewUserConfig(username string ,autoplay bool, newDownloadRoute string, newPreferredColor string) *UserConfig{
	return &UserConfig{
		Username: username,
		AutoplayOn: autoplay,
		DownloadRoute: newDownloadRoute,
		PreferredColor: newPreferredColor,
	}
}