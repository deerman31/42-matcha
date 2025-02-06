package utils

import (
	"time"
)

func stringToTime(str string, layout string) time.Time {
	t, _ := time.Parse(layout, str)
	return t
}

func CalculateAgeFromBirthDate(birthDate string) int {
	date := stringToTime(birthDate[:10], "2006-01-02")
	now := time.Now()
	age := now.Year() - date.Year()

	if now.YearDay() < date.YearDay() {
		age--
	}
	return age
}
