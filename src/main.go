package main

import (
	"fmt"

	"profile-ms/handlers"
	//"profile-ms/models"

	"github.com/gofiber/fiber/v2"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func test(ctx *fiber.Ctx) error {
	return ctx.SendString("This is a test");
}

func initDB(){
	err := mgm.SetDefaultConfig(nil, "profiles_db",options.Client().ApplyURI("mongodb://localhost:27017"));
	if err != nil{
		panic("Error connecting to the database");
	}else{
		fmt.Println("Successfully connected to the database")
	}
}

func main(){
	app := fiber.New()
	initDB()
	
	/* userConfig := models.NewUserConfig(false,"/downloads","Black")

	err := mgm.Coll(userConfig).Create(userConfig)
	fmt.Println(err) */

	handlers.SetupUserConfigRoutes(app)

	app.Listen(":3000")

}