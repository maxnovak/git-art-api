package designs

import (
	"git-art/src/helpers"
	"time"
)

func DrawWord(word string, date time.Time) {
	var characters = []rune(word)
	for _, character := range characters {
		if character == 'a' {
			date = drawA(date)
			date = helpers.AddDays(date, 8)
		}
		if character == 'b' {
			date = drawB(date)
			date = helpers.AddDays(date, 9)
		}
	}
}

func drawA(date time.Time) time.Time {
	date = helpers.AddDays(date, 2)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 1)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 1)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 1)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 1)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 2)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 3)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 3)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 4)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 4)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 3)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 5)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 1)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 1)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 1)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 1)
	helpers.CreateCommit(date)

	return date
}

func drawB(date time.Time) time.Time {
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 1)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 1)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 1)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 1)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 1)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 1)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 1)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 3)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 3)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 1)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 3)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 3)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 1)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 3)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 3)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 2)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 1)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 2)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 1)
	helpers.CreateCommit(date)

	return date
}
