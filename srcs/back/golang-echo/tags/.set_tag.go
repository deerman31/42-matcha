package tags

import (
	"database/sql"
	"golang-echo/jwt_token"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const (
	setTagSuccessMessage = "Tag set successfully"
)

type setTagRequest struct {
	Tag string `json:"tag" validate:"required,tag"`
}

func SetTag(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		claims, ok := c.Get("user").(*jwt_token.Claims)
		if !ok {
			return echo.NewHTTPError(http.StatusInternalServerError, "user claims not found")
		}
		userID := claims.UserID
		req := new(setTagRequest)
		if err := c.Bind(req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
		}
		// validationをここで行う
		if err := c.Validate(req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		}

		//TagName := strings.Title(strings.ToLower(req.Tag))
		TagName := cases.Title(language.Und).String(strings.ToLower(req.Tag))


		// トランザクションを開始
		tx, err := db.Begin()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not start transaction"})
		}
		defer tx.Rollback() // エラーが発生した場合はロールバック

		// tagをtagsに追加
		if err := addTag(tx, TagName); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		// tagsからtag_idを取得する
		tagID, err := getTagIDByName(tx, TagName)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		// tagIDとuserIDを使って、user_tagに要素を追加
		if err := addUserTag(tx, userID, tagID); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		// トランザクションのコミット
		if err = tx.Commit(); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not commit transaction"})
		}
		return c.JSON(http.StatusOK, map[string]string{"message": setTagSuccessMessage})

	}
}