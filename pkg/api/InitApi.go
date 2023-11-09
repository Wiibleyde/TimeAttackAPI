package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func InitApi() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// API status
	app.Get("/api", homeApi)
	app.Get("/api/status", statusApi)

	// API Races
	app.Get("/api/races", racesApi)
	app.Get("/api/race", raceApi)
	app.Get("/api/raceSearch", raceSearchApi)
	app.Post("/api/createRace", createRaceApi)
	app.Post("/api/removeRace", removeRaceApi)

	app.Listen(":3001")
}
