package models

import "github.com/kamva/mgm/v3"

type MainSongsList struct {
	mgm.DefaultModel `bson:",inline"`

	Username string `json:"username" bson:"username"`
	// Key: songID - Value: number of reproductions
	Songs map[int]int `json:"songsList" bson:"songsList"`
}

func NewMainSongsList (username string) *MainSongsList {
	return &MainSongsList{
		Username: username,
		Songs: make(map[int]int),
	}
}