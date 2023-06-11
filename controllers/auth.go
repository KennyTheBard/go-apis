package controllers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"hellkite.eu/go-api/utils"
)

func (c *Controller) Authenticate(ctx *fiber.Ctx) error {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Date(2023, 10, 10, 12, 0, 0, 0, time.UTC)),
		Issuer:    "kenny",
	})
	tokenString, err := token.SignedString(utils.JwtSecret)
	if err != nil {
		panic(err)
	}
	response := &AuthenticateResponse{
		Token: tokenString,
	}
	return ctx.JSON(response)
}

type AuthenticateResponse struct {
	Token string `json:"token"`
}
