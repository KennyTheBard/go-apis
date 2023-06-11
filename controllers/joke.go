package controllers

import (
	"github.com/gofiber/fiber/v2"
	"hellkite.eu/go-api/services"
)

func (c *Controller) GetJoke(ctx *fiber.Ctx) error {
	joke := services.GetJoke()
	return ctx.JSON(joke)
}
