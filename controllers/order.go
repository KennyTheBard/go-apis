package controllers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"hellkite.eu/go-api/models"
)

func CreateOrder(c *fiber.Ctx) error {
	var request CreateOrderRequest
	if err := c.BodyParser(&request); err != nil {
		return err
	}

	// check that the user exists
	var user models.User
	if err := models.DB.First(&user, request.UserID).Error; err != nil {
		return err
	}

	order := models.Order{
		UserID:      request.UserID,
		OrderDate:   request.OrderDate,
		TotalAmount: request.TotalAmount,
	}
	if err := models.DB.Create(&order).Error; err != nil {
		return err
	}
	return c.JSON(order)
}

type CreateOrderRequest struct {
	UserID      int       `json:"userId"`
	OrderDate   time.Time `json:"orderDate"`
	TotalAmount float64   `json:"totalAmount"`
}

func GetAllOrders(c *fiber.Ctx) error {
	var orders []models.Order
	if err := models.DB.Find(&orders).Error; err != nil {
		return err
	}
	return c.JSON(orders)
}
