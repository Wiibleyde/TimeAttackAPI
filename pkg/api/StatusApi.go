package api

import "github.com/gofiber/fiber/v2"

func homeApi(c *fiber.Ctx) error {
	return c.SendString("Hello \"users\"!")
}

func statusApi(c *fiber.Ctx) error {
	return c.SendString("OK")
}