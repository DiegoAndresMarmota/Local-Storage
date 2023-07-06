package pkg

import (
	"database/sql"
)

var db *sql.DB

func Ping() {
	if err := db.Ping(); err != nil {
		panic(err)
	}
}
