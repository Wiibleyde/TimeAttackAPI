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
	reqToken := c.Get("Authorization")
	if reqToken == "" {
		return c.SendString("Unauthorized")
	}

	authKey := data.GetLastAuthKey()
	if reqToken != authKey.AuthKey {
		return c.SendString("Unauthorized")
	}

	name := c.Query("search")
	races := data.SearchRaces(name)

	return c.JSON(races)
}

type createRaceApiStruct struct {
	Name string
}

func createRaceApi(c *fiber.Ctx) error {
	reqToken := c.Get("Authorization")
	if reqToken == "" {
		return c.SendString("Unauthorized")
	}

	authKey := data.GetLastAuthKey()
	if reqToken != authKey.AuthKey {
		return c.SendString("Unauthorized")
	}

	var body createRaceApiStruct
	err := c.BodyParser(&body)
	if err != nil {
		return c.SendString("Invalid body")
	}

	data.InsertRace(body.Name)

	return c.SendString("OK")
}

type removeRaceApiStruct struct {
	ID int
}

func removeRaceApi(c *fiber.Ctx) error {
	reqToken := c.Get("Authorization")
	if reqToken == "" {
		return c.SendString("Unauthorized")
	}

	authKey := data.GetLastAuthKey()
	if reqToken != authKey.AuthKey {
		return c.SendString("Unauthorized")
	}

	var body removeRaceApiStruct
	err := c.BodyParser(&body)
	if err != nil {
		return c.SendString("Invalid body")
	}

	data.DeleteRace(body.ID)

	return c.SendString("OK")
}

type updateRaceApiStruct struct {
	ID   int
	Name string
}

func updateRaceApi(c *fiber.Ctx) error {
	reqToken := c.Get("Authorization")
	if reqToken == "" {
		return c.SendString("Unauthorized")
	}

	authKey := data.GetLastAuthKey()
	if reqToken != authKey.AuthKey {
		return c.SendString("Unauthorized")
	}

	var body updateRaceApiStruct
	err := c.BodyParser(&body)
	if err != nil {
		return c.SendString("Invalid body")
	}

	data.UpdateRace(body.ID, body.Name)

	return c.SendString("OK")
}