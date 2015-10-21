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

	result.addFromTime(easter.Add(Day), &Holiday{"Ostermontag"})
	result.addFromTime(easter.Add(-2*Day), &Holiday{"Karfreitag"})
	result.addFromTime(easter.Add(39*Day), &Holiday{"Christi Himmelfahrt"})
	result.addFromTime(easter.Add(50*Day), &Holiday{"Pfingstmontag"})

	return result
}

func Holidays_DE_NRW(year int) YearHolidays {
	result := Holidays_DE(year)

	result.add(&YearDay{time.November, 1}, &Holiday{"Allerheiligen"})
	result.addFromTime(easter(year).Add(60*Day), &Holiday{"Fronleichnam"})

	return result
}
