package controllers

import (
	"github.com/gofiber/fiber/v2"
	"hellkite.eu/go-api/models"
)

func CreateUser(c *fiber.Ctx) error {
	var request RegisterRequest
	if err := c.BodyParser(&request); err != nil {
		return err
	}
	user := models.User{Name: request.Name, Email: request.Email}
	if err := models.DB.Create(&request).Error; err != nil {
		return err
	}
	return c.JSON(user)
}

type RegisterRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func GetAllUsers(c *fiber.Ctx) error {
	var users []models.User
	if err := models.DB.Find(&users).Error; err != nil {
		return err
	}
	return c.JSON(users)
}

func GetUserById(c *fiber.Ctx) error {
	var request GetUserByIdRequest
	if err := c.BodyParser(&request); err != nil {
		return err
	}
	var user models.User
	err := models.DB.First(&user, request.ID).Error
	if err != nil {
		return err
	}
	return c.JSON(user)
}

type GetUserByIdRequest struct {
	ID int `json:"id"`
}

func UpdateUserName(c *fiber.Ctx) error {
	var request UpdateUserNameRequest
	if err := c.BodyParser(&request); err != nil {
		return err
	}
	var user models.User
	if err := models.DB.First(&user, request.ID).Error; err != nil {
		return err
	}
	user.Name = request.Name
	if err := models.DB.Save(&user).Error; err != nil {
		return err
	}
	return c.JSON(user)
}

type UpdateUserNameRequest struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
