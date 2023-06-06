package main

import (
	"github.com/gofiber/fiber/v2"
	"hellkite.eu/go-api/controllers"
)

func main() {
	app := fiber.New()

	app.Get("/authenticate", controllers.Authenticate)

	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
