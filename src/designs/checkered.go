package designs

import (
	"git-art/src/helpers"
	"time"
)

func DrawCheckered(date time.Time, yearForDesign int) {
	for {
		if date.Year() > yearForDesign {
			break
		}
		helpers.CreateCommit(date)
		date = date.Add(time.Hour * 48)
	}
}
