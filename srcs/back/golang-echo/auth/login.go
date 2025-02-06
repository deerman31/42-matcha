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
	result, err := tx.Exec(updateUserOnlineStatusQuery, user.ID)
	if err != nil {
		return nil, "", ErrTransactionFailed
	}
	// 更新が成功したか確認
	rows, err := result.RowsAffected()
	if err != nil {
		return nil, "", ErrTransactionFailed
	}
	// userが見つからなかった場合
	if rows == 0 {
		return nil, "", ErrUserNotFound
	}
	// トランザクションのコミット
	if err = tx.Commit(); err != nil {
		return nil, "", ErrTransactionFailed
	}
	return user, accessToken, nil
}

// func Login(db *sql.DB) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		req := new(LoginRequest)
// 		if err := c.Bind(req); err != nil {
// 			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
// 		}
// 		// validationをここで行う
// 		if err := c.Validate(req); err != nil {
// 			return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
// 		}
// 		// トランザクションを開始
// 		tx, err := db.Begin()
// 		if err != nil {
// 			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not start transaction"})
// 		}
// 		defer tx.Rollback() // エラーが発生した場合はロールバック
// 		user, status, err := searchUserDB(tx, req.Username)
// 		if err != nil {
// 			return c.JSON(status, map[string]string{"error": err.Error()})
// 		}
// 		// userがメールで認証済みかどうか確認
// 		if !user.isRegistered {
// 			return c.JSON(http.StatusForbidden, map[string]string{"error": "Email not verified"})
// 		}
// 		if bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)) != nil {
// 			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid password"})
// 		}

// 		accessToken, err := jwt_token.GenerateAccessToken(user.ID)
// 		if err != nil {
// 			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
// 		}
// 		result, err := tx.Exec(updateUserOnlineStatusQuery, user.ID)
// 		if err != nil {
// 			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
// 		}
// 		// 更新が成功したか確認
// 		rows, err := result.RowsAffected()
// 		if err != nil {
// 			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
// 		}
// 		// userが見つからなかった場合
// 		if rows == 0 {
// 			return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
// 		}
// 		// トランザクションのコミット
// 		if err = tx.Commit(); err != nil {
// 			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not commit transaction"})
// 		}
// 		return c.JSON(http.StatusOK, TokenResponse{IsPreparation: user.isPreparation, AccessToken: accessToken})
// 	}
// }

func searchUserDB(tx *sql.Tx, username string) (*User, int, error) {
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
