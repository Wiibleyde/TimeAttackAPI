package data

import (
	"logger"

	_ "github.com/go-sql-driver/mysql"
)

func InsertTimer(RaceId int, RacerName string, TimeInSec int) bool {
	InitDatabase()
	defer CloseDatabase()
	var err error

	var count int
	err = db.QueryRow("SELECT * FROM Timers WHERE RacerName = ? AND RaceID = ?", RacerName, RaceId).Scan(&count)
	if err != nil {
		logger.ErrorLogger.Println(err.Error())
		return false
	}
	if count > 0 {
		logger.ErrorLogger.Println("Race already exists.")
		err = db.QueryRow("UPDATE Timers SET Duration = ? WHERE RacerName = ? AND RaceID = ?", TimeInSec, RacerName, RaceId).Scan(&count)
		if err != nil {
			logger.ErrorLogger.Println(err.Error())
			return false
		}
	} else {
		_, err = db.Exec("INSERT INTO Timers (RacerName, Duration, RaceID) VALUES (?, ?, ?)", RacerName, TimeInSec, RaceId)
		if err != nil {
			logger.ErrorLogger.Println(err.Error())
			return false
		}
	}
	return true
}

func DeleteTimer(RaceID int, RacerName string) bool {
	InitDatabase()
	defer CloseDatabase()
	var err error

	_, err = db.Exec("DELETE FROM Timers WHERE RaceID = ? AND RacerName = ?", RaceID, RacerName)
	if err != nil {
		logger.ErrorLogger.Println(err.Error())
		return false
	}

	return true
}

func GetTimersAsc(RaceID int) []Timers {
	InitDatabase()
	defer CloseDatabase()

	var err error
	var timers []Timers

	err = db.QueryRow("SELECT * FROM Timers WHERE RaceID = ? ORDER BY Duration ASC", RaceID).Scan(&timers)
	if err != nil {
		logger.ErrorLogger.Println(err.Error())
	}

	return timers
}

func GetTimersDesc(RaceID int) []Timers {
	InitDatabase()
	defer CloseDatabase()

	var err error
	var timers []Timers

	err = db.QueryRow("SELECT * FROM Timers WHERE RaceID = ? ORDER BY Duration DESC", RaceID).Scan(&timers)
	if err != nil {
		logger.ErrorLogger.Println(err.Error())
	}

	return timers
}

func GetTimers(RaceID int) []Timers {
	InitDatabase()
	defer CloseDatabase()

	var err error
	var timers []Timers

	err = db.QueryRow("SELECT * FROM Timers WHERE RaceID = ?", RaceID).Scan(&timers)
	if err != nil {
		logger.ErrorLogger.Println(err.Error())
	}

	return timers
}

func GetRacerTimers(RacerName string) []Timers {
	InitDatabase()
	defer CloseDatabase()

	var err error
	var timers []Timers

	err = db.QueryRow("SELECT * FROM Timers WHERE RacerName = ?", RacerName).Scan(&timers)
	if err != nil {
		logger.ErrorLogger.Println(err.Error())
	}

	return timers
}
