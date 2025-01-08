package like

import (
	"database/sql"
	"fmt"
	"golang-echo/jwt_token"
	"net/http"

	"github.com/labstack/echo/v4"
)

type unLikeRequest struct {
	UserName string `json:"username" validate:"required,username"`
}

func UnLike(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		claims, ok := c.Get("user").(*jwt_token.Claims)
		if !ok {
			return echo.NewHTTPError(http.StatusInternalServerError, "user claims not found")
		}
		userID := claims.UserID
		req := new(unLikeRequest)
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

		var likedID int
		if err = tx.QueryRow(getUserIDQuery, req.UserName).Scan(&likedID); err != nil {
			if err == sql.ErrNoRows {
				return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
			}
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": fmt.Sprintf("failed to query: %v", err)})
		}
		if userID == likedID {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Cannot unlike your own profile"})
		}

		result, err := tx.Exec(deleteLikeQuery, userID, likedID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to remove existing like"})
		}
		rowsAffected, err := result.RowsAffected()
        if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to confirm like removal"})
		}
		if rowsAffected == 0 {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Like not found"})
		}

		if err = tx.Commit(); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to commit transaction"})
		}
		return c.JSON(http.StatusOK, map[string]string{"message": "Successfully unliked the user"})
	}
}
