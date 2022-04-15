package handlers

import (
	"net/http"
	"profile-ms/dtos"
	"profile-ms/helpers"
	"profile-ms/models"
	"time"

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
	
	playedArtists := make(map[string]models.SongInfo)

	//Gathering all artists reproductions according to the last update date
	for _, artistInfo := range playedSongsList.Songs {
		currentSongInfo := playedArtists[artistInfo.ArtistName]
		
		//Checking last update date
		lastUpdate := artistInfo.LastUpdate
		newDate :=lastUpdate.AddDate(0, 0, topPlayedDto.Gap)

		//Adding reproductions
		if(newDate.After(time.Now())){
			currentSongInfo.Reproductions += artistInfo.Reproductions
		}
		
		playedArtists[artistInfo.ArtistName] = currentSongInfo
	}

	// Sorting
	artists := make(map[string]int)
	for key, value := range playedArtists {
		artists[key] = value.Reproductions
	}
	filteredAndSortedArtists := helpers.SortMapString_Int(artists)

	const NUMBER_OF_ARTISTS int = 5 

	if len(filteredAndSortedArtists) > NUMBER_OF_ARTISTS {
		filteredAndSortedArtists = filteredAndSortedArtists[len(filteredAndSortedArtists) - NUMBER_OF_ARTISTS:]
	}
	
	// Key: Artist name - Value: number of reproductions
	return ctx.JSON(fiber.Map{
		"artists" : filteredAndSortedArtists,
	})
}