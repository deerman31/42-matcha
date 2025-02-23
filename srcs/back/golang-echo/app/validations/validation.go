package validations

import (
	"fmt"

	"github.com/go-playground/validator"
)

type CustomValidator struct {
	validator *validator.Validate
}

func validateMinCommonTags(fl validator.FieldLevel) bool {
	minCommonTags := fl.Field().Int()
	if minCommonTags < 0 || minCommonTags > 5 {
		return false
	}
	return true
}

// min_fame_rating
func validateMinFameRating(fl validator.FieldLevel) bool {
	minFameRating := fl.Field().Int()
	if minFameRating < 0 || minFameRating > 5 {
		return false
	}
	return true
}

func NewValidator() *CustomValidator {
	v := validator.New()

	v.RegisterValidation("username", validateUsername)
	v.RegisterValidation("password", validatePassword)
	v.RegisterValidation("repassword", validatePassword)
	v.RegisterValidation("name", validateName)
	v.RegisterValidation("area", validateArea)
	v.RegisterValidation("birthdate", validateBirthdate)
	v.RegisterValidation("self_intro", validateSelfIntro)

	v.RegisterValidation("tag", validateTag)

	v.RegisterValidation("age_range", validateAgeRange)
	v.RegisterValidation("distance_range", validateDistanceRange)

	v.RegisterValidation("min_common_tags", validateMinCommonTags)
	v.RegisterValidation("min_fame_rating", validateMinFameRating)

	return &CustomValidator{validator: v}
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		formattedErr := formatValidationError(err)
		return formattedErr
	}
	return nil
}

func formatValidationError(err error) error {
	validationErrors := err.(validator.ValidationErrors)
	if len(validationErrors) == 0 {
		return nil
	}

	// 最初のエラーのみを取得
	firstErr := validationErrors[0]
	var errMsg string

	switch firstErr.Tag() {
	case "required":
		errMsg = fmt.Sprintf("%s is required", firstErr.Field())
	case "email":
		errMsg = "Invalid email format"
	case "min":
		errMsg = fmt.Sprintf("%s must be at least %s characters", firstErr.Field(), firstErr.Param())
	case "max":
		errMsg = fmt.Sprintf("%s must be at most %s characters", firstErr.Field(), firstErr.Param())
	case "username":
		errMsg = "Username must contain only alphanumeric characters and underscores"
	case "password":
		errMsg = "Password must contain at least one uppercase letter, one lowercase letter, one number, and one special character"
	case "name":
		errMsg = fmt.Sprintf("%s must contain only letters, spaces, and hyphens", firstErr.Field())
	case "oneof":
		errMsg = fmt.Sprintf("%s must be one of: %s", firstErr.Field(), firstErr.Param())
	case "eria":
		errMsg = "Invalid prefecture"
	case "birthdate":
		errMsg = "Invalid date"
	case "self_intro":
		errMsg = "Invalid self intro"
	case "tag":
		errMsg = "Invalid tag"
	case "age_range":
		errMsg = "Invalid age_range"
	case "distance_range":
		errMsg = "Invalid distance_range"
	case "min_common_tags":
		errMsg = "Invalid min_common_tags"
	case "min_fame_rating":
		errMsg = "Invalid min_fame_rating"
	}

	return fmt.Errorf("validation failed: %s", errMsg)
}
