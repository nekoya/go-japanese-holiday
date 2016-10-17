package holidays

import (
	"strconv"
	"strings"
	"time"
)

func IsWeekend(t time.Time) bool {
	w := t.Weekday()
	if w == time.Sunday || w == time.Saturday {
		return true
	}
	return false
}

func GetHolidays() []time.Time {
	holidays := make([]time.Time, 0, 100)
	for _, l := range Holidays() {
		e := strings.Split(l, "/")
		year, _ := strconv.Atoi(e[0])
		month, _ := strconv.Atoi(e[1])
		day, _ := strconv.Atoi(e[2])
		t := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
		holidays = append(holidays, t)
	}
	return holidays
}
