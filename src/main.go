package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"

	"profile-ms/handlers"
)

func initDB(){
	err := mgm.SetDefaultConfig(nil, "profiles_db",options.Client().ApplyURI("mongodb://localhost:27017"));
	if err != nil{
		panic("Error connecting to the database");
	}else{
		fmt.Println("Successfully connected to the database")
	}
}

func SetupAllRoutes(app *fiber.App){
	handlers.SetupLikedArtistsRoutes(app)
	handlers.SetupUserConfigRoutes(app)
	handlers.SetupMainSongsRoutes(app)
	handlers.SetupPlayedArtistsRoutes(app)
}

func main(){

	app := fiber.New()
	initDB()

	SetupAllRoutes(app)

	app.Listen(":3000")

}