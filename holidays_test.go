package holidays

import (
	"testing"
	"time"
)

func TestIsHoliday(t *testing.T) {
	SetHolidaysFunction(Holidays_DE)

	assertIsHoliday(t, 2009, time.January, 1)
	assertIsNotHoliday(t, 2009, time.January, 2)
}

func assertHoliday(t *testing.T, year int, month time.Month, day int, expectedHoliday string) {
	expectedTime := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	holidays := GetHolidays(expectedTime)
	for _, holiday := range holidays {
		if holiday.Name == expectedHoliday {
			return
		}
	}
	t.Errorf("Expect holiday '%s' on '%s'", expectedHoliday, expectedTime)
}

func assertIsHoliday(t *testing.T, year int, month time.Month, day int) {
	expectedTime := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	if !IsHoliday(expectedTime) {
		t.Errorf("Expect '%s' to be a holiday", expectedTime)
	}
}

func assertIsNotHoliday(t *testing.T, year int, month time.Month, day int) {
	expectedTime := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	if IsHoliday(expectedTime) {
		t.Errorf("Expect '%s' not to be a holiday", expectedTime)
	}
}
