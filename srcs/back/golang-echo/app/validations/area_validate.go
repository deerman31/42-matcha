package validations

import (
	"github.com/go-playground/validator"
)

func validateArea(fl validator.FieldLevel) bool {
	eria := fl.Field().String()

	validErias := []string{
		"Hokkaido", "Aomori", "Iwate", "Miyagi", "Akita", "Yamagata", "Fukushima",
		"Ibaraki", "Tochigi", "Gunma", "Saitama", "Chiba", "Tokyo", "Kanagawa",
		"Niigata", "Toyama", "Ishikawa", "Fukui", "Yamanashi", "Nagano",
		"Gifu", "Shizuoka", "Aichi", "Mie",
		"Shiga", "Kyoto", "Osaka", "Hyogo", "Nara", "Wakayama",
		"Tottori", "Shimane", "Okayama", "Hiroshima", "Yamaguchi",
		"Tokushima", "Kagawa", "Ehime", "Kochi",
		"Fukuoka", "Saga", "Nagasaki", "Kumamoto", "Oita", "Miyazaki", "Kagoshima", "Okinawa",
	}

	for _, validEria := range validErias {
		if eria == validEria {
			return true
		}
	}
	return false
}
