package dtos

type TopPlayedSongsDto struct {
	Username string `json:"username" bson:"username"`
	Gap int `json:"gap" bson:"gap"`
}