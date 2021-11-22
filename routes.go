package main

import (
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {

	api := app.Group("/v1")
	api.Get("/", func (c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
}
