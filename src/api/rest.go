package api

import (
	"log"

	"github.com/gin-gonic/gin"
)

// RestAPI implementation for creating repos as a zip
func RestAPI() {
	router := gin.Default()
	router.GET("/ping", Ping)
	router.POST("/make-art", MakeArt)

	router.Run() // listen and serve on localhost:8080
	log.Print("Server Started")
}
