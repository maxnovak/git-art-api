package api

import "github.com/gin-gonic/gin"

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
