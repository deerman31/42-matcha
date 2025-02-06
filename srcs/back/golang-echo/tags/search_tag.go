package tags

import (
	"database/sql"
	"errors"
	"fmt"
	"golang-echo/jwt_token"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func (t *TagHandler) SearchTag(c echo.Context) error {
	// claims, ok := c.Get("user").(*jwt_token.Claims)
	_, ok := c.Get("user").(*jwt_token.Claims)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "user claims not found")
	}
	//userID := claims.UserID
	req := new(SearchTagRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, SearchTagResponse{
			Error: "Invalid request body",
		})
	}
	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, SearchTagResponse{
			Error: err.Error(),
		})
	}

	tagName := cases.Title(language.Und).String(strings.ToLower(req.TagName))

	tags, err := t.service.SearchTag(tagName)
	if err != nil {
		switch {
		case errors.Is(err, ErrTransactionFailed):
			return c.JSON(http.StatusInternalServerError, SetTagResponse{
				Error: "Internal server error",
			})
		case errors.Is(err, ErrTag):
			return c.JSON(http.StatusInternalServerError, SetTagResponse{
				Error: "Failed to set tag",
			})
		default:
			return c.JSON(http.StatusInternalServerError, SetTagResponse{
				Error: "Unexpected error occurred",
			})
		}
	}
	return c.JSON(http.StatusOK, SearchTagResponse{
		Tags: tags,
	})
}

func (t *TagService) SearchTag(tagName string) ([]string, error) {
	tx, err := t.db.Begin()
	if err != nil {
		return nil, ErrTransactionFailed
	}
	defer tx.Rollback() // エラーが発生した場合はロールバック

	tags, err := searchTagsHelper(tx, tagName)
	if err != nil {
		return nil, ErrTag
	}
	if err := tx.Commit(); err != nil {
		return nil, ErrTransactionFailed
	}
	return tags, nil
}

func searchTagsHelper(tx *sql.Tx, tagName string) ([]string, error) {
	const SearchTagsQuery = `
        SELECT tag_name 
		FROM tags
		WHERE tag_name LIKE $1 || '%'
		ORDER BY tag_name
    `
	// クエリを実行
	rows, err := tx.Query(SearchTagsQuery, tagName)
	if err != nil {
		return nil, fmt.Errorf("failed to query user tags: %v", err)
	}
	defer rows.Close()

	// タグ名を格納するスライス
	var tags []string

	// 結果を処理
	for rows.Next() {
		var tagName string
		if err := rows.Scan(&tagName); err != nil {
			return nil, fmt.Errorf("failed to scan tag name: %v", err)
		}
		tags = append(tags, tagName)
	}

	// rows.Next()のエラーチェック
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating tag rows: %v", err)
	}

	return tags, nil
}
