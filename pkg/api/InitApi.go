package api

import (
	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/v2/middleware/cors"
)

func InitApi() {
	app := fiber.New()

	// app.Use(cors.New(cors.Config{
	// 	AllowOrigins:     "*",
	// 	AllowMethods:     "GET,POST,PUT,DELETE",
	// 	AllowHeaders:     "Origin, Content-Type, Accept, Authorization, Access-Control-Allow-Origin",
	// 	AllowCredentials: true,
	// }))

	// API status
	app.Get("/api", homeApi)
	app.Get("/api/status", statusApi)

	// API Races
	app.Get("/api/races", racesApi)
	app.Get("/api/race", raceApi)
	app.Get("/api/raceSearch", raceSearchApi)
	app.Post("/api/createRace", createRaceApi)
	app.Post("/api/removeRace", removeRaceApi)
	app.Post("/api/updateRace", updateRaceApi)

	// API Timers
	app.Post("/api/addTimer", addTimerApi)
	app.Post("/api/removeTimer", removeTimerApi)
	app.Get("/api/raceLeaderboard", raceLeaderboardApi)
	app.Get("/api/raceTimers", raceTimersApi)
	
	app.Get("/api/racerTimers", racerTimersApi)
	app.Get("/api/racerLeaderboard", racerLeaderboardApi)
	
	app.Listen(":3001")
}
