package famerating

import (
	"database/sql"
	"golang-echo/jwt_token"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

/*
	1. 1ヶ月でどれだけ自分のプロフィールが閲覧されてたか　+ *1
	2. 1ヶ月でLikeされたか + *3
	3. friendの数 + *5
	4. blockされた数 - *5
	5 偽アカウントとして報告された数 -*5
*/

const (
	viewedPoint      = 1
	likedPoint       = 3
	friendPoint      = 5
	blockPoint       = 5
	fakeAccountPoint = 5
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

		vieweds, err := viewedCount(tx, userID, 30*24*time.Hour)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		likeds, err := likedCount(tx, userID, 30*24*time.Hour)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		friends, err := friendCount(tx, userID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		blocks, err := blockCount(tx, userID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		fakeAccounts, err := fakeAccountCount(tx, userID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		// トランザクションをコミット
		if err = tx.Commit(); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not commit transaction"})
		}
		rating := fameRatingCalculation(vieweds, likeds, friends, blocks, fakeAccounts)
		return c.JSON(http.StatusOK, map[string]int{"fame_rating": rating})

	}
}

func fameRatingCalculation(vieweds, likeds, friends, blocks, fakeAccounts int) int {
	point := (vieweds * viewedPoint) + (likeds * likedPoint) + (friends * friendPoint) - (blocks * blockPoint) - (fakeAccounts * fakeAccountPoint)

	if point >= 100 {
		return 5
	} else if point >= 80 {
		return 4
	} else if point >= 60 {
		return 3
	} else if point >= 40 {
		return 2
	} else if point >= 20 {
		return 1
	} else {
		return 0
	}
}
