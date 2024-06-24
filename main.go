package main

import (
	"database/sql"
	"fmt"

	"github.com/abe444/BUBBLE_ENGINE/controller"
	"github.com/gin-gonic/gin"
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

func main() {
	Connect()
	router := gin.Default()

	router.MaxMultipartMemory = 8 << 20 // 8 MB max upload size

	router.LoadHTMLGlob("templates/*")
	router.Static("/app", "./app")
	router.Static("/entries", "./entries")
	router.Static("/static", "./static")
	router.StaticFile("/favicon.ico", "./favicon.ico")

	controller.StartupRoutes(router)

	router.Run(":8080")
}
