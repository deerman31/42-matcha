package validations

import (
	"time"

	"github.com/go-playground/validator"
)

func validateBirthdate(fl validator.FieldLevel) bool {
	birthDate := fl.Field().String()

	if _, err := time.Parse("2006-01-02", birthDate); err != nil {
		return false
	}
	return true
}
