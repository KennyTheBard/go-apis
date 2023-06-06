package controllers

import "github.com/gofiber/fiber/v2"

type Thing struct {
	Name  string `json:name`
	Value int    `json:value`
}

var db []Thing = make([]Thing, 0)

func GetStuff(c *fiber.Ctx) error {
	return c.JSON(db)
}

func PushThingy(c *fiber.Ctx) error {
	thingy := Thing{}
	c.BodyParser(thingy)
	db = append(db, thingy)

	return c.SendStatus(201)
}
