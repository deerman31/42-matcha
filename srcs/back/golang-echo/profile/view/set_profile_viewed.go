package view

import (
	"database/sql"
	"fmt"
	"golang-echo/jwt_token"
	"net/http"

	"github.com/labstack/echo/v4"
)

const (
	//query = `SELECT id FROM users WHERE username = $1;`
	getUserIDQuery = `SELECT id FROM users WHERE username = $1;`

	upsertProfileViewQuery = `
        INSERT INTO profile_views (viewer_id, viewed_id)
        VALUES ($1, $2)
        ON CONFLICT (viewer_id, viewed_id) 
        DO UPDATE SET viewed_at = CURRENT_TIMESTAMP
        WHERE profile_views.viewer_id = EXCLUDED.viewer_id 
        AND profile_views.viewed_id = EXCLUDED.viewed_id`
)

type setProfileViewedRequest struct {
	UserName string `json:"username" validate:"required,username"`
}

func SetProfileViewed(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		claims, ok := c.Get("user").(*jwt_token.Claims)
		if !ok {
			return echo.NewHTTPError(http.StatusInternalServerError, "user claims not found")
		}
		userID := claims.UserID
		req := new(setProfileViewedRequest)
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

		// 閲覧した相手ユーザーのusernameから相手のuser_idを取得する
		var viewedID int
		if err = tx.QueryRow(getUserIDQuery, req.UserName).Scan(&viewedID); err != nil {
			if err == sql.ErrNoRows {
				return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
			}
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": fmt.Sprintf("failed to query user tags: %v", err)})
		}
		if userID == viewedID {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Cannot record profile view for your own profile"})
		}

		if _, err := tx.Exec(upsertProfileViewQuery, userID, viewedID); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": fmt.Sprintf("failed to upsert profile view: %v", err)})
		}

		// トランザクションのコミット
		if err = tx.Commit(); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not commit transaction"})
		}

		return c.JSON(http.StatusCreated, map[string]string{"message": "Profile viewed created successfully."})
	}
}
