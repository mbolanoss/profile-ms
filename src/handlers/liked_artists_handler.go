package handlers

import (
	"context"
	"net/http"
	"profile-ms/models"

	"github.com/gofiber/fiber/v2"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
)

func SetupLikedArtistsRoutes(app *fiber.App){
	// Used to fetch all the liked artists from an user
	app.Get("/likedArtists", GetLikedArtists)

	// Used to create the list in the DB
	app.Post("/likedArtists/create", CreateLikedArtistsList)

	// Used to add a new liked artist to a user's list
	app.Post("/likedArtists", AddLikedArtist)

	// Used to delete an artist from a list
	app.Delete("/likedArtists", DeleteArtist)
}

func GetLikedArtists(ctx *fiber.Ctx) error {
	username := ctx.Query("username")

	if username == "" {
		return ctx.Status(http.StatusBadRequest).SendString("No username query found in url")
	}

	var likedArtistsList models.LikedArtistList
	err := mgm.Coll(&likedArtistsList).First(bson.M{"username" : username}, &likedArtistsList)

	if err != nil{
		return ctx.Status(http.StatusBadRequest).SendString("User does not exist")
	}

	return ctx.Status(http.StatusOK).JSON(likedArtistsList)
}

func CreateLikedArtistsList(ctx *fiber.Ctx) error {

	username := ctx.Query("username")

	if username == "" {
		return ctx.Status(http.StatusBadRequest).SendString("No username query found in url")
	}

	likedArtistsList := models.NewLikedArtistList(username)

	//The list is saved in the DB
	_, err := mgm.Coll(likedArtistsList).InsertOne(context.TODO(), likedArtistsList)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).SendString("Error while creating the user's liked artists list in the DB")
	}

	return ctx.SendStatus(http.StatusOK)
}

func AddLikedArtist(ctx *fiber.Ctx) error {
	artistName := ctx.Query("artistName")
	username := ctx.Query("username")

	if artistName == "" || username == ""{
		return ctx.Status(http.StatusBadRequest).SendString("No artist name or username query found in url")
	}

	var likedArtistsList models.LikedArtistList
	err := mgm.Coll(&likedArtistsList).First(bson.M{"username" : username}, &likedArtistsList)

	if err != nil{
		return ctx.Status(http.StatusBadRequest).SendString("User does not exist")
	}

	//The artist gets appended and the list is updated in the DB
	likedArtistsList.Artists = append(likedArtistsList.Artists, artistName)
	err = mgm.Coll(&likedArtistsList).Update(&likedArtistsList)

	if err != nil{
		return ctx.Status(http.StatusInternalServerError).SendString("Error while adding the artist to the user's list")
	}

	return ctx.SendStatus(http.StatusOK)
}

func DeleteArtist(ctx *fiber.Ctx) error {
	artistName := ctx.Query("artistName")
	username := ctx.Query("username")

	if artistName == "" || username == ""{
		return ctx.Status(http.StatusBadRequest).SendString("No artist name or username query found in url")
	}

	var likedArtistsList models.LikedArtistList
	err := mgm.Coll(&likedArtistsList).First(bson.M{"username" : username}, &likedArtistsList)

	if err != nil{
		return ctx.Status(http.StatusBadRequest).SendString("User does not exist")
	}

	for index, artist := range likedArtistsList.Artists{
		if artist == artistName {
			// Delete artist from the list
			likedArtistsList.Artists = models.DeleteArtist(likedArtistsList.Artists, index)
			break
		}

		// If artist not found
		if index == len(likedArtistsList.Artists) - 1 {
			return ctx.Status(http.StatusBadRequest).SendString("Artist does not exist")
		}
	}

	// The list gets updated in the DB
	err = mgm.Coll(&likedArtistsList).Update(&likedArtistsList)

	if err != nil{
		return ctx.Status(http.StatusInternalServerError).SendString("Error while deleting artist from the list in the DB")
	}

	return ctx.SendStatus(http.StatusOK)
}