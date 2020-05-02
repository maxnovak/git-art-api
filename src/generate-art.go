package main

import (
	"bufio"
	"fmt"
	"git-art/src/api"
	"log"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Console or REST API start? [console/rest]")
	apiType, err := reader.ReadString('\n')
	apiType = strings.Replace(apiType, "\n", "", -1)
	if err != nil {
		log.Fatal(err)
	}

	if apiType == "console" {
		api.Console()
	}
	if apiType == "rest" {
		api.RestAPI()
	}
	os.Exit(0)
}
