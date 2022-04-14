package dtos

type AddReproductionDto struct {
	Username string `json:"username" bson:"username"`
	SongId int `json:"songId" bson:"songId"`
}

func NewAddReproductionDto(username string, songId int) AddReproductionDto {
	return AddReproductionDto{
		Username: username,
		SongId: songId,
	}
}