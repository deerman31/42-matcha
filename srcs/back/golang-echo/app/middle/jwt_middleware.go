package middle

import (
	"golang-echo/app/schemas"
	"golang-echo/app/utils/jwt_token"
	"net/http"

	"github.com/labstack/echo/v4"
)

func JWTMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			tokenString, err := jwt_token.GetAuthToken(c)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, schemas.ErrorResponse{
					Error: err.Error(),
				})
			}

			claims, err := jwt_token.ParseAndValidateAccessToken(tokenString)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, schemas.ErrorResponse{
					Error: err.Error(),
				})
			}

			// コンテキストにユーザー情報を設定
			c.Set("user", claims)

			return next(c)
		}
	}
}
