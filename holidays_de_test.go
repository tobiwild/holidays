package holidays

import (
	"testing"
	"time"
)

func TestHolidays_DE(t *testing.T) {
	SetHolidaysFunction(Holidays_DE)

	assertHoliday(t, 2009, time.January, 1, "Neujahrstag")
	assertHoliday(t, 2009, time.May, 1, "Tag der Arbeit")
	assertHoliday(t, 2009, time.October, 3, "Tag der Deutschen Einheit")
	assertHoliday(t, 2009, time.December, 25, "1. Weihnachstag")
	assertHoliday(t, 2009, time.December, 26, "2. Weihnachstag")
	assertHoliday(t, 2009, time.April, 13, "Ostermontag")
	assertHoliday(t, 2009, time.April, 10, "Karfreitag")
	assertHoliday(t, 2009, time.May, 21, "Christi Himmelfahrt")
	assertHoliday(t, 2009, time.June, 1, "Pfingstmontag")
}

func TestHolidays_DE_NRW(t *testing.T) {
	SetHolidaysFunction(Holidays_DE_NRW)

	assertHoliday(t, 2009, time.November, 1, "Allerheiligen")
	assertHoliday(t, 2009, time.June, 11, "Fronleichnam")
}
