package data

import (
	"logger"

	_ "github.com/go-sql-driver/mysql"
)

func GetLastAuthKey(id int) AuthKey {
	InitDatabase()
	defer CloseDatabase()

	var err error
	var authKeys []AuthKey

	err = db.QueryRow("SELECT * FROM AuthKeys WHERE ID = ?", id).Scan(&authKeys)
	if err != nil {
		logger.ErrorLogger.Println(err.Error())
	}

	return authKeys[len(authKeys)-1]
}
