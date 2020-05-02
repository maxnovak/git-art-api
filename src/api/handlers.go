package api

import "github.com/gin-gonic/gin"

// Handler for ping/pong default endpoint
func Handler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
