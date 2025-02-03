package tags

import (
	"errors"
	"golang-echo/jwt_token"
	"net/http"

	"github.com/labstack/echo/v4"
)

type GetUserTagResponse struct {
	Tags  []string `json:"tags,omitempty"`
	Error string   `json:"error,omitempty"`
}

func (t *TagHandler) GetUserTag(c echo.Context) error {
	claims, ok := c.Get("user").(*jwt_token.Claims)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "user claims not found")
	}
	userID := claims.UserID

	tags, err := t.service.GetUserTag(userID)
	if err != nil {
		switch {
		case errors.Is(err, ErrUserNotFound):
			return c.JSON(http.StatusNotFound, GetUserTagResponse{
				Error: "user not found",
			})
		case errors.Is(err, ErrTransactionFailed):
			return c.JSON(http.StatusInternalServerError, GetUserTagResponse{
				Error: "internal server error",
			})
		case errors.Is(err, ErrTag):
			return c.JSON(http.StatusInternalServerError, GetUserTagResponse{
				Error: "failed to get user tags",
			})
		default:
			return c.JSON(http.StatusInternalServerError, GetUserTagResponse{
				Error: "unexpected error occurred",
			})
		}
	}
	return c.JSON(http.StatusOK, GetUserTagResponse{
		Tags: tags,
	})
}

func (t *TagService) GetUserTag(myID int) ([]string, error) {
	tx, err := t.db.Begin()
	if err != nil {
		return nil, ErrTransactionFailed
	}
	defer tx.Rollback() // エラーが発生した場合はロールバック

	tags, err := getUserTags(tx, myID)
	if err != nil {
		return nil, ErrTag
	}
	if err := tx.Commit(); err != nil {
		return nil, ErrTransactionFailed
	}
	return tags, nil
}
