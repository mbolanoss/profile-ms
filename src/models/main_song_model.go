package models

import (
	"time"

	"github.com/kamva/mgm/v3"
)


type Reproduction struct {
	Reproductions int `json:"reproductions" bson:"reproductions"`
	LastUpdate time.Time `json:"lastUpdate" bson:"LastUpdate"`
}
type MainSongsList struct {
	mgm.DefaultModel `bson:",inline"`

	Username string `json:"username" bson:"username"`
	// Key: songID - Value: number of reproductions
	Songs map[int]Reproduction `json:"songsList" bson:"songsList"`
}

func NewMainSongsList (username string) *MainSongsList {
	return &MainSongsList{
		Username: username,
		Songs: make(map[int]Reproduction),
	}
}