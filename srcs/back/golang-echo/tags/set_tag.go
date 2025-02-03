package tags

import (
	"errors"
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

type SetTagRequest struct {
	Tag string `json:"tag" validate:"required,tag"`
}

type SetTagResponse struct {
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}

func (t *TagHandler) SetTag(c echo.Context) error {
	claims, ok := c.Get("user").(*jwt_token.Claims)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "user claims not found")
	}
	userID := claims.UserID

	req := new(SetTagRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, SetTagResponse{
			Error: "Invalid request body",
		})
	}

	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, SetTagResponse{
			Error: err.Error(),
		})
	}

	tagName := cases.Title(language.Und).String(strings.ToLower(req.Tag))

	err := t.service.SetTag(userID, tagName)
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

	return c.JSON(http.StatusOK, SetTagResponse{
		Message: setTagSuccessMessage,
	})
}

func (t *TagService) SetTag(userID int, tagName string) error {
	tx, err := t.db.Begin()
	if err != nil {
		return ErrTransactionFailed
	}
	defer tx.Rollback()

	if err := addTag(tx, tagName); err != nil {
		return ErrTag
	}

	tagID, err := getTagIDByName(tx, tagName)
	if err != nil {
		return ErrTag
	}

	if err := addUserTag(tx, userID, tagID); err != nil {
		return ErrTag
	}

	if err := tx.Commit(); err != nil {
		return ErrTransactionFailed
	}

	return nil
}
