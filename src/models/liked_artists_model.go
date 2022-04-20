package models

import "github.com/kamva/mgm/v3"

type LikedArtistList struct {
	mgm.DefaultModel `bson:",inline"`

	Username string `json:"username" bson:"username"`
	Artists []string `json:"artists" bson:"artists"`
}

func NewLikedArtistList(username string) *LikedArtistList {
	return &LikedArtistList{
		Username: username,
		Artists: make([]string, 0),
	}
}

func DeleteArtist(artistsList []string, index int) []string{
	if index == 0 {
		return artistsList[1:]
	} else if index == len(artistsList) - 1 {
		length := len(artistsList)
		return artistsList[:length-1]
	} else {
		return append(artistsList[:index], artistsList[index+1:]...)
	}
} 