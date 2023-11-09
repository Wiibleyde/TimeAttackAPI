package data

import (
	"database/sql"
	"logger"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const (
	databaseHost     = "localhost"
	databasePort     = "3306"
	databaseUser     = "root"
	databasePassword = ""
	databaseName     = "TimeAttack"
)

var db *sql.DB

func InitDatabase() {
	var err error
	if os.Getenv("STARTED_BY_DOCKER") == "true" {
		db, err = sql.Open("mysql", databaseUser+":"+databasePassword+"@tcp(db)/"+databaseName)
		if err != nil {
			logger.ErrorLogger.Println(err.Error())
		}
	} else {
		db, err = sql.Open("mysql", databaseUser+":"+databasePassword+"@tcp("+databaseHost+":"+databasePort+")/"+databaseName)
		if err != nil {
			logger.ErrorLogger.Println(err.Error())
		}
	}
}

func GetDatabase() *sql.DB {
	return db
}

func CloseDatabase() {
	db.Close()
}

func InitTables() bool {
	InitDatabase()
	defer CloseDatabase()
	var err error

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS Races (ID INT NOT NULL AUTO_INCREMENT, Name VARCHAR(255) NOT NULL, PRIMARY KEY (ID))")
	if err != nil {
		logger.ErrorLogger.Println(err.Error())
		return false
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS Timers (ID INT NOT NULL AUTO_INCREMENT, RacerName VARCHAR(255) NOT NULL, Duration INT NOT NULL, RaceID INT NOT NULL, PRIMARY KEY (ID))")
	if err != nil {
		logger.ErrorLogger.Println(err.Error())
		return false
	}

	return false
}
