package auth

import (
	"database/sql"
	"errors"
	"fmt"
	"golang-echo/jwt_token"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Username string `json:"username" validate:"required,username"`
	Password string `json:"password" validate:"required,password"`
}

// User はデータベースのユーザー情報を表す構造体
type User struct {
	ID           int
	Username     string
	PasswordHash string
	isOnline     bool
	isRegistered bool
}

// トークンのレスポンス用構造体を追加
type TokenResponse struct {
	AccessToken string `json:"access_token"`
}

const (
	// ユーザー名からユーザー情報を取得するクエリ
	selectUserByUsernameQuery = `
        SELECT id, username, password_hash, is_online, is_registered
        FROM users 
        WHERE username = $1
        LIMIT 1
    `

	// ユーザーのオンラインステータスを更新するクエリ
	updateUserOnlineStatusQuery = `
        UPDATE users 
        SET is_online = TRUE 
        WHERE id = $1
	`
)

func Login(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := new(LoginRequest)
		if err := c.Bind(req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
		}
		// validationをここで行う
		if err := c.Validate(req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		}
		// トランザクションを開始
		tx, err := db.Begin()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not start transaction"})
		}
		defer tx.Rollback() // エラーが発生した場合はロールバック
		user, status, err := searchUserDB(tx, req.Username)
		if err != nil {
			return c.JSON(status, map[string]string{"error": err.Error()})
		}
		// userがメールで認証済みかどうか確認
		if !user.isRegistered {
			return c.JSON(http.StatusForbidden, map[string]string{"error": "Email not verified"})
		}
		if bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)) != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid password"})
		}

		accessToken, err := jwt_token.GenerateAccessToken(user.ID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		result, err := tx.Exec(updateUserOnlineStatusQuery, user.ID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		// 更新が成功したか確認
		rows, err := result.RowsAffected()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		// userが見つからなかった場合
		if rows == 0 {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
		}
		// トランザクションのコミット
		if err = tx.Commit(); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not commit transaction"})
		}
		return c.JSON(http.StatusOK, TokenResponse{AccessToken: accessToken})
	}
}

func searchUserDB(tx *sql.Tx, username string) (*User, int, error) {
	user := &User{}
	if err := tx.QueryRow(selectUserByUsernameQuery, username).Scan(
		&user.ID,
		&user.Username,
		&user.PasswordHash,
		&user.isOnline,
		&user.isRegistered,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, http.StatusUnauthorized, errors.New("user not found")
		}
		return nil, http.StatusInternalServerError, fmt.Errorf("database error: %v", err)
	}
	return user, 0, nil
}
