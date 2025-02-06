// delete_tag.go
package tags

import (
	"database/sql"
	"errors"
	"golang-echo/jwt_token"
	"net/http"

	"github.com/labstack/echo/v4"
)

const (
	deleteTagSuccessMessage = "Tag deleted successfully"
)



func (t *TagHandler) DeleteTag(c echo.Context) error {
	claims, ok := c.Get("user").(*jwt_token.Claims)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "user claims not found")
	}
	userID := claims.UserID

	req := new(DeleteTagRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, DeleteTagResponse{
			Error: "Invalid request body",
		})
	}

	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, DeleteTagResponse{
			Error: err.Error(),
		})
	}

	err := t.service.DeleteTag(userID, req.Tag)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return c.JSON(http.StatusNotFound, DeleteTagResponse{
				Error: "Tag not found",
			})
		case errors.Is(err, ErrTransactionFailed):
			return c.JSON(http.StatusInternalServerError, DeleteTagResponse{
				Error: "Internal server error",
			})
		case errors.Is(err, ErrTag):
			return c.JSON(http.StatusInternalServerError, DeleteTagResponse{
				Error: "Failed to delete tag",
			})
		default:
			return c.JSON(http.StatusInternalServerError, DeleteTagResponse{
				Error: "Unexpected error occurred",
			})
		}
	}

	return c.JSON(http.StatusOK, DeleteTagResponse{
		Message: deleteTagSuccessMessage,
	})
}

// service.go に追加
func (t *TagService) DeleteTag(userID int, tagName string) error {
	tx, err := t.db.Begin()
	if err != nil {
		return ErrTransactionFailed
	}
	defer tx.Rollback()

	tagID, err := getTagIDByName(tx, tagName)
	if err == sql.ErrNoRows {
		return sql.ErrNoRows
	} else if err != nil {
		return ErrTag
	}

	if err := deleteUserTag(tx, userID, tagID); err != nil {
		return ErrTag
	}

	if err := tx.Commit(); err != nil {
		return ErrTransactionFailed
	}

	return nil
}
