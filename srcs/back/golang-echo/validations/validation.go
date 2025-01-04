package validations

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validator *validator.Validate
}

func NewValidator() *CustomValidator {
	v := validator.New()

	v.RegisterValidation("username", validateUsername)
	v.RegisterValidation("password", validatePassword)
	v.RegisterValidation("repassword", validatePassword)
	//v.RegisterValidation("last_name", validateName)
	v.RegisterValidation("name", validateName)
	v.RegisterValidation("area", validateArea)
	v.RegisterValidation("birthdate", validateBirthdate)
	v.RegisterValidation("self_intro", validateSelfIntro)

	return &CustomValidator{validator: v}
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

// バリデーションエラーのフォーマット
func formatValidationError(err error) error {
	var errMsgs []string

	for _, err := range err.(validator.ValidationErrors) {
		switch err.Tag() {
		case "required":
			errMsgs = append(errMsgs, fmt.Sprintf("%s is required", err.Field()))
		case "email":
			errMsgs = append(errMsgs, "Invalid email format")
		case "min":
			errMsgs = append(errMsgs, fmt.Sprintf("%s must be at least %s characters", err.Field(), err.Param()))
		case "max":
			errMsgs = append(errMsgs, fmt.Sprintf("%s must be at most %s characters", err.Field(), err.Param()))
		case "username":
			errMsgs = append(errMsgs, "Username must contain only alphanumeric characters and underscores")
		case "password":
			errMsgs = append(errMsgs, "Password must contain at least one uppercase letter, one lowercase letter, one number, and one special character")
		case "name":
			errMsgs = append(errMsgs, fmt.Sprintf("%s must contain only letters, spaces, and hyphens", err.Field()))
		case "oneof":
			errMsgs = append(errMsgs, fmt.Sprintf("%s must be one of: %s", err.Field(), err.Param()))
		case "eria":
			errMsgs = append(errMsgs, "Invalid prefecture")
		case "birthdate":
			errMsgs = append(errMsgs, "Invalid date")
		case "self_intro":
			errMsgs = append(errMsgs, "Invalid self intro")
		}
	}

	return fmt.Errorf("Validation failed: %s", strings.Join(errMsgs, "; "))
}