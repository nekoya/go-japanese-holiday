package main

import (
	"./holidays"
	"os"
	"time"
)

func main() {
	now := time.Now()
	today := time.Date(now.Year(), time.Month(now.Month()), now.Day(), 0, 0, 0, 0, time.Local)
	if holidays.IsWeekend(today) {
		os.Exit(0)
	}
	if in(holidays.GetHolidays(), today) {
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
