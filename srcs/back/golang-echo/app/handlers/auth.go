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

// Register godoc
// @Summary ユーザー登録
// @Description 新規ユーザーを登録します。ユーザー名、メールアドレス、パスワードが必要です。
// @Tags 認証
// @Accept json
// @Produce json
// @Param request body schemas.RegisterRequest true "登録情報"
// @Success 201 {object} schemas.Response "登録成功"
// @Failure 400 {object} schemas.ErrorResponse "リクエスト不正"
// @Failure 409 {object} schemas.ErrorResponse "ユーザー名またはメールアドレスが既に使用されています"
// @Failure 500 {object} schemas.ErrorResponse "サーバーエラー"
// @Router /api/auth/register [post]
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

// Login godoc
// @Summary ログイン
// @Description ユーザー認証を行い、アクセストークンを発行します。
// @Tags 認証
// @Accept json
// @Produce json
// @Param request body schemas.LoginRequest true "ログイン情報"
// @Success 200 {object} schemas.LoginResponse "ログイン成功"
// @Failure 400 {object} schemas.ErrorResponse "リクエスト不正"
// @Failure 403 {object} schemas.ErrorResponse "メール未認証またはパスワード不一致"
// @Failure 404 {object} schemas.ErrorResponse "ユーザーが見つかりません"
// @Failure 500 {object} schemas.ErrorResponse "サーバーエラー"
// @Router /api/auth/login [post]
func (a *AuthHandler) LoginHandler(c echo.Context) error {
	req := new(schemas.LoginRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, schemas.ErrorResponse{Error: schemas.InvalidRequestMessage})
	}
	// validationをここで行う
	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, schemas.ErrorResponse{Error: err.Error()})
	}
	isPreparation, accessToken, err := a.service.LoginService(req)
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
	return c.JSON(http.StatusOK, schemas.LoginResponse{IsPreparation: isPreparation, AccessToken: accessToken})
}

// Logout godoc
// @Summary ログアウト
// @Description ユーザーをログアウトし、アクセストークンを無効化します。
// @Tags 認証
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} schemas.Response "ログアウト成功"
// @Failure 401 {object} schemas.ErrorResponse "認証エラー（無効なトークンまたは期限切れ）"
// @Failure 404 {object} schemas.ErrorResponse "ユーザーが見つかりません"
// @Failure 500 {object} schemas.ErrorResponse "サーバーエラー"
// @Router /api/auth/logout [post]
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
	return c.JSON(http.StatusOK, schemas.Response{Message: schemas.LogoutSuccessMessage})
}
