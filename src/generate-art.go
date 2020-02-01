package main

import (
	"bufio"
	"fmt"
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

	fmt.Print("Select Design [checkered]: ")
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
			createCommit(date)
			date = date.Add(time.Hour * 48)
		}
	}

}

func createCommit(date time.Time) {
	os.Setenv("GIT_AUTHOR_DATE", date.String())
	os.Setenv("GIT_COMMITTER_DATE", date.String())

	args := []string{"commit", "--allow-empty", "--allow-empty-message", "-m "}
	createRepo := exec.Command("git", args...)

	_, err := createRepo.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
}

func findFirstSunday(date time.Time) time.Time {
	if int(date.Weekday()) > 0 {
		return date.Add(time.Hour * 24 * time.Duration(7-int(date.Weekday())))
	}
	return date
}
