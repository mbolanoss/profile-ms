package handlers

import (
	"context"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"

	"profile-ms/dtos"
	"profile-ms/helpers"
	"profile-ms/models"
)

func SetupMainSongsRoutes(app *fiber.App){
	// Used to create the list in the DB
	app.Post("/mainSongs", CreateMainSongsList)

	// Used to add 1 to the number of reproductions of a song
	app.Put("/mainSongs", AddReproduction)

	// Used to get the top played songs
	app.Get("/mainSongs", GetTopPlayedSongs)
}

func CreateMainSongsList(ctx *fiber.Ctx) error {
	username := ctx.Query("username")

	if username == "" {
		return ctx.Status(http.StatusBadRequest).SendString("No username query found in url")
	}

	mainSongsList := models.NewMainSongsList(username)

	mgm.Coll(mainSongsList).InsertOne(context.TODO(), mainSongsList)

	return ctx.SendStatus(http.StatusOK)
}

func AddReproduction(ctx *fiber.Ctx) error {
	var err error

	addReproductionDto := new(dtos.AddReproductionDto)
	err = ctx.BodyParser(addReproductionDto)
	if err != nil {
		ctx.Status(http.StatusInternalServerError).SendString("Error parsing add reproduction body")
	}

	var mainSongsList models.MainSongsList

	err = mgm.Coll(&mainSongsList).First(bson.M{"username" : addReproductionDto.Username}, &mainSongsList)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).SendString("Username does not exist")
	}

	mainSongsList.Songs[addReproductionDto.SongId] += 1
	err = mgm.Coll(&mainSongsList).Update(&mainSongsList)

	if err != nil{
		return ctx.Status(http.StatusInternalServerError).SendString("Error while updating song's number of reproductions in the DB")
	}

	return ctx.SendStatus(http.StatusOK)
}

func GetTopPlayedSongs(ctx *fiber.Ctx) error {
	var err error

	topPlayedSongsDto := new(dtos.TopPlayedSongsDto)
	err = ctx.BodyParser(topPlayedSongsDto)
	if err != nil {
		ctx.Status(http.StatusInternalServerError).SendString("Error parsing top played songs body")
	}

	var mainSongsList models.MainSongsList

	err = mgm.Coll(&mainSongsList).First(bson.M{"username" : topPlayedSongsDto.Username}, &mainSongsList)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).SendString("Username does not exist")
	}

	
	const NUMBER_OF_SONGS int = 5 

	sorted := helpers.SortMap(mainSongsList.Songs)
	sorted = sorted[len(mainSongsList.Songs) - NUMBER_OF_SONGS:]
	
	// Key: songID - Value: number of reproductions
	return ctx.JSON(fiber.Map{
		"songs" : sorted,
	})
}