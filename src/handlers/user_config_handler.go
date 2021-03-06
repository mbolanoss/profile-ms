package handlers

import (
	"context"
	"profile-ms/dtos"
	"profile-ms/models"

	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
)

func SetupUserConfigRoutes(app *fiber.App){

	//Used to fetch user config
	app.Get("/userConfig", GetUserConfig)

	// Used when creating an user, in order to save default config
	app.Post("/userConfig", CreateUserConfig)

	//Used to update an existing userConfig
	app.Put("/userConfig", UpdateUserConfig)
}

func CreateUserConfig(ctx *fiber.Ctx) error {
	userConfig := new(models.UserConfig)
	var err error

	if err = ctx.BodyParser(userConfig); err != nil {
		return ctx.Status(http.StatusInternalServerError).SendString("Error parsing user config body")
	}

	// Setting default config
	userConfig.AutoplayOn = false
	userConfig.DownloadRoute = "/downloads"
	userConfig.PreferredColor = "White"

	_ ,err = mgm.Coll(userConfig).InsertOne(context.TODO(), userConfig)
	if err != nil{
		return ctx.Status(http.StatusInternalServerError).SendString("Error while creating the user config in the DB")
	}

	return ctx.SendStatus(http.StatusOK)
}

func UpdateUserConfig(ctx *fiber.Ctx) error {
	var userConfigDto = new(dtos.UserConfigDto)
	var err error

	if err = ctx.BodyParser(userConfigDto); err != nil {
		return ctx.Status(http.StatusInternalServerError).SendString("Error parsing update user config body")
	}

	var userConfig models.UserConfig
	err = mgm.Coll(&userConfig).First(bson.M{"username" : userConfigDto.Username}, &userConfig)

	if err != nil{
		return ctx.Status(http.StatusBadRequest).SendString("User does not exist")
	}
	
	//User's config data is reasigned to the values of the dto
	userConfig.AutoplayOn = userConfigDto.AutoplayOn
	userConfig.PreferredColor = userConfigDto.PreferredColor
	userConfig.DownloadRoute = userConfigDto.DownloadRoute

	//User's config gets updated in the DB
	err = mgm.Coll(&userConfig).Update(&userConfig)

	if err != nil{
		return ctx.Status(http.StatusInternalServerError).SendString("Error while updating user config in the DB")
	}

	return ctx.SendStatus(http.StatusOK)
}

func GetUserConfig(ctx *fiber.Ctx) error {

	username := ctx.Query("username")

	if username == "" {
		return ctx.Status(http.StatusBadRequest).SendString("No username query found")
	}

	var userConfig models.UserConfig
	err := mgm.Coll(&userConfig).First(bson.M{"username" : username}, &userConfig)

	if err != nil{
		return ctx.Status(http.StatusBadRequest).SendString("User does not exist")
	}

	return ctx.JSON(bson.M{
		"autoplayOn" : userConfig.AutoplayOn,
		"downloadRoute" : userConfig.DownloadRoute,
		"preferredColor" : userConfig.PreferredColor,
	})
}