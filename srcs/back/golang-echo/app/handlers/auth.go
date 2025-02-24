package handlers

import (
	"golang-echo/app/schemas"
	"golang-echo/app/schemas/errors"
	"golang-echo/app/services"
	"golang-echo/app/utils/jwt_token"
	"net/http"

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
// @Tags auth
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
// @Tags auth
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
// @Tags auth
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} schemas.Response "ログアウト成功"
// @Failure 401 {object} schemas.ErrorResponse "認証エラー（無効なトークンまたは期限切れ）"
// @Failure 404 {object} schemas.ErrorResponse "ユーザーが見つかりません"
// @Failure 500 {object} schemas.ErrorResponse "サーバーエラー"
// @Router /api/auth/logout [post]
func (a *AuthHandler) LogoutHandler(c echo.Context) error {
	// Authorizationヘッダーを取得
	tokenString, err := jwt_token.GetAuthToken(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, schemas.ErrorResponse{Error: err.Error()})
	}
	claims, err := jwt_token.VerifyTokenClaims(tokenString)
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

// VerifyEmail godoc
// @Summary メールアドレス認証
// @Description メール認証用トークンを検証し、ユーザーアカウントを有効化します
// @Tags auth
// @Accept json
// @Produce json
// @Param token path string true "認証トークン"
// @Success 200 {object} schemas.Response "認証成功"
// @Failure 401 {object} schemas.ErrorResponse "トークン認証エラー（無効なトークンまたは期限切れ）"
// @Failure 404 {object} schemas.ErrorResponse "ユーザーが見つかりません"
// @Failure 500 {object} schemas.ErrorResponse "サーバーエラー"
// @Router /api/auth/verify-email/{token} [get]
func (a *AuthHandler) VerifyEmailHandler(c echo.Context) error {
	token := c.Param("token")
	err := a.service.VerifyEmailService(token)
	if err != nil {
		switch err {
		case errors.ErrTokenUnauthorized:
			return c.JSON(http.StatusUnauthorized, schemas.ErrorResponse{Error: err.Error()})
		case errors.ErrUserNotFound:
			return c.JSON(http.StatusNotFound, schemas.ErrorResponse{Error: err.Error()})
		default:
			return c.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Error: schemas.ServErrMessage})
		}
	}
	return c.JSON(http.StatusOK, schemas.Response{Message: schemas.VerifyEmailSuccessMessage})
}

// ResetPasswordEmail godoc
// @Summary PasswordをResetする
// @Description PasswordをResetするリンクのあるメールを送信
// @Tags auth
// @Accept json
// @Produce json
// @Param request body schemas.ResetPasswordEmailRequest true "送信するメールアドレス"
// @Success 200 {object} schemas.Response "認証成功"
// @Failure 404 {object} schemas.ErrorResponse "ユーザーが見つかりません"
// @Failure 500 {object} schemas.ErrorResponse "サーバーエラー"
// @Router /api/auth/reset-password-email [post]
func (a *AuthHandler) ResetPasswordEmailHandler(c echo.Context) error {
	req := new(schemas.ResetPasswordEmailRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, schemas.ErrorResponse{Error: schemas.InvalidRequestMessage})
	}
	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, schemas.ErrorResponse{Error: err.Error()})
	}
	err := a.service.ResetPasswordEmailService(req.Email)
	if err != nil {
		switch err {
		case errors.ErrUserNotFound:
			return c.JSON(http.StatusNotFound, schemas.ErrorResponse{Error: err.Error()})
		default:
			return c.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Error: schemas.ServErrMessage})
		}
	}
	return c.JSON(http.StatusOK, schemas.Response{Message: schemas.ResetPasswordEmailSend})
}

// パスワードリセットを送信する godoc
// @Summary パスワードをリセット
// @Description パスワードをリセット
// @Tags auth
// @Accept json
// @Produce json
// @Param token path string true "認証token"
// @Param request body schemas.ResetPasswordRequest true "パスワードのjson"
// @Success 201 {object} schemas.Response "変更成功"
// @Failure 400 {object} schemas.ErrorResponse "リクエスト不正"
// @Failure 401 {object} schemas.ErrorResponse "認証エラー（無効なトークンまたは期限切れ）"
// @Failure 409 {object} schemas.ErrorResponse "ユーザー名またはメールアドレスが既に使用されています"
// @Failure 500 {object} schemas.ErrorResponse "サーバーエラー"
// @Router /api/auth/reset-password/{token} [post]
func (a *AuthHandler) ResetPasswordHandler(c echo.Context) error {
	token := c.Param("token")
	req := new(schemas.ResetPasswordRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, schemas.ErrorResponse{Error: schemas.InvalidRequestMessage})
	}
	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, schemas.ErrorResponse{Error: err.Error()})
	}

	if req.Password != req.RePassword {
		return c.JSON(http.StatusBadRequest, schemas.ErrorResponse{Error: schemas.PasswordNoMatchMessage})
	}

	err := a.service.ResetPasswordService(token, req.Password)
	if err != nil {
		switch err {
		case errors.ErrUserNotFound:
			return c.JSON(http.StatusNotFound, schemas.ErrorResponse{Error: err.Error()})
		case errors.ErrTokenUnauthorized:
			return c.JSON(http.StatusUnauthorized, schemas.ErrorResponse{Error: err.Error()})
		default:
			return c.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Error: schemas.ServErrMessage})
		}
	}
	return c.JSON(http.StatusOK, schemas.Response{Message: schemas.ResetPasswordSuccess})
}
