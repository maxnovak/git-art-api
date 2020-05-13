package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Request object that api accepts for generating git pattern
type Request struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Pattern  string `json:"pattern" binding:"required"`
	Word     string `json:"word"`
	Year     int    `json:"year" binding:"required"`
	RepoName string `json:"repoName" binding:"required"`
}

// Ping for ping/pong default endpoint
func Ping(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "pong",
	})
}

// MakeArt accepts the json request
func MakeArt(context *gin.Context) {
	var json Request
	if err := context.ShouldBindJSON(&json); err != nil {
		if err.Error() == "EOF" {
			context.JSON(
				http.StatusBadRequest,
				gin.H{"error": "Empty Payload"},
			)
			return
		}

		context.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	log.Printf("Request: %+v\n", json)
}
