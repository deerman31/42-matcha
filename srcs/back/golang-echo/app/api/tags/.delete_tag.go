package tags

import (
	"database/sql"
	"golang-echo/jwt_token"
	"net/http"

	"github.com/labstack/echo/v4"
)

type deleteTagRequest struct {
	Tag string `json:"tag" validate:"required,tag"`
}

func DeleteTag(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		claims, ok := c.Get("user").(*jwt_token.Claims)
		if !ok {
			return echo.NewHTTPError(http.StatusInternalServerError, "user claims not found")
		}
		userID := claims.UserID

		req := new(deleteTagRequest)
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

		// tagsからtag_idを取得する
		tagID, err := getTagIDByName(tx, req.Tag)
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Tag not found"})
		} else if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to get tag"})
		}

		if err := deleteUserTag(tx, userID, tagID); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		// exists, err := existsTagInUserTags(tx, tagID)
		// if err != nil {
		// 	return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		// }

		// if !exists {
		// 	if err := deleteTag(tx, tagID); err != nil {
		// 		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		// 	}
		// }

		// トランザクションをコミット
		if err = tx.Commit(); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not commit transaction"})
		}
		// 成功時のレスポンスを返す
		return c.JSON(http.StatusOK, map[string]string{"message": "Tag deleted successfully"})
	}
}
