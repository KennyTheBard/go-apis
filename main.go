package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"hellkite.eu/go-api/controllers"
	"hellkite.eu/go-api/data"
	"hellkite.eu/go-api/middlewares"

	"github.com/apex/log"
)

func main() {
	var err error

	// init logger
	logger := log.WithFields(log.Fields{
		"file": "something.png",
		"type": "image/png",
		"user": "tobi",
	})

	// database connection
	db, err := data.InitDatabaseConnection()
	if err != nil {
		panic(err)
	}

	// configure server
	app := fiber.New()
	app.Use(cors.New())
	app.Use(recover.New())
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("logger", logger)
		return c.Next()
	})

	controller := &controllers.Controller{Logger: logger, DB: db}

	// configure routes
	publicGroup := app.Group("/public")
	publicGroup.Get("/authenticate", controller.Authenticate)
	publicGroup.Post("/register", controller.CreateUser)
	publicGroup.Post("/getAllUsers", controller.GetAllUsers)
	publicGroup.Post("/getUserById", controller.GetUserById)
	publicGroup.Post("/updateUserName", controller.UpdateUserName)
	publicGroup.Post("/createOrder", controller.CreateOrder)
	publicGroup.Post("/getAllOrders", controller.GetAllOrders)
	publicGroup.Get("/getJoke", controller.GetJoke)
	publicGroup.Static("/", "./public/index.html")

	apiGroup := app.Group("/api", middlewares.JwtAuthMiddleware)
	apiGroup.Get("/getStuff", controllers.GetStuff)
	apiGroup.Post("/pushThingy", controllers.PushThingy)

	err = app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
