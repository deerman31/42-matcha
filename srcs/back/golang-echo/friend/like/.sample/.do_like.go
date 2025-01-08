package like

import (
	"database/sql"
	"fmt"
	"golang-echo/jwt_token"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

type doLikeRequest struct {
	UserName string `json:"username" validate:"required,username"`
}

const (
	getUserIDQuery = `SELECT id FROM users WHERE username = $1;`

	checkLikeExistsQuery = `
    SELECT EXISTS (
        SELECT 1 
        FROM user_likes 
        WHERE liker_id = $1 AND liked_id = $2
    )
`
	checkFriendExistsQuery = `
    SELECT EXISTS (
        SELECT 1 
        FROM user_friends 
        WHERE user_id1 = $1 AND user_id2 = $2
    )
`
	deleteLikeQuery = `
		DELETE FROM user_likes
		WHERE liker_id = $1 AND liked_id = $2
	`

	insertLikeQuery = `
	INSERT INTO user_likes (liker_id, liked_id)
	VALUES ($1, $2)
	`

	insertFriendQuery = `
	INSERT INTO user_friends (user_id1, user_id2)
	VALUES ($1, $2)
	`
)

/*
関数の目的: 相手ユーザーにlikeする
1 すでにfriendの場合はlikeは意味をなさないので何も処理をせず終わる
2 相手からlikeされている場合はのちの処理が変わるので 相手からlikeされているかを調べる
3 相手からlikeされている場合は既存のuser_likesテーブルを削除し、user_friendsに挿入する
likeされていない枚はuser_likesテーブルに挿入する
*/

func DoLike(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		claims, ok := c.Get("user").(*jwt_token.Claims)
		if !ok {
			return echo.NewHTTPError(http.StatusInternalServerError, "user claims not found")
		}
		userID := claims.UserID

		req := new(doLikeRequest)
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
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": fmt.Sprintf("failed to query user tags: %v", err)})
		}
		if userID == likedID {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Cannot record profile like for your own profile"})
		}

		// すでにfriendかどうかを調べるすでにfriendの相手にlikeを遅れるのはおかしいので
		friendExists, err := isFriend(tx, userID, likedID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to check friend status"})
		}
		if friendExists {
			return c.JSON(http.StatusConflict, map[string]string{"error": "Already friends with this user"})
		}

		// 相手からのlikeが存在するかをチェック
		var otherLikeExists bool
		if err := tx.QueryRow(checkLikeExistsQuery, likedID, userID).Scan(&otherLikeExists); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to check existing likes"})
		}

		if otherLikeExists {
			// 相手からのlikeが存在する場合には、friend関係を作成する。
			// まず、存在するlikeを削除
			if _, err := tx.Exec(deleteLikeQuery, likedID, userID); err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to remove existing like"})
			}
			// insert
			if _, err := tx.Exec(insertFriendQuery, min(userID, likedID), max(userID, likedID)); err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create friend relationship"})
			}
			if err = tx.Commit(); err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to commit transaction"})
			}
			return c.JSON(http.StatusCreated, map[string]string{"message": "Successfully matched! You are now friends."})
		} else {
			// 通常のlike処理
			if _, err := tx.Exec(insertLikeQuery, userID, likedID); err != nil {
				if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
					return c.JSON(http.StatusConflict, map[string]string{"error": "You have already liked this user"})
				}
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create like"})
			}
			if err = tx.Commit(); err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to commit transaction"})
			}
			return c.JSON(http.StatusCreated, map[string]string{"message": "Successfully liked the user"})
		}
	}
}

// すでにfriendかどうかを調べる関数
func isFriend(tx *sql.Tx, myID, otherID int) (bool, error) {
	var exists bool
	var ID1, ID2 int
	if myID < otherID {
		ID1 = myID
		ID2 = otherID
	} else {
		ID1 = otherID
		ID2 = myID
	}
	err := tx.QueryRow(checkFriendExistsQuery, ID1, ID2).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func min(num1, num2 int) int {
	if num1 < num2 {
		return num1
	} else {
		return num2
	}
}

func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	} else {
		return num2
	}
}
