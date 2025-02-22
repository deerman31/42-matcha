package jwt_token

import (
	"fmt"
	"strings"

	"github.com/labstack/echo/v4"
)

func GetAuthToken(c echo.Context) (string, error) {
	authHeader := c.Request().Header.Get("Authorization")
	if authHeader == "" {
		return "", fmt.Errorf("authorization header is required")
	}
	strs := strings.Split(authHeader, " ")
	if len(strs) != 2 || strs[0] != "Bearer" {
		return "", fmt.Errorf("invalid authorization format")
	}
	return strs[1], nil
}
