package models

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitConnect() {
	if connection, err := sql.Open("mysql", urlParams() ); err != nil {
		panic(err)
	} else {
		fmt.Println("Connection success")
		db = connection
	}
}

func CloseConnect() {
	fmt.Println("Connection closed")
	db.Close()
}

func urlParams() string {
	return "DB_USER" + ":" + "DB_PASSWORD" + "@tcp(" + "DB_HOST" + ":" + "DB_PORT" + ")/" + "DB_NAME"
}
