package designs

import (
	"git-art/src/helpers"
	"time"
)

func DrawGive(date time.Time) {
	// Shift by a month first
	date = date.Add(time.Hour * 24 * 7 * 4)

	// Left Arm
	date = helpers.AddDays(date, 2)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 6)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 4)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 3)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 3)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 5)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 1)
	helpers.CreateCommit(date)

	// Left Eye
	date = helpers.AddDays(date, 20)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 6)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 1)
	helpers.CreateCommit(date)

	// Mouth
	date = helpers.AddDays(date, 9)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 7)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 7)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 7)
	helpers.CreateCommit(date)

	// Right Eye
	date = helpers.AddDays(date, 5)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 6)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 1)
	helpers.CreateCommit(date)

	// Right Side
	date = helpers.AddDays(date, 13)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 8)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 8)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 1)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 8)
	helpers.CreateCommit(date)

	// Right Arm
	date = helpers.AddDays(date, 4)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 6)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 4)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 3)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 3)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 5)
	helpers.CreateCommit(date)
	date = helpers.AddDays(date, 1)
	helpers.CreateCommit(date)
}
