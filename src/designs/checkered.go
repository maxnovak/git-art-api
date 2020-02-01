package designs

import (
	"git-art/src/helpers"
	"time"
)

// DrawCheckered creates a git history with a checkered pattern
func DrawCheckered(date time.Time, yearForDesign int) {
	for {
		if date.Year() > yearForDesign {
			break
		}
		helpers.CreateCommit(date)
		date = date.Add(time.Hour * 48)
	}
}
