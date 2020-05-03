package api

import (
	"log"

	"github.com/gin-gonic/gin"
)

// RestAPI implementation for creating repos as a zip
func RestAPI() {
	r := gin.Default()
	r.GET("/ping", Handler)
	r.Run() // listen and serve on localhost:8080
	log.Print("Server Started")
}
