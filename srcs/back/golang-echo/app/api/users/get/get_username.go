package get

import (
	"database/sql"
	"fmt"
	"golang-echo/app/utils/jwt_token"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUserName(db *sql.DB) echo.HandlerFunc {
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
		const query = `SELECT username FROM users WHERE id = $1;`
		// クエリを実行
		var userName string
		if err = tx.QueryRow(query, userID).Scan(&userName); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": fmt.Sprintf("failed to query user tags: %v", err)})
		}

		// トランザクションをコミット
		if err = tx.Commit(); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not commit transaction"})
		}
		return c.JSON(http.StatusOK, map[string]string{"username": userName})
	}
}
