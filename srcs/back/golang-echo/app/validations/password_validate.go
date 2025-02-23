package validations

import (
	"regexp"

	"github.com/go-playground/validator"
)

func validatePassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	// 文字数Check
	if len(password) < 8 || len(password) > 30 {
		return false
	}

	// 半角大文字小文字の英字、数字、特定の記号のみを許可する正規表現, それ以外があるとfalse
	re := regexp.MustCompile(`^[a-zA-Z0-9!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]+$`)
	if !re.MatchString(password) {
		return false
	}

	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(password)
	hasSpecial := regexp.MustCompile(`[!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]`).MatchString(password)
	return hasLower && hasUpper && hasNumber && hasSpecial
}