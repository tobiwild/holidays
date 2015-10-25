package holidays

import (
	"github.com/vjeantet/eastertime"
	"time"
)

var (
	easterCache = make(map[int]*time.Time)
	holidays    = &Holidays{Function: Holidays_DE_NRW}
)

type Holidays struct {
	Function func(year int) YearHolidays
	cache    map[int]*YearHolidays
}

type YearDay struct {
	Month time.Month
	Day   int
}

type Holiday struct {
	Name string
}

type YearHolidays map[YearDay][]*Holiday

func (y YearHolidays) add(yearDay *YearDay, holiday *Holiday) {
	if _, ok := y[*yearDay]; ok {
		y[*yearDay] = append(y[*yearDay], holiday)
	} else {
		y[*yearDay] = []*Holiday{holiday}
	}
}

func (y YearHolidays) addFromTime(t time.Time, holiday *Holiday) {
	y.add(&YearDay{t.Month(), t.Day()}, holiday)
}

func (h *Holidays) GetHolidays(t time.Time) []*Holiday {
	yearDay := YearDay{t.Month(), t.Day()}

	yearHolidays := h.yearHolidays(t.Year())
	return yearHolidays[yearDay]
}

func GetHolidays(t time.Time) []*Holiday {
	return holidays.GetHolidays(t)
}

func (h *Holidays) IsHoliday(t time.Time) bool {
	return len(h.GetHolidays(t)) > 0
}

func IsHoliday(t time.Time) bool {
	return holidays.IsHoliday(t)
}

func SetHolidaysFunction(function func(year int) YearHolidays) {
	holidays = &Holidays{Function: function}
}

func (h *Holidays) yearHolidays(year int) YearHolidays {
	if h.cache == nil {
		h.cache = make(map[int]*YearHolidays)
	}
	if _, ok := h.cache[year]; !ok {
		yearHolidays := h.Function(year)
		h.cache[year] = &yearHolidays
	}

	return *h.cache[year]
}

func easter(year int) time.Time {
	if _, ok := easterCache[year]; !ok {
		easter, _ := eastertime.CatholicByYear(year)
		easterCache[year] = &easter
	}

	return *easterCache[year]
}
