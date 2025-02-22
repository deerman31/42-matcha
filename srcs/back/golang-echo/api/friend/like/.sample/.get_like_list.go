package like

import (
	"database/sql"
	"golang-echo/jwt_token"
	"net/http"

	"github.com/labstack/echo/v4"
)

const (
	getLikedUsersQuery = `
        SELECT DISTINCT u.username 
        FROM user_likes ul
        JOIN users u ON ul.liked_id = u.id
        WHERE ul.liker_id = $1
        ORDER BY ul.created_at DESC
    `

	getLikerUsersQuery = `
        SELECT DISTINCT u.username 
        FROM user_likes ul
        JOIN users u ON ul.liker_id = u.id
        WHERE ul.liked_id = $1
        ORDER BY ul.created_at DESC
    `
)

type LikeType string

const (
	likedUsers LikeType = "liked_users" // 自分がlikeしたユーザー
	likerUsers LikeType = "liker_users" // 自分をlikeしたユーザー
)

func getUserList(tx *sql.Tx, userID int, likeType LikeType) ([]string, error) {
	query := getLikedUsersQuery
	if likeType == likerUsers {
		query = getLikerUsersQuery
	}

	rows, err := tx.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var usernames []string
	for rows.Next() {
		var username string
		if err := rows.Scan(&username); err != nil {
			return nil, err
		}
		usernames = append(usernames, username)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return usernames, nil
}
func handleLikeList(db *sql.DB, likeType LikeType) echo.HandlerFunc {
	return func(c echo.Context) error {
		claims, ok := c.Get("user").(*jwt_token.Claims)
		if !ok {
			return echo.NewHTTPError(http.StatusInternalServerError, "user claims not found")
		}
		userID := claims.UserID

		tx, err := db.Begin()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not start transaction"})
		}
		defer tx.Rollback()

		usernames, err := getUserList(tx, userID, likeType)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		if err = tx.Commit(); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not commit transaction"})
		}

		return c.JSON(http.StatusOK, map[string][]string{string(likeType): usernames})
	}
}

// 公開用のハンドラー関数
func GetLikedUsers(db *sql.DB) echo.HandlerFunc {
	return handleLikeList(db, likedUsers)
}

func GetLikerUsers(db *sql.DB) echo.HandlerFunc {
	return handleLikeList(db, likerUsers)
}
