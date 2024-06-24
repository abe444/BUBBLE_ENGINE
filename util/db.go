package util

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func Connect() {
	var err error
	Db, err = sql.Open("mariadb", "abe:pass1234@tcp(127.0.0.1:3306)/bubble")
	if err != nil {
		panic(err.Error())
	}
	if err = Db.Ping(); err != nil {
		panic(err.Error())
	}
	fmt.Println("Successfully connected to the database.")
}
