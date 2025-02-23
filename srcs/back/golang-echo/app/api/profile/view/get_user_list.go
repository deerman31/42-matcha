package view

import (
	"database/sql"
	"golang-echo/app/utils/jwt_token"
	"net/http"

	"github.com/labstack/echo/v4"
)

const (
	getViewedUsersQuery = `
        SELECT DISTINCT u.username 
        FROM profile_views pv
        JOIN users u ON pv.viewed_id = u.id
        WHERE pv.viewer_id = $1
        ORDER BY pv.viewed_at DESC
    `

	getViewerUsersQuery = `
        SELECT DISTINCT u.username 
        FROM profile_views pv
        JOIN users u ON pv.viewer_id = u.id
        WHERE pv.viewed_id = $1
        ORDER BY pv.viewed_at DESC
    `
)

type ViewType string

const (
	ViewedUsers ViewType = "viewed_users" // 自分が閲覧したユーザー
	ViewerUsers ViewType = "viewer_users" // 自分を閲覧したユーザー
)

func getUserList(tx *sql.Tx, userID int, viewType ViewType) ([]string, error) {
	query := getViewedUsersQuery
	if viewType == ViewerUsers {
		query = getViewerUsersQuery
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

func handleUserList(db *sql.DB, viewType ViewType) echo.HandlerFunc {
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

		usernames, err := getUserList(tx, userID, viewType)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		if err = tx.Commit(); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not commit transaction"})
		}

		return c.JSON(http.StatusOK, map[string][]string{string(viewType): usernames})
	}
}

// 元のハンドラー関数をラップする形で定義
func GetViewedUsers(db *sql.DB) echo.HandlerFunc {
	return handleUserList(db, ViewedUsers)
}

func GetViewerUsers(db *sql.DB) echo.HandlerFunc {
	return handleUserList(db, ViewerUsers)
}
