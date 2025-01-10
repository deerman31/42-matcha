package browse

import (
	"golang-echo/jwt_token"
	"net/http"

	"github.com/labstack/echo/v4"
)

type BrowseHandler struct {
	service *BrowseService
}

func NewBrowseHandler(service *BrowseService) *BrowseHandler {
	return &BrowseHandler{service: service}
}

func (b *BrowseHandler) GetBrowseUser(c echo.Context) error {
	claims, ok := c.Get("user").(*jwt_token.Claims)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "user claims not found")
	}
	req := new(BrowseRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, BrowseResponse{Error: "Invalid request body"})
	}
	// validationをここで行う
	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, BrowseResponse{Error: err.Error()})
	}
	users, err := b.service.GetBrowseUser(*req, claims.UserID)
	if err != nil {
		switch err {
		case ErrUserNotFound:
			return c.JSON(http.StatusNotFound, BrowseResponse{Error: "User not found"})
		default:
			return c.JSON(http.StatusInternalServerError, BrowseResponse{Error: "Internal server error"})
		}
	}
	return c.JSON(http.StatusOK, BrowseResponse{UserInfos: users})
}
