package data

import (
	"logger"

	_ "github.com/go-sql-driver/mysql"
)

func GetLastAuthKey() AuthKey {
	InitDatabase()
	defer CloseDatabase()

	var err error
	var authKeys AuthKey

	err = db.QueryRow("SELECT * FROM AuthKeys ORDER BY ID DESC LIMIT 1").Scan(&authKeys.ID, &authKeys.AuthKey)
	if err != nil {
		logger.ErrorLogger.Println(err.Error())
	}

	return authKeys
}
