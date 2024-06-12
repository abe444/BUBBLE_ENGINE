package main

import (
	"github.com/abe444/BUBBLE_ENGINE/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

    router.MaxMultipartMemory = 8 << 20  // 8 MiB

	router.LoadHTMLGlob("templates/*")
	router.Static("/entries", "./entries")
	router.Static("/static", "./static")
	router.StaticFile("/favicon.ico", "./favicon.ico")

    controller.StartupRoutes(router)

	router.Run(":8080")
}
