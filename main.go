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

	userController := &controllers.Controller{Logger: logger, DB: db}

	// configure routes
	publicGroup := app.Group("/public")
	publicGroup.Get("/authenticate", controllers.Authenticate)
	publicGroup.Post("/register", userController.CreateUser)
	publicGroup.Post("/getAllUsers", userController.GetAllUsers)
	publicGroup.Post("/getUserById", userController.GetUserById)
	publicGroup.Post("/updateUserName", userController.UpdateUserName)
	publicGroup.Post("/createOrder", userController.CreateOrder)
	publicGroup.Post("/getAllOrders", userController.GetAllOrders)
	publicGroup.Static("/", "./public/index.html")

	apiGroup := app.Group("/api", middlewares.JwtAuthMiddleware)
	apiGroup.Get("/getStuff", controllers.GetStuff)
	apiGroup.Post("/pushThingy", controllers.PushThingy)

	err = app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
