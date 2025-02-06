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

func (a *AuthHandler) Login(c echo.Context) error {
	req := new(LoginRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}
	// validationをここで行う
	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	user, accessToken, err := a.service.Authenticate(req)
	if err != nil {
		switch err {
		case ErrUserNotFound:
			return c.JSON(http.StatusNotFound, LoginResponse{Error: "User not found"})
		case ErrStatusForbidden:
			return c.JSON(http.StatusForbidden, LoginResponse{Error: "Email not verified"})
		case ErrPasswordUnauthorized:
			return c.JSON(http.StatusForbidden, LoginResponse{Error: "Invalid password"})
		default:
			return c.JSON(http.StatusInternalServerError, LoginResponse{Error: "Internal server error"})

		}
	}
	return c.JSON(http.StatusOK, LoginResponse{IsPreparation: user.isPreparation, AccessToken: accessToken})
}

func (a *AuthService) Authenticate(req *LoginRequest) (*User, string, error) {
	// トランザクションを開始
	tx, err := a.db.Begin()
	if err != nil {
		return nil, "", ErrTransactionFailed
	}
	defer tx.Rollback() // エラーが発生した場合はロールバック

	user, status, err := searchUserDB(tx, req.Username)
	if err != nil {
		if status == http.StatusNotFound {
			return nil, "", ErrUserNotFound
		} else {
			return nil, "", ErrTransactionFailed
		}
	}
	// userがメールで認証済みかどうか確認
	if !user.isRegistered {
		return nil, "", ErrStatusForbidden
	}

	if bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)) != nil {
		return nil, "", ErrPasswordUnauthorized
	}

	accessToken, err := jwt_token.GenerateAccessToken(user.ID)
	if err != nil {
		return nil, "", ErrTransactionFailed
	}

	if err := updateUserStatusOn(tx, user.ID); err != nil {
		return nil, "", err
	}

	// トランザクションのコミット
	if err = tx.Commit(); err != nil {
		return nil, "", ErrTransactionFailed
	}
	return user, accessToken, nil
}

func updateUserStatusOn(tx *sql.Tx, userID int) error {
	const updateUserOnlineStatusQuery = `
        UPDATE users 
        SET is_online = TRUE 
        WHERE id = $1
	`
	result, err := tx.Exec(updateUserOnlineStatusQuery, userID)
	if err != nil {
		return ErrTransactionFailed
	}
	// 更新が成功したか確認
	rows, err := result.RowsAffected()
	if err != nil {
		return ErrTransactionFailed
	}
	// userが見つからなかった場合
	if rows == 0 {
		return ErrUserNotFound
	}
	return nil
}

func searchUserDB(tx *sql.Tx, username string) (*User, int, error) {
	// ユーザー名からユーザー情報を取得するクエリ
	const selectUserByUsernameQuery = `
        SELECT id, username, password_hash, is_online, is_registered, is_preparation
        FROM users 
        WHERE username = $1
        LIMIT 1
    `

	user := &User{}
	if err := tx.QueryRow(selectUserByUsernameQuery, username).Scan(
		&user.ID,
		&user.Username,
		&user.PasswordHash,
		&user.isOnline,
		&user.isRegistered,
		&user.isPreparation,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, http.StatusNotFound, errors.New("user not found")
		}
		return nil, http.StatusInternalServerError, fmt.Errorf("database error: %v", err)
	}
	return user, 0, nil
}
