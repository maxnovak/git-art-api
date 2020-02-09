package designs

import (
	"fmt"
	"git-art/src/helpers"
	"time"
)

// DrawMatixPatern creates a git history with the passed in matrix design
func DrawMatixPatern(date time.Time, drawingMatrix [][]int) {
	for _, element := range drawingMatrix {
		fmt.Println(element)
		for _, item := range element {
			for i := 0; i < item; i++ {
				helpers.CreateCommit(date)
			}
			date = helpers.AddDays(date, 1)
		}
	}
}
