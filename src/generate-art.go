package main

import (
	"bufio"
	"fmt"
	"git-art/src/helpers"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func main() {
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

	fmt.Print("Select Design [checkered, give]: ")
	design, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	design = strings.Replace(design, "\n", "", -1)

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

	os.Chdir(repoName)
	date := time.Date(yearParsed, time.January, 1, 0, 0, 0, 0, time.UTC)
	if design == "checkered" {
		for {
			if date.Year() > yearParsed {
				break
			}
			helpers.CreateCommit(date)
			date = date.Add(time.Hour * 48)
		}
	}
	if design == "give" {
		date = helpers.FindFirstSunday(date)
		drawGive(date)
	}
}

func drawGive(date time.Time) {
	// Shift by a month first
	date = date.Add(time.Hour * 24 * 7 * 4)

	// Left Arm
	date = helpers.AddDays(date, 2)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 6)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 4)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 3)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 3)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 5)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 1)
	helpers.CreateCommit(date)

	// Left Eye
	date = helpers.AddDays(date, 20)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 6)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 1)
	helpers.CreateCommit(date)

	// Mouth
	date = helpers.AddDays(date, 9)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 7)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 7)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 7)
	helpers.CreateCommit(date)

	// Right Eye
	date = helpers.AddDays(date, 5)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 6)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 1)
	helpers.CreateCommit(date)

	// Right Side
	date = helpers.AddDays(date, 13)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 8)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 8)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 1)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 8)
	helpers.CreateCommit(date)

	// Right Arm
	date = helpers.AddDays(date, 4)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 6)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 4)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 3)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 3)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 5)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 1)
	helpers.CreateCommit(date)
}
