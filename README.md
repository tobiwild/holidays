# golang holidays [![Build Status](https://travis-ci.org/tobiwild/holidays.svg?branch=master)](https://travis-ci.org/tobiwild/holidays)

With this GOLANG package it is possible to check for holidays on given [Time](https://golang.org/pkg/time/#Time)

## Installation

    go get github.com/tobiwild/holidays

## Usage

```go
package main

import (
	"fmt"
	. "github.com/tobiwild/holidays"
	"time"
)

func main() {
	SetHolidaysFunction(Holidays_DE)
	date := time.Date(2009, time.January, 1, 0, 0, 0, 0, time.UTC)

	holidays := GetHolidays(date)
	for _, holiday := range holidays {
		fmt.Println(date, "is", holiday.Name)
	}

	printHolidayInfo(date)
	date = date.AddDate(0, 0, 1)
	printHolidayInfo(date)
}

func printHolidayInfo(t time.Time) {
	if IsHoliday(t) {
		fmt.Println(t, "is a holiday")
	} else {
		fmt.Println(t, "is not a holiday")
	}
}
```

Output:

```
2009-01-01 00:00:00 +0000 UTC is Neujahrstag
2009-01-01 00:00:00 +0000 UTC is a holiday
2009-01-02 00:00:00 +0000 UTC is not a holiday
```

## Todos

At the moment there are only these functions available:

* Holidays_DE
* Holidays_DE_NRW

It would be nice to have some more. A good place to look is the [ruby holidays gem](https://github.com/holidays/holidays/tree/master/data).
