package holidays

import (
	"time"
)

func HolidaysBR(year int) YearHolidays {
	result := YearHolidays{}

	result.add(&YearDay{time.January, 1}, &Holiday{"Ano Novo"})
	result.add(&YearDay{time.April, 21}, &Holiday{"Dia de Tiradentes"})
	result.add(&YearDay{time.May, 1}, &Holiday{"Dia do Trabalho"})
	result.add(&YearDay{time.June, 3}, &Holiday{"Corpus Christi"})
	result.add(&YearDay{time.September, 7}, &Holiday{"Independência do Brasil"})
	result.add(&YearDay{time.October, 12}, &Holiday{"Nossa Senhora Aparecida"})
	result.add(&YearDay{time.November, 2}, &Holiday{"Dia de Finados"})
	result.add(&YearDay{time.November, 15}, &Holiday{"Proclamação da República"})
	result.add(&YearDay{time.December, 25}, &Holiday{"Natal"})

	easter := easter(year)

	result.addFromTime(easter.AddDate(0, 0, 0), &Holiday{"Domingo de Páscoa"})
	result.addFromTime(easter.AddDate(0, 0, -2), &Holiday{"Sexta-Feira da Paixão"})
	result.addFromTime(easter.AddDate(0, 0, -47).UTC(), &Holiday{"Carnaval"})
	return result
}
