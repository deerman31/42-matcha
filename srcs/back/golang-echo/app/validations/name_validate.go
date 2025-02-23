package validations

import (
	"regexp"

	"github.com/go-playground/validator"
)

func validateName(fl validator.FieldLevel) bool {
	name := fl.Field().String()

	// 文字数Check
	if len(name) < 1 || len(name) > 50 {
		return false
	}

	// スペースと一般的な文字のみ許可
	matched, _ := regexp.MatchString("^[a-zA-Z\\s-]+$", name)
	return matched
}
