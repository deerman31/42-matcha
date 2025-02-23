package validations

import (
	"regexp"

	"github.com/go-playground/validator"
)

func validateUsername(fl validator.FieldLevel) bool {
	username := fl.Field().String()

	// 文字数が3~30文字でなければfalse
	if len(username) < 3 || len(username) > 30 {
		return false
	}
	// 英数字とアンダースコアのみ許可
	matched, _ := regexp.MatchString("^[a-zA-Z0-9_]+$", username)
	return matched
}
