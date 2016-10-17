package main

import (
	"./holidays"
	"os"
	"time"
)

func main() {
	t := time.Now()
	if holidays.IsWeekend(t) {
		os.Exit(0)
	}
	if in(holidays.GetHolidays(), t) {
		os.Exit(0)
	}
	os.Exit(1)
}

func in(haystack []time.Time, needle time.Time) bool {
	for _, t := range haystack {
		if t == needle {
			return true
		}
	}
	return false
}
