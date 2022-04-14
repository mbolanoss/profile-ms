package models

import (
	"time"

	"github.com/kamva/mgm/v3"
)


type SongInfo struct {
	Reproductions int `json:"reproductions" bson:"reproductions"`
	LastUpdate time.Time `json:"lastUpdate" bson:"lastUpdate"`
	ArtistName string `json:"artistName" bson:"artistName"`
}
type PlayedSongsList struct {
	mgm.DefaultModel `bson:",inline"`

	Username string `json:"username" bson:"username"`
	// Key: songID
	Songs map[int]SongInfo `json:"songsList" bson:"songsList"`
}

func NewPlayedSongsList (username string) *PlayedSongsList {
	return &PlayedSongsList{
		Username: username,
		Songs: make(map[int]SongInfo),
	}
}