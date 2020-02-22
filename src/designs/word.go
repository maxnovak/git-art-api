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
	var characters = []rune(strings.ToLower(word))
	for _, character := range characters {
		if character == 'a' {
			data, err := ioutil.ReadFile("../src/images/a.json")
			if err != nil {
				log.Fatal(err)
			}
			date = drawLetter(date, data)
		}
		if character == 'b' {
			data, err := ioutil.ReadFile("../src/images/b.json")
			if err != nil {
				log.Fatal(err)
			}
			date = drawLetter(date, data)
		}
		if character == 'c' {
			data, err := ioutil.ReadFile("../src/images/c.json")
			if err != nil {
				log.Fatal(err)
			}
			date = drawLetter(date, data)
		}
		if character == 'd' {
			data, err := ioutil.ReadFile("../src/images/d.json")
			if err != nil {
				log.Fatal(err)
			}
			date = drawLetter(date, data)
		}
		if character == 'e' {
			data, err := ioutil.ReadFile("../src/images/e.json")
			if err != nil {
				log.Fatal(err)
			}
			date = drawLetter(date, data)
		}
		if character == 'f' {
			data, err := ioutil.ReadFile("../src/images/f.json")
			if err != nil {
				log.Fatal(err)
			}
			date = drawLetter(date, data)
		}
		if character == 'g' {
			data, err := ioutil.ReadFile("../src/images/g.json")
			if err != nil {
				log.Fatal(err)
			}
			date = drawLetter(date, data)
		}
		if character == 'h' {
			data, err := ioutil.ReadFile("../src/images/h.json")
			if err != nil {
				log.Fatal(err)
			}
			date = drawLetter(date, data)
		}
	}
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
