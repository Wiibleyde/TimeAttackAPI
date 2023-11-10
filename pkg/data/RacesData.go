package data

import (
	"database/sql"
	"logger"

	_ "github.com/go-sql-driver/mysql"
)

func InsertRace(Name string) bool {
	InitDatabase()
	defer CloseDatabase()
	var err error

	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM Races WHERE Name = ?", Name).Scan(&count)
	if err != nil && err.Error() != sql.ErrNoRows.Error() {
		logger.ErrorLogger.Println(err.Error())
		return false
	}
	if count > 0 {
		logger.ErrorLogger.Println("Races already exists.")
		return false
	}

	_, err = db.Exec("INSERT INTO Races (Name) VALUES (?)", Name)
	if err != nil {
		logger.ErrorLogger.Println(err.Error())
		return false
	}

	return true
}

func GetRace(id int) Race {
	InitDatabase()
	defer CloseDatabase()

	var err error
	var race Race

	err = db.QueryRow("SELECT * FROM Races WHERE ID = ?", id).Scan(&race.ID, &race.Name)
	if err != nil {
		logger.ErrorLogger.Println(err.Error())
	}

	return race
}

func GetRaces() []Race {
	InitDatabase()
	defer CloseDatabase()

	var err error
	var races []Race

	rows, err := db.Query("SELECT * FROM Races ORDER BY ID ASC")
	if err != nil {
		logger.ErrorLogger.Println(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var race Race
		err = rows.Scan(&race.ID, &race.Name)
		if err != nil {
			logger.ErrorLogger.Println(err.Error())
		}
		races = append(races, race)
	}

	return races
}