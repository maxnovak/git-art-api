package api

import (
	"bufio"
	"encoding/json"
	"fmt"
	"git-art/src/designs"
	"git-art/src/helpers"
	"git-art/src/models"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

// Console implementation of api
func Console() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Name of Repository: ")
	repoName, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	repoName = strings.Replace(repoName, "\n", "", -1)

	fmt.Print("Year for design: ")
	year, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	year = strings.Replace(year, "\n", "", -1)
	yearParsed, err := strconv.Atoi(year)

	fmt.Print("Select Design [checkered, give, matrix, table flip, word]: ")
	design, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	design = strings.Replace(design, "\n", "", -1)

	var word string
	if design == "word" {
		fmt.Print("What word would you like to display? Must be no more than 8 characters:")
		word, err = reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		word = strings.Replace(word, "\n", "", -1)
	}

	var matrixFile string
	if design == "matrix" {
		fmt.Print("Please enter the filepath to the matrix you would like to draw: ")
		matrixFile, err = reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		matrixFile = strings.Replace(matrixFile, "\n", "", -1)
	}

	fmt.Printf("Name of Repo '%s', design '%s', & year '%v' (y/n): ", repoName, design, yearParsed)
	confirmation, _, err := reader.ReadRune()
	if err != nil {
		log.Fatal(err)
	}
	if confirmation != 'y' {
		os.Exit(0)
	}

	fmt.Println("Creating repository")
	os.MkdirAll("./"+repoName, os.ModePerm)
	createRepo := exec.Command("git", "init")
	createRepo.Dir = repoName
	output, err := createRepo.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(output))

	date := time.Date(yearParsed, time.January, 1, 0, 0, 0, 0, time.UTC)
	if design == "checkered" {
		os.Chdir(repoName)
		designs.DrawCheckered(date, yearParsed)
	}
	if design == "give" {
		date = helpers.FindFirstSunday(date)
		os.Chdir(repoName)
		designs.DrawGive(date)
	}
	if design == "table flip" {
		date = helpers.FindFirstSunday(date)
		os.Chdir(repoName)
		designs.DrawTableFlip(date)
	}
	if design == "word" {
		date = helpers.FindFirstSunday(date)
		os.Chdir(repoName)
		designs.DrawWord(word, date)
	}
	if design == "matrix" {
		date = helpers.FindFirstSunday(date)
		data, err := ioutil.ReadFile(matrixFile)
		if err != nil {
			log.Fatal(err)
		}
		var matrixRequest models.MatrixRequest
		if err := json.Unmarshal(data, &matrixRequest); err != nil {
			fmt.Printf("Error whilde decoding %v\n", err)
			log.Fatal(err)
		}
		os.Chdir(repoName)
		designs.DrawMatixPatern(date, matrixRequest.Matrix)
	}
}
