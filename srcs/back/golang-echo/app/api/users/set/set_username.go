package set

import (
	"database/sql"
	"fmt"
	"golang-echo/app/utils/jwt_token"
	"net/http"

	"github.com/labstack/echo/v4"
)

const (
	setUserNameSuccessMessage = "UserName updated successfully"
	setUserNameQuery          = `
        UPDATE users
        SET username = $1
        WHERE id = $2
    `
	checkDuplicateUsernameQuery = `
        SELECT EXISTS(
			SELECT 1 FROM users
			WHERE username = $1
		) as username_exists
	`
)

type setUserNameRequest struct {
	UserName string `json:"username" validate:"required,username"`
}

func SetUserName(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		claims, ok := c.Get("user").(*jwt_token.Claims)
		if !ok {
			return echo.NewHTTPError(http.StatusInternalServerError, "user claims not found")
		}
		userID := claims.UserID

		req := new(setUserNameRequest)
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

		if status, err := checkDuplicateUsername(tx, req.UserName); err != nil {
			return c.JSON(status, map[string]string{"error": err.Error()})
		}
		if status, err := executeUserNameUpdate(tx, req.UserName, userID); err != nil {
			return c.JSON(status, map[string]string{"error": err.Error()})
		}
		// トランザクションのコミット
		if err = tx.Commit(); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not commit transaction"})
		}

		return c.JSON(http.StatusOK, map[string]string{"message": setUserNameSuccessMessage})
	}
}

func executeUserNameUpdate(tx *sql.Tx, username string, userID int) (int, error) {
	result, err := tx.Exec(setUserNameQuery, username, userID)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	// 更新が成功したか確認
	rows, err := result.RowsAffected()
	if err != nil {
		return http.StatusInternalServerError, err
	}
	// userが見つからなかった場合
	if rows == 0 {
		return http.StatusNotFound, fmt.Errorf("User not found")
	}
	return 0, nil
}

func checkDuplicateUsername(tx *sql.Tx, username string) (int, error) {
	if username == "" {
		return http.StatusBadRequest, fmt.Errorf("username cannot be empty")
	}
	var usernameExists bool
	err := tx.QueryRow(checkDuplicateUsernameQuery, username).Scan(&usernameExists)

	switch {
	case err == sql.ErrNoRows:
		return http.StatusNotFound, fmt.Errorf("database query returned no rows")
	case err != nil:
		return http.StatusInternalServerError, fmt.Errorf("database error while checking username: %w", err)
	case usernameExists:
		return http.StatusConflict, fmt.Errorf("username %q is already taken", username)
	default:
		return http.StatusOK, nil
	}
}
