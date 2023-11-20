package api

import (
	"data"
	"logger"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type AddTimerRequest struct {
	RaceId    int
	RacerName string
	Timer     int
}

func addTimerApi(c *fiber.Ctx) error {
	reqToken := c.Get("Authorization")
	if reqToken == "" {
		return c.SendString("Unauthorized")
	}

	authKey := data.GetLastAuthKey()
	if reqToken != authKey.AuthKey {
		return c.SendString("Unauthorized")
	}

	var body AddTimerRequest
	err := c.BodyParser(&body)
	if err != nil {
		return c.SendString("Invalid body")
	}

	data.InsertTimer(body.RaceId, body.RacerName, body.Timer)

	return c.SendString("OK")
}

type RemoveTimerRequest struct {
	RaceId    int
	RacerName string
}

func removeTimerApi(c *fiber.Ctx) error {
	reqToken := c.Get("Authorization")
	if reqToken == "" {
		return c.SendString("Unauthorized")
	}

	authKey := data.GetLastAuthKey()
	if reqToken != authKey.AuthKey {
		return c.SendString("Unauthorized")
	}

	var body RemoveTimerRequest
	err := c.BodyParser(&body)
	if err != nil {
		return c.SendString("Invalid body")
	}

	data.DeleteTimer(body.RaceId, body.RacerName)

	return c.SendString("OK")
}

func raceLeaderboardApi(c *fiber.Ctx) error {
	reqToken := c.Get("Authorization")
	if reqToken == "" {
		return c.SendString("Unauthorized")
	}

	authKey := data.GetLastAuthKey()
	if reqToken != authKey.AuthKey {
		return c.SendString("Unauthorized")
	}


	raceId := c.Query("raceId", "0")
	raceIdInt, err := strconv.Atoi(raceId)
	if err != nil {
		logger.ErrorLogger.Fatal(err)
		return c.SendString("Invalid raceId")
	}

	leaderboard := data.GetTimersAsc(raceIdInt)

	return c.JSON(leaderboard)
}

func raceTimersApi(c *fiber.Ctx) error {
	raceId := c.Query("raceId", "0")
	raceIdInt, err := strconv.Atoi(raceId)
	if err != nil {
		logger.ErrorLogger.Fatal(err)
		return c.SendString("Invalid raceId")
	}

	timers := data.GetTimers(raceIdInt)

	return c.JSON(timers)
}

func racerTimersApi(c *fiber.Ctx) error {
	racerName := c.Query("racerName", "")
	if racerName == "" {
		return c.SendString("Invalid racerName")
	}

	timers := data.GetRacerTimers(racerName)

	return c.JSON(timers)
}

func racerLeaderboardApi(c *fiber.Ctx) error {
	racerName := c.Query("racerName", "")
	if racerName == "" {
		return c.SendString("Invalid racerName")
	}

	leaderboard := data.GetRacerLeaderboard(racerName)

	return c.JSON(leaderboard)
}