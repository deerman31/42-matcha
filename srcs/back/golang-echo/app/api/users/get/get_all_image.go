package get

import (
	"database/sql"
	"golang-echo/app/utils"
	"golang-echo/app/utils/jwt_token"
	"golang-echo/app/utils/query"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAllImage(db *sql.DB) echo.HandlerFunc {
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
		allImagePath, err := query.GetAllImagePath(tx, userID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Failed to get existing image path",
			})
		}

		retImages, err := utils.SetAllImageURI(allImagePath)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": err.Error(),
			})
		}

		// トランザクションのコミット
		if err = tx.Commit(); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Could not commit transaction",
			})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"all_image": retImages,
		})
	}
}
