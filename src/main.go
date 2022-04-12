package main

import "github.com/gofiber/fiber/v2"

func test(ctx *fiber.Ctx) error {
	return ctx.SendString("This is a test");
}

func main(){
	app := fiber.New()

	app.Get("/",  test)

	app.Listen(":3000")

}