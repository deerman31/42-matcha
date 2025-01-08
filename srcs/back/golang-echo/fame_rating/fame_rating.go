package famerating

import (
	"database/sql"
	"golang-echo/jwt_token"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetFameRating(db *sql.DB) echo.HandlerFunc {
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



		// トランザクションをコミット
		if err = tx.Commit(); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not commit transaction"})
		}

		/*
		1. 1ヶ月でどれだけ自分のプロフィールが閲覧されてたか　+ *1
		2. 1ヶ月でLikeされたか + *3
		3. friendの数 + *5
		4. blockされた数 - *5
		5 偽アカウントとして報告された数 -*5
		*/



	}
}
