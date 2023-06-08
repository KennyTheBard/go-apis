package controllers

import (
	"github.com/gofiber/fiber/v2"
	"hellkite.eu/go-api/models"
	"hellkite.eu/go-api/services"
)

func Register(c *fiber.Ctx) error {
	var err error
	request := RegisterRequest{}
	if err = c.BodyParser(&request); err != nil {
		return err
	}
	var user *models.User
	if user, err = services.CreateUser(request.Name, request.Email); err != nil {
		return err
	}
	return c.JSON(user)
}

type RegisterRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
