package designs

import (
	"fmt"
	"time"
)

func DrawMatixPatern(date time.Time, drawingMatrix [][]int) {
	for _, element := range drawingMatrix {
		fmt.Println(element)
		for _, item := range element {
			fmt.Println(item)
		}
	}
}
