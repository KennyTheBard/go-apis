package controllers

import (
	"strconv"
	"time"

	"github.com/apex/log"
	"github.com/gofiber/fiber/v2"
	"hellkite.eu/go-api/models"
)

func (c *Controller) CreateOrder(ctx *fiber.Ctx) error {
	var request CreateOrderRequest
	if err := ctx.BodyParser(&request); err != nil {
		return err
	}

	// check that the user exists
	var user models.User
	if err := c.DB.First(&user, request.UserID).Error; err != nil {
		logger, ok := ctx.Locals("logger").(*log.Entry)
		if !ok {
			// Handle the case when the logger is not found in locals
			return fiber.ErrInternalServerError
		}
		logger.Error("User " + strconv.Itoa(request.UserID) + " does not exist")
		return err
	}

	order := models.Order{
		UserID:      request.UserID,
		OrderDate:   request.OrderDate,
		TotalAmount: request.TotalAmount,
	}
	if err := c.DB.Create(&order).Error; err != nil {
		return err
	}
	return ctx.JSON(order)
}

type CreateOrderRequest struct {
	UserID      int       `json:"userId"`
	OrderDate   time.Time `json:"orderDate"`
	TotalAmount float64   `json:"totalAmount"`
}

func (c *Controller) GetAllOrders(ctx *fiber.Ctx) error {
	var orders []models.Order
	if err := c.DB.Find(&orders).Error; err != nil {
		return err
	}
	return ctx.JSON(orders)
}
