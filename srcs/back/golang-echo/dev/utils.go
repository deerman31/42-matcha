package dev

import (
	"fmt"
	"time"

	"golang.org/x/exp/rand"
)

func getImagePath(gender GenderType) string {
	num := getRandomNumber(5)
	if gender == GMale {
		return fmt.Sprintf("/home/appuser/uploads/images1/male%d", num) + ".png"
	} else {
		return fmt.Sprintf("/home/appuser/uploads/images1/famale%d", num) + ".png"
	}
}

func getSexuality() string {
	switch getRandomNumber(3) {
	case 1:
		return string(SMale)
	case 2:
		return string(SFamale)
	default:
		return string(SMaleFamale)
	}
}

func getArea() string {
	prefectures := []string{
		"Hokkaido", "Aomori", "Iwate", "Miyagi", "Akita", "Yamagata", "Fukushima",
		"Ibaraki", "Tochigi", "Gunma", "Saitama", "Chiba", "Tokyo", "Kanagawa",
		"Niigata", "Toyama", "Ishikawa", "Fukui", "Yamanashi", "Nagano",
		"Gifu", "Shizuoka", "Aichi", "Mie",
		"Shiga", "Kyoto", "Osaka", "Hyogo", "Nara", "Wakayama",
		"Tottori", "Shimane", "Okayama", "Hiroshima", "Yamaguchi",
		"Tokushima", "Kagawa", "Ehime", "Kochi",
		"Fukuoka", "Saga", "Nagasaki", "Kumamoto", "Oita", "Miyazaki", "Kagoshima", "Okinawa",
	}
	return prefectures[rand.Intn(len(prefectures))]
}

func getBirthDate() string {
	start := time.Date(1964, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2006, 12, 31, 0, 0, 0, 0, time.UTC)

	diff := end.Unix() - start.Unix()
	randomSeconds := rand.Int63n(diff)

	date := start.Add(time.Duration(randomSeconds) * time.Second)
	return date.Format("2006-01-02")
}

func getUserName(gender GenderType) string {
	var username string
	switch getRandomNumber(5) {
	case 1:
		if gender == GMale {
			username = string(Takashi)
		} else {
			username = string(Yui)
		}
	case 2:
		if gender == GMale {
			username = string(Yutaka)
		} else {
			username = string(Mayumi)
		}
	case 3:
		if gender == GMale {
			username = string(Koji)
		} else {
			username = string(Miyu)
		}
	case 4:
		if gender == GMale {
			username = string(Kai)
		} else {
			username = string(Meiko)
		}
	case 5:
		if gender == GMale {
			username = string(Shin)
		} else {
			username = string(Keiko)
		}
	}
	return username

}

func getRandomNumber(randomNum int) int {
	return rand.Intn(randomNum) + 1
}
