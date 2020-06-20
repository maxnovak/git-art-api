package main

import (
	"git-art/src/api"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

// RestAPI implementation for creating repos as a zip
func main() {
	router := gin.Default()
	router.GET("/ping", api.Ping)
	router.POST("/make-art", api.MakeArt)

	router.Run(":" + os.Getenv("PORT")) // listen and serve on localhost:8080
	log.Print("Server Started")
}
