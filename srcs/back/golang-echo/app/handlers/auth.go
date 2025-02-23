package handlers

import (
	"golang-echo/app/cruds"
	"golang-echo/app/schemas"
	"golang-echo/app/schemas/errors"
	"golang-echo/app/services"
	"golang-echo/app/utils/jwt_token"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	service *services.AuthService
}

func NewAuthHandler(service *services.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

func (a *AuthHandler) RegisterHandler(c echo.Context) error {
	req := new(schemas.RegisterRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, schemas.ErrorResponse{Error: schemas.InvalidRequestMessage})
	}
	if req.Password != req.RePassword {
		return c.JSON(http.StatusBadRequest, schemas.ErrorResponse{Error: schemas.PasswordNoMatchMessage})
	}
	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, schemas.ErrorResponse{Error: err.Error()})
	}
	err := a.service.RegisterService(req)
	if err != nil {
		switch err {
		case errors.ErrUserNameEmailConflict:
			return c.JSON(http.StatusConflict, schemas.ErrorResponse{Error: err.Error()})
		default:
			return c.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Error: schemas.ServErrMessage})
		}
	}
	return c.JSON(http.StatusCreated, schemas.Response{Message: schemas.RegisterSuccessMessage})
}

func (a *AuthHandler) LoginHandler(c echo.Context) error {
	req := new(schemas.LoginRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, schemas.ErrorResponse{Error: schemas.InvalidRequestMessage})
	}
	// validationをここで行う
	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, schemas.ErrorResponse{Error: err.Error()})
	}
	user, accessToken, err := a.service.LoginService(req)
	if err != nil {
		switch err {
		case errors.ErrUserNotFound:
			return c.JSON(http.StatusNotFound, schemas.ErrorResponse{Error: err.Error()})
		case errors.ErrStatusForbidden:
			return c.JSON(http.StatusForbidden, schemas.ErrorResponse{Error: err.Error()})
		case errors.ErrPasswordUnauthorized:
			return c.JSON(http.StatusForbidden, schemas.ErrorResponse{Error: err.Error()})
		default:
			return c.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Error: schemas.ServErrMessage})
		}
	}
	return c.JSON(http.StatusOK, schemas.LoginResponse{IsPreparation: user.IsPreparation, AccessToken: accessToken})
}

func (a *AuthHandler) LogoutHandler(c echo.Context) error {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	// Authorizationヘッダーを取得
	tokenString, err := jwt_token.GetAuthToken(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, schemas.ErrorResponse{Error: err.Error()})
	}
	claims, err := cruds.VerifyTokenClaims(tokenString, secretKey)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, schemas.ErrorResponse{Error: err.Error()})
	}
	userID := claims.UserID
	err = a.service.LogoutService(userID)
	if err != nil {
		switch err {
		case errors.ErrUserNotFound:
			return c.JSON(http.StatusNotFound, schemas.ErrorResponse{Error: err.Error()})
		default:
			return c.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Error: schemas.ServErrMessage})
		}
	}
	return c.JSON(http.StatusInternalServerError, schemas.Response{Message: schemas.LogoutSuccessMessage})
}
