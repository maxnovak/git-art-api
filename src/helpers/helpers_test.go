package helpers

import (
	"testing"
	"time"
)

func TestAddDays(t *testing.T) {
	date := time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)
	actual := AddDays(date, 5)
	expected := time.Date(2000, time.January, 6, 0, 0, 0, 0, time.UTC)
	if actual != expected {
		t.Errorf("Date was not properly incremented, got: %v, want: %v", actual, expected)
	}
}

func TestFindFirstSunday(t *testing.T) {
	startsRandomDay := time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)
	actual := FindFirstSunday(startsRandomDay)
	expected := time.Date(2000, time.January, 2, 0, 0, 0, 0, time.UTC)
	if actual != expected {
		t.Errorf("Found wrong day, got: %v, want: %v", actual, expected)
	}

	date := time.Date(2017, time.January, 1, 0, 0, 0, 0, time.UTC)
	actual = FindFirstSunday(date)
	expected = time.Date(2017, time.January, 1, 0, 0, 0, 0, time.UTC)
	if actual != expected {
		t.Errorf("Found wrong day, got: %v, want: %v", actual, expected)
	}
}
