package designs

import (
	"git-art/src/helpers"
	"strings"
	"time"
)

func DrawWord(word string, date time.Time) {
	var characters = []rune(strings.ToLower(word))
	for _, character := range characters {
		if character == 'a' {
			date = drawA(date)
			date = helpers.AddDays(date, 8)
		}
		if character == 'b' {
			date = drawB(date)
			date = helpers.AddDays(date, 9)
		}
		if character == 'c' {
			date = drawC(date)
			date = helpers.AddDays(date, 9)
		}
		if character == 'd' {
			date = drawD(date)
			date = helpers.AddDays(date, 9)
		}
		if character == 'e' {
			date = drawE(date)
			date = helpers.AddDays(date, 8)
		}
		if character == 'f' {
			date = drawF(date)
			date = helpers.AddDays(date, 14)
		}
		if character == 'g' {
			date = drawG(date)
			date = helpers.AddDays(date, 9)
		}
		if character == 'h' {
			date = drawH(date)
			date = helpers.AddDays(date, 8)
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

func drawC(date time.Time) time.Time {
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
	date = helpers.AddDays(date, 2)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 6)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 1)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 6)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 1)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 6)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 2)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 4)
	helpers.CreateCommit(date)

	return date
}

func drawD(date time.Time) time.Time {
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
	date = helpers.AddDays(date, 6)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 1)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 6)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 1)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 6)
	helpers.CreateCommit(date)
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

	return date
}

func drawE(date time.Time) time.Time {
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
	date = helpers.AddDays(date, 1)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 6)
	helpers.CreateCommit(date)

	return date
}

func drawF(date time.Time) time.Time {
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
	date = helpers.AddDays(date, 4)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 3)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 4)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 7)
	helpers.CreateCommit(date)

	return date
}

func drawG(date time.Time) time.Time {
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
	date = helpers.AddDays(date, 2)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 6)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 1)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 6)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 1)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 4)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 2)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 2)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 3)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 1)
	helpers.CreateCommit(date)

	return date
}

func drawH(date time.Time) time.Time {
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
	date = helpers.AddDays(date, 4)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 7)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 7)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 4)
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

	return date
}
