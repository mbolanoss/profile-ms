package dtos

type AddReproductionDto struct {
	SongId int `json:"songId" bson:"songId"`
	Username string `json:"username" bson:"username"`
	ArtistName string `json:"artistName" bson:"artistName"`
}