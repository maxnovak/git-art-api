package helpers

import (
	"archive/zip"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
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

// Zipit zips up a directory into a zip file
func Zipit(source, target string) error {
	zipfile, err := os.Create(target)
	if err != nil {
		return err
	}
	defer zipfile.Close()

	archive := zip.NewWriter(zipfile)
	defer archive.Close()

	info, err := os.Stat(source)
	if err != nil {
		return nil
	}

	var baseDir string
	if info.IsDir() {
		baseDir = filepath.Base(source)
	}

	filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		if baseDir != "" {
			header.Name = filepath.Join(baseDir, strings.TrimPrefix(path, source))
		}

		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}

		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()
		_, err = io.Copy(writer, file)
		return err
	})

	return err
}
