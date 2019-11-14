package modules

import "time"

var JSONContentType = "application/json; charset=utf-8"

func MonthYearDiff(a, b time.Time) (years, months int) {
	m := a.Month()
	for a.Before(b) {
		a = a.Add(time.Hour * 24)
		m2 := a.Month()
		if m2 != m {
			months++
		}
		m = m2
	}
	years = months / 12
	months = months % 12
	return
}
