package tags

import (
	"database/sql"
	"golang-echo/jwt_token"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUserTag(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		claims, ok := c.Get("user").(*jwt_token.Claims)
		if !ok {
			return echo.NewHTTPError(http.StatusInternalServerError, "user claims not found")
		}
		userID := claims.UserID

		// トランザクションを開始
		tx, err := db.Begin()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not start transaction"})
		}
		defer tx.Rollback() // エラーが発生した場合はロールバック
		tags, err := getUserTags(tx, userID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		// トランザクションをコミット
		if err = tx.Commit(); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not commit transaction"})
		}
		// タグのリストを返す
		return c.JSON(http.StatusOK, map[string][]string{"tags": tags})
	}
}