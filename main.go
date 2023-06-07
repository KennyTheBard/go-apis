package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"hellkite.eu/go-api/controllers"
	"hellkite.eu/go-api/middlewares"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())

	publicGroup := app.Group("/public")
	publicGroup.Get("/authenticate", controllers.Authenticate)
	publicGroup.Static("/", "./public/index.html")

	apiGroup := app.Group("/api", middlewares.JwtAuthMiddleware)
	apiGroup.Get("/getStuff", controllers.GetStuff)
	apiGroup.Post("/pushThingy", controllers.PushThingy)

	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
