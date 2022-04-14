package handlers

import (
	"net/http"
	"profile-ms/dtos"
	"profile-ms/helpers"
	"profile-ms/models"

	"github.com/gofiber/fiber/v2"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
)

func SetupPlayedArtistsRoutes(app *fiber.App) {
	// Used to get the top played artists
	app.Get("/playedArtists", GetTopPlayedArtists)
}

func GetTopPlayedArtists(ctx *fiber.Ctx) error {
	
	var err error

	topPlayedSongsDto := new(dtos.TopPlayedDto)
	err = ctx.BodyParser(topPlayedSongsDto)
	if err != nil {
		ctx.Status(http.StatusInternalServerError).SendString("Error parsing top played songs body")
	}

	var playedSongsList models.PlayedSongsList

	err = mgm.Coll(&playedSongsList).First(bson.M{"username" : topPlayedSongsDto.Username}, &playedSongsList)

	if err != nil {
		return ctx.Status(http.StatusBadRequest).SendString("Username does not exist")
	}
	
	playedArtists := make(map[string]int)

	for _, songInfo := range playedSongsList.Songs {
		playedArtists[songInfo.ArtistName] += songInfo.Reproductions
	}

	// Sorting
	sorted := helpers.SortMapString_Int(playedArtists)

	const NUMBER_OF_ARTISTS int = 5 

	if len(sorted) > NUMBER_OF_ARTISTS {
		sorted = sorted[len(playedArtists) - NUMBER_OF_ARTISTS:]
	}
	
	// Key: songID - Value: number of reproductions
	return ctx.JSON(fiber.Map{
		"songs" : sorted,
	})

	/* for artist, times := range topPlayedArtists {
		fmt.Println(artist, times)
	} */
}