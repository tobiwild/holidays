package holidays

import (
	"testing"
	"time"
)

func TestHolidays_BR(t *testing.T) {
	SetHolidaysFunction(HolidaysBR)

	assertHoliday(t, 2017, time.January, 1, "Ano Novo")
	assertHoliday(t, 2017, time.April, 21, "Dia de Tiradentes")
	assertHoliday(t, 2017, time.May, 1, "Dia do Trabalho")
	assertHoliday(t, 2017, time.June, 15, "Corpus Christi")
	assertHoliday(t, 2017, time.September, 7, "Independência do Brasil")
	assertHoliday(t, 2017, time.October, 12, "Nossa Senhora Aparecida")
	assertHoliday(t, 2017, time.November, 2, "Dia de Finados")
	assertHoliday(t, 2017, time.November, 15, "Proclamação da República")
	assertHoliday(t, 2017, time.December, 25, "Natal")

	assertHoliday(t, 2017, time.April, 14, "Sexta-Feira da Paixão")
	assertHoliday(t, 2017, time.April, 16, "Domingo de Páscoa")
	assertHoliday(t, 2017, time.February, 28, "Carnaval")

}
