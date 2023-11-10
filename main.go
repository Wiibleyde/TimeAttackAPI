package main

import (
	"api"
	"data"
	"logger"
)

func main() {
	logger.InitLogger()
	logger.InfoLogger.Println("Initiliazing program...")
	data.InitDatabase()
	data.InitTables()
	logger.InfoLogger.Println("Program initialized.")
	api.InitApi()
}
