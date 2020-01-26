package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	fmt.Printf("Name of Repo '%s' & design '%s'", repoName, design)
}
