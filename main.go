package main

import (
	"data"
	"logger"
)

func main() {
	logger.InitLogger()
	logger.InfoLogger.Println("Initiliazing program...")
	data.InitDatabase()
	data.InitTables()
	logger.InfoLogger.Println("Program initialized.")
	data.InsertRace("TestRace")
	data.InsertTimer(1, "TestRacer", 100)
	data.InsertTimer(1, "TestRacer", 200)
	data.InsertTimer(1, "TestRacer", 300)
	data.GetTimersAsc(1)
}
