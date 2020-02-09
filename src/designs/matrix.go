package designs

import (
	"fmt"
	"git-art/src/helpers"
	"time"
)

func DrawMatixPatern(date time.Time, drawingMatrix [][]int) {
	for _, element := range drawingMatrix {
		fmt.Println(element)
		for _, item := range element {
			if item == 1 {
				helpers.CreateCommit(date)
			}
			date = helpers.AddDays(date, 1)
		}
	}
}
