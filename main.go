package main

import (
	"api"
	"data"
	"logger"
)

func main() {
	logger.InitLogger()
	logger.InfoLogger.Println("Initiliazing program...")
	status := data.InitDatabase()
	if !status {
		logger.ErrorLogger.Println("Database initialization failed.")
		return
	}
	data.InitTables()
	logger.InfoLogger.Println("Program initialized.")
	api.InitApi()
}
