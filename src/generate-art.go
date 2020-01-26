package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Name of Repository: ")
	repoName, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	repoName = strings.Replace(repoName, "\n", "", -1)

	fmt.Print("Select Design: ")
	design, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	design = strings.Replace(design, "\n", "", -1)

	fmt.Printf("Name of Repo '%s' & design '%s:' (y/n): ", repoName, design)
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
}
