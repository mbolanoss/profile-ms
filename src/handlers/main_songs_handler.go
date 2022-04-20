package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"

	"profile-ms/dtos"
	"profile-ms/helpers"
	"profile-ms/models"
)

func SetupMainSongsRoutes(app *fiber.App){
	// Used to create the list in the DB
	app.Post("/playedSongs", CreatePlayedSongsList)

	// Used to add 1 to the number of reproductions of a song
	app.Put("/playedSongs", AddReproduction)

	// Used to get the top played songs
	app.Get("/playedSongs", GetTopPlayedSongs)
}

func CreatePlayedSongsList(ctx *fiber.Ctx) error {
	username := ctx.Query("username")

	if username == "" {
		return ctx.Status(http.StatusBadRequest).SendString("No username query found in url")
	}

	playedSongsList := models.NewPlayedSongsList(username)

	mgm.Coll(playedSongsList).InsertOne(context.TODO(), playedSongsList)

	return ctx.SendStatus(http.StatusOK)
}

func AddReproduction(ctx *fiber.Ctx) error {
	var err error

	addReproductionDto := new(dtos.AddReproductionDto)
	err = ctx.BodyParser(addReproductionDto)
	if err != nil {
		ctx.Status(http.StatusInternalServerError).SendString("Error parsing add reproduction body")
	}

	var playedSongsList models.PlayedSongsList

	err = mgm.Coll(&playedSongsList).First(bson.M{"username" : addReproductionDto.Username}, &playedSongsList)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).SendString("Username does not exist")
	}

	// Adding 1 to reproductions counter and replacing update date
	currentSong := playedSongsList.Songs[addReproductionDto.SongId]
	currentSong.Reproductions += 1
	currentSong.ArtistName = addReproductionDto.ArtistName
	currentSong.LastUpdate = time.Now()
	playedSongsList.Songs[addReproductionDto.SongId] = currentSong

	err = mgm.Coll(&playedSongsList).Update(&playedSongsList)

	if err != nil{
		return ctx.Status(http.StatusInternalServerError).SendString("Error while updating song's number of reproductions in the DB")
	}

	return ctx.SendStatus(http.StatusOK)
}

func GetTopPlayedSongs(ctx *fiber.Ctx) error {
	var err error

	topPlayedDto := new(dtos.TopPlayedDto)
	err = ctx.BodyParser(topPlayedDto)
	if err != nil {
		ctx.Status(http.StatusInternalServerError).SendString("Error parsing top played songs body")
	}

	var playedSongsList models.PlayedSongsList

	err = mgm.Coll(&playedSongsList).First(bson.M{"username" : topPlayedDto.Username}, &playedSongsList)

	if err != nil {
		return ctx.Status(http.StatusBadRequest).SendString("Username does not exist")
	}

	//Time gap filtering
	filteredSongs := make(map[int]models.SongInfo)

	for songId, reproduction := range playedSongsList.Songs {
		lastUpdate := reproduction.LastUpdate
		newDate :=lastUpdate.AddDate(0, 0, topPlayedDto.Gap)

		if(newDate.After(time.Now())){
			filteredSongs[songId] = reproduction
		}
	}

	// Sorting
	songs := make(map[int]int)
	for key, value := range filteredSongs {
		songs[key] = value.Reproductions
	}
	sorted := helpers.SortMapInt_Int(songs)

	const NUMBER_OF_SONGS int = 5 

	if len(sorted) > NUMBER_OF_SONGS {
		sorted = sorted[len(filteredSongs) - NUMBER_OF_SONGS:]
	}
	
	// Key: songID - Value: number of reproductions
	return ctx.JSON(fiber.Map{
		"songs" : sorted,
	})
}