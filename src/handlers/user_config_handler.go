package handlers

import (
	"profile-ms/models"

	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/kamva/mgm/v3"
)

func SetupUserConfigRoutes(app *fiber.App){
	// Used when creating an user, in order to save default config
	app.Post("/userConfig", CreateUserConfig)
}

func CreateUserConfig(ctx *fiber.Ctx) error {
	userConfig := models.NewUserConfig(true, "/downloads", "White")

	if err := mgm.Coll(userConfig).Create(userConfig); err != nil{
		return ctx.Status(http.StatusInternalServerError).SendString("Error while saving the user config in the DB")
	}

	return ctx.SendStatus(http.StatusOK)
}