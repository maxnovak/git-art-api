package main

import (
	"git-art/src/api"
	"git-art/src/env"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// RestAPI implementation for creating repos as a zip
func main() {
	router := gin.Default()
	env.Load()

	router.Use(cors.Default())

	router.GET("/ping", api.Ping)
	router.POST("/make-art", api.MakeArt)

	router.Run(":" + os.Getenv("PORT")) // listen and serve on localhost:8080
	log.Print("Server Started")
}
