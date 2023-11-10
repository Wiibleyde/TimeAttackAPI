package api

import (
	"data"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func racesApi(c *fiber.Ctx) error {
	reqToken := c.Get("Authorization")
	if reqToken == "" {
		return c.SendString("Unauthorized")
	}

	authKey := data.GetLastAuthKey()
	if reqToken != authKey.AuthKey {
		return c.SendString("Unauthorized")
	}

	races := data.GetRaces()

	return c.JSON(races)
}

func raceApi(c *fiber.Ctx) error {
	reqToken := c.Get("Authorization")
	if reqToken == "" {
		return c.SendString("Unauthorized")
	}

	authKey := data.GetLastAuthKey()
	if reqToken != authKey.AuthKey {
		return c.SendString("Unauthorized")
	}

	id := c.Query("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		return c.SendString("Invalid ID")
	}

	race := data.GetRace(intId)

	return c.JSON(race)
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
