package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"

	"profile-ms/handlers"
)

func initDB(){
	// Loading env variables
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	DB_NAME := os.Getenv("DB_NAME")
	DB_PASS := os.Getenv("DB_PASS")
	
	// Local DB
	//dbURL := fmt.Sprintf("mongodb://%s:%s", DB_HOST, DB_PORT)
	//DB_HOST := os.Getenv("DB_HOST")
	//DB_PORT := os.Getenv("DB_PORT")

	//Deployed DB
	dbURL := fmt.Sprintf("mongodb+srv://root:%s@cluster0.lno2t.mongodb.net/%s?retryWrites=true&w=majority", DB_PASS, DB_NAME)
	
	err = mgm.SetDefaultConfig(nil, DB_NAME,options.Client().ApplyURI(dbURL));

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

func dockerTest(ctx *fiber.Ctx) error {
	return ctx.SendString("Docker test")
}

func main(){

	app := fiber.New()
	initDB()

	SetupAllRoutes(app)
	app.Get("/", dockerTest)

	app.Listen(":3000")

}