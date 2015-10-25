package holidays

import (
	"time"
)

func Holidays_DE(year int) YearHolidays {
	result := YearHolidays{}

	result.add(&YearDay{time.January, 1}, &Holiday{"Neujahrstag"})
	result.add(&YearDay{time.May, 1}, &Holiday{"Tag der Arbeit"})
	result.add(&YearDay{time.October, 3}, &Holiday{"Tag der Deutschen Einheit"})
	result.add(&YearDay{time.December, 25}, &Holiday{"1. Weihnachstag"})
	result.add(&YearDay{time.December, 26}, &Holiday{"2. Weihnachstag"})

	easter := easter(year)

	result.addFromTime(easter.AddDate(0, 0, 1), &Holiday{"Ostermontag"})
	result.addFromTime(easter.AddDate(0, 0, -2), &Holiday{"Karfreitag"})
	result.addFromTime(easter.AddDate(0, 0, 39), &Holiday{"Christi Himmelfahrt"})
	result.addFromTime(easter.AddDate(0, 0, 50), &Holiday{"Pfingstmontag"})

	return result
}

func Holidays_DE_NRW(year int) YearHolidays {
	result := Holidays_DE(year)

	result.add(&YearDay{time.November, 1}, &Holiday{"Allerheiligen"})
	result.addFromTime(easter(year).AddDate(0, 0, 60), &Holiday{"Fronleichnam"})

	return result
}
