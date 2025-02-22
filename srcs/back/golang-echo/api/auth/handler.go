package auth

import (
	"golang-echo/pkg/errors"
	"golang-echo/pkg/jwt_token"
	"golang-echo/pkg/response"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	service *AuthService
}

func newAuthHandler(service *AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

func (a *AuthHandler) registerHandler(c echo.Context) error {
	req := new(RegisterRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{Success: false, Error: response.InvalidRequestMessage})
	}
	if req.Password != req.RePassword {
		return c.JSON(http.StatusBadRequest, response.Response{Success: false, Error: response.PasswordNoMatchMessage})
	}
	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{Success: false, Error: err.Error()})
	}
	err := a.service.registerService(req)
	if err != nil {
		switch err {
		case errors.ErrUserNameEmailConflict:
			return c.JSON(http.StatusConflict, response.Response{Success: false, Error: err.Error()})
		default:
			return c.JSON(http.StatusInternalServerError, response.Response{Success: false, Error: response.ServErrMessage})
		}
	}
	return c.JSON(http.StatusCreated, response.Response{Success: true, Data: response.RegisterData{Message: response.RegisterSuccessMessage}})
}

func (a *AuthHandler) loginHandler(c echo.Context) error {
	req := new(LoginRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{Success: false, Error: response.InvalidRequestMessage})
	}
	// validationをここで行う
	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, response.Response{Success: false, Error: err.Error()})
	}
	user, accessToken, err := a.service.loginService(req)
	if err != nil {
		switch err {
		case errors.ErrUserNotFound:
			return c.JSON(http.StatusNotFound, response.Response{Success: false, Error: err.Error()})
		case errors.ErrStatusForbidden:
			return c.JSON(http.StatusForbidden, response.Response{Success: false, Error: err.Error()})
		case errors.ErrPasswordUnauthorized:
			return c.JSON(http.StatusForbidden, response.Response{Success: false, Error: err.Error()})
		default:
			return c.JSON(http.StatusInternalServerError, response.Response{Success: false, Error: response.ServErrMessage})
		}
	}
	return c.JSON(http.StatusOK, response.Response{Success: true, Data: response.LoginData{IsPreparation: user.isPreparation, AccessToken: accessToken}})
}

func (a *AuthHandler) logoutHandler(c echo.Context) error {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	// Authorizationヘッダーを取得
	tokenString, err := jwt_token.GetAuthToken(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, response.Response{Success: false, Error: err.Error()})
	}
	claims, err := verifyTokenClaims(tokenString, secretKey)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, response.Response{Success: false, Error: err.Error()})
	}
	userID := claims.UserID
	err = a.service.logoutService(userID)
	if err != nil {
		switch err {
		case errors.ErrUserNotFound:
			return c.JSON(http.StatusNotFound, response.Response{Success: false, Error: err.Error()})
		default:
			return c.JSON(http.StatusInternalServerError, response.Response{Success: false, Error: response.ServErrMessage})
		}
	}
	return c.JSON(http.StatusInternalServerError, response.Response{Success: true, Data: response.LogoutData{Message: response.LogoutSuccessMessage}})
}
