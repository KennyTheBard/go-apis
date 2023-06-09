package controllers

import (
	"github.com/gofiber/fiber/v2"
	"hellkite.eu/go-api/services"
)

func CreateUser(c *fiber.Ctx) error {
	request := RegisterRequest{}
	if err := c.BodyParser(&request); err != nil {
		return err
	}
	user, err := services.CreateUser(request.Name, request.Email)
	if err != nil {
		return err
	}
	return c.JSON(user)
}

type RegisterRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func GetAllUsers(c *fiber.Ctx) error {
	users, err := services.GetAllUsers()
	if err != nil {
		return err
	}
	return c.JSON(users)
}

func GetUserById(c *fiber.Ctx) error {
	request := GetUserByIdRequest{}
	if err := c.BodyParser(&request); err != nil {
		return err
	}
	user, err := services.GetUserById(request.ID)
	if err != nil {
		return err
	}
	return c.JSON(user)
}

type GetUserByIdRequest struct {
	ID int `json:"id"`
}

func UpdateUserName(c *fiber.Ctx) error {
	request := UpdateUserNameRequest{}
	if err := c.BodyParser(&request); err != nil {
		return err
	}
	user, err := services.UpdateUserById(request.ID, request.Name)
	if err != nil {
		return err
	}
	return c.JSON(user)
}

type UpdateUserNameRequest struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
