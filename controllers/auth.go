package controllers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"hellkite.eu/go-api/utils"
)

var jwtSecret = []byte(utils.String(20))

func Authenticate(c *fiber.Ctx) error {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Date(2023, 10, 10, 12, 0, 0, 0, time.UTC)),
		Issuer:    "kenny",
	})
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		panic(err)
	}
	response := &AuthenticateResponse{
		Token: tokenString,
	}
	return c.JSON(response)
}

type AuthenticateResponse struct {
	Token string `json:"token"`
}
