package models

import "github.com/kamva/mgm/v3"

type M map[string]interface{}

type UserConfig struct {
	mgm.DefaultModel `bson:",inline"`

	AutoplayOn bool `json:"autoplayOn" bson:"autoplayOn"`
	DownloadRoute string `json:"downloadRoute" bson:"downloadRoute"`
	PreferredColor string `json:"preferredColor" bson:"preferredColor"`
}

func NewUserConfig(autoplay bool, newDownloadRoute string, newPreferredColor string) *UserConfig{
	return &UserConfig{
		AutoplayOn: autoplay,
		DownloadRoute: newDownloadRoute,
		PreferredColor: newPreferredColor,
	}
}