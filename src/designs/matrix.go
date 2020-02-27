package designs

import (
	"git-art/src/helpers"
	"time"
)

// DrawMatixPatern creates a git history with the passed in matrix design
func DrawMatixPatern(date time.Time, drawingMatrix [][]int) time.Time {
	for _, element := range drawingMatrix {
		for _, item := range element {
			for i := 0; i < item; i++ {
				helpers.CreateCommit(date)
			}
			date = helpers.AddDays(date, 1)
		}
	}
	return date
}
