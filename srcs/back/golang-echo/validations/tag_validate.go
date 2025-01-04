package validations

import (
	"regexp"

	"github.com/go-playground/validator"
)

func validateTag(fl validator.FieldLevel) bool {
	tag := fl.Field().String()

	// 文字数Check
	if len(tag) < 1 || len(tag) > 50 {
		return false
	}
	// 英数字のみ許可
	matched, _ := regexp.MatchString("^[a-zA-Z0-9]+$", tag)
	return matched
}
