package helpers

import (
	"log"
	"os"
	"os/exec"
	"time"
)

// CreateCommit makes an empty commit with no message
func CreateCommit(date time.Time) {
	os.Setenv("GIT_AUTHOR_DATE", date.String())
	os.Setenv("GIT_COMMITTER_DATE", date.String())

	args := []string{"commit", "--allow-empty", "--allow-empty-message", "-m "}
	createRepo := exec.Command("git", args...)

	_, err := createRepo.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
}

// FindFirstSunday finds to first Sunday of the year for make the pattern always line up right
func FindFirstSunday(date time.Time) time.Time {
	if int(date.Weekday()) > 0 {
		return date.Add(time.Hour * 24 * time.Duration(7-int(date.Weekday())))
	}
	return date
}

// AddDays is a helper function to make the copy pasta simpler
func AddDays(date time.Time, days int) time.Time {
	return date.Add(time.Hour * 24 * time.Duration(days))
}
