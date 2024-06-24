package main

import (
	"github.com/abe444/BUBBLE_ENGINE/controller"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.MaxMultipartMemory = 8 << 20 // 8 MB max upload size

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	router.LoadHTMLGlob("templates/*")
	router.Static("/app", "./app")
	router.Static("/entries", "./entries")
	router.Static("/static", "./static")
	router.StaticFile("/favicon.ico", "./favicon.ico")

	controller.StartupRoutes(router)

	router.Run(":8080")
}
