package middlewares

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"hellkite.eu/go-api/utils"
)

func JwtAuthMiddleware(c *fiber.Ctx) error {
	authorization := c.Request().Header.Peek("Authorization")
	authorizationParts := strings.Split(string(authorization), " ")
	if len(authorizationParts) != 2 || authorizationParts[0] != "Bearer" {
		err := errors.New("Invalid authorization")
		fmt.Println(err)
		return err
	}
	tokenString := authorizationParts[1]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return utils.JwtSecret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims)
	} else {
		fmt.Println(err)
		return err
	}

	return c.Next()
}
