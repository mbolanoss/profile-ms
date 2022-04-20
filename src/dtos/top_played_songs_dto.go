package dtos

type TopPlayedDto struct {
	Username string `json:"username" bson:"username"`
	Gap int `json:"gap" bson:"gap"`
}