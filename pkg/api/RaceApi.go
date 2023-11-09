package api

import (
	"github.com/gofiber/fiber/v2"
)

func racesApi(c *fiber.Ctx) error {
	return c.SendString("OK")
}

func raceApi(c *fiber.Ctx) error {
	return c.SendString("OK")
}

func raceSearchApi(c *fiber.Ctx) error {
	return c.SendString("OK")
}

func createRaceApi(c *fiber.Ctx) error {
	return c.SendString("OK")
}

func removeRaceApi(c *fiber.Ctx) error {
	return c.SendString("OK")
}
