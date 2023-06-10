package controllers

import (
	"github.com/gofiber/fiber/v2"
	"hellkite.eu/go-api/models"
)

func (c *Controller) CreateUser(ctx *fiber.Ctx) error {
	var request RegisterRequest
	if err := ctx.BodyParser(&request); err != nil {
		return err
	}
	user := models.User{Name: request.Name, Email: request.Email}
	if err := c.DB.Create(&request).Error; err != nil {
		return err
	}
	return ctx.JSON(user)
}

type RegisterRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (c *Controller) GetAllUsers(ctx *fiber.Ctx) error {
	var users []models.User
	if err := c.DB.Find(&users).Error; err != nil {
		return err
	}
	return ctx.JSON(users)
}

func (c *Controller) GetUserById(ctx *fiber.Ctx) error {
	var request GetUserByIdRequest
	if err := ctx.BodyParser(&request); err != nil {
		return err
	}
	var user models.User
	err := c.DB.First(&user, request.ID).Error
	if err != nil {
		return err
	}
	return ctx.JSON(user)
}

type GetUserByIdRequest struct {
	ID int `json:"id"`
}

func (c *Controller) UpdateUserName(ctx *fiber.Ctx) error {
	var request UpdateUserNameRequest
	if err := ctx.BodyParser(&request); err != nil {
		return err
	}
	var user models.User
	if err := c.DB.First(&user, request.ID).Error; err != nil {
		return err
	}
	user.Name = request.Name
	if err := c.DB.Save(&user).Error; err != nil {
		return err
	}
	return ctx.JSON(user)
}

type UpdateUserNameRequest struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
