package designs

import (
	"encoding/json"
	"fmt"
	"git-art/src/helpers"
	"git-art/src/models"
	"io/ioutil"
	"log"
	"strings"
	"time"
)

// DrawWord creates a git history based on the passed in word
func DrawWord(word string, date time.Time) {
	date = helpers.AddDays(date, 1)
	var characters = []rune(strings.ToLower(word))
	for _, character := range characters {
		data := getLetter(character)
		date = drawLetter(date, data)
	}
}

func getLetter(character rune) []byte {
	data, err := ioutil.ReadFile(fmt.Sprintf("../src/images/%v.json", string(character)))
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func drawLetter(date time.Time, data []byte) time.Time {

	var matrixRequest models.MatrixRequest
	if err := json.Unmarshal(data, &matrixRequest); err != nil {
		fmt.Printf("Error whilde decoding %v\n", err)
		log.Fatal(err)
	}
	date = DrawMatixPatern(date, matrixRequest.Matrix)
	date = helpers.AddDays(date, 7)

	return date
}
