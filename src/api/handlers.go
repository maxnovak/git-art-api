package api

import (
	"fmt"
	"git-art/src/designs"
	"git-art/src/helpers"
	"git-art/src/models"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"

	"github.com/gin-gonic/gin"
)

// Request object that api accepts for generating git pattern
type Request struct {
	Name     string               `json:"name" binding:"required"`
	Email    string               `json:"email" binding:"required"`
	Pattern  string               `json:"pattern" binding:"required"`
	Word     string               `json:"word"`
	Year     int                  `json:"year,string" binding:"required"`
	RepoName string               `json:"repoName" binding:"required"`
	Matrix   models.MatrixRequest `json:"matrixObj"`
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

	log.Println("Creating repository")
	os.MkdirAll("./"+json.RepoName, os.ModePerm)
	os.Chdir(json.RepoName)

	createRepo := exec.Command("git", "init")
	output, err := createRepo.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(output))

	configureGitAuthor := exec.Command("git", "config", "user.name", json.Name)
	err = configureGitAuthor.Run()
	if err != nil {
		log.Fatal(err)
	}

	configureGitEmail := exec.Command("git", "config", "user.email", json.Email)
	err = configureGitEmail.Run()
	if err != nil {
		log.Fatal(err)
	}

	date := time.Date(json.Year, time.January, 1, 0, 0, 0, 0, time.UTC)
	if json.Pattern == "checkered" {
		designs.DrawCheckered(date, json.Year)
	}
	if json.Pattern == "give" {
		date = helpers.FindFirstSunday(date)
		designs.DrawGive(date)
	}
	if json.Pattern == "table flip" {
		date = helpers.FindFirstSunday(date)
		designs.DrawTableFlip(date)
	}
	if json.Pattern == "word" {
		date = helpers.FindFirstSunday(date)
		designs.DrawWord(json.Word, date)
	}
	if json.Pattern == "matrix" {
		date = helpers.FindFirstSunday(date)
		designs.DrawMatixPatern(date, json.Matrix.Matrix)
	}

	os.Chdir("..")

	zipFile := fmt.Sprintf("%v.zip", json.RepoName)
	helpers.Zipit(json.RepoName, zipFile)

	context.FileAttachment(zipFile, "file")

	os.RemoveAll(json.RepoName)
	os.Remove(zipFile)
}
