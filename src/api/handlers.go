package api

import (
	"encoding/json"
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
	var request Request
	if err := context.ShouldBindJSON(&request); err != nil {
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

	log.Printf("Request: %+v\n", request)

	log.Println("Creating repository")
	os.MkdirAll("./"+request.RepoName, os.ModePerm)
	os.Chdir(request.RepoName)

	createRepo := exec.Command("git", "init")
	output, err := createRepo.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(output))

	configureGitAuthor := exec.Command("git", "config", "user.name", request.Name)
	err = configureGitAuthor.Run()
	if err != nil {
		log.Fatal(err)
	}

	configureGitEmail := exec.Command("git", "config", "user.email", request.Email)
	err = configureGitEmail.Run()
	if err != nil {
		log.Fatal(err)
	}

	date := time.Date(request.Year, time.January, 1, 0, 0, 0, 0, time.UTC)
	if request.Pattern == "checkered" {
		designs.DrawCheckered(date, request.Year)
	}
	if request.Pattern == "give" {
		date = helpers.FindFirstSunday(date)
		designs.DrawGive(date)
	}
	if request.Pattern == "table flip" {
		date = helpers.FindFirstSunday(date)
		designs.DrawTableFlip(date)
	}
	if request.Pattern == "word" {
		date = helpers.FindFirstSunday(date)
		designs.DrawWord(request.Word, date)
	}
	if request.Pattern == "shrug" {
		date = helpers.FindFirstSunday(date)
		matrix := helpers.GetDesign("shrug")
		var matrixRequest models.MatrixRequest
		if err := json.Unmarshal(matrix, &matrixRequest); err != nil {
			fmt.Printf("Error whilde decoding %v\n", err)
			log.Fatal(err)
		}
		designs.DrawMatixPatern(date, matrixRequest.Matrix)
	}
	os.Chdir("..")

	zipFile := fmt.Sprintf("%v.zip", request.RepoName)
	helpers.Zipit(request.RepoName, zipFile)

	context.FileAttachment(zipFile, "file")

	os.RemoveAll(request.RepoName)
	os.Remove(zipFile)
}
