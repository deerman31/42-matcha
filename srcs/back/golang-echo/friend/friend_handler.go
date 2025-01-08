package friend

import (
	"golang-echo/jwt_token"
	"net/http"

	"github.com/labstack/echo/v4"
)

type FriendHandler struct {
	friendService *FriendService
}

func NewFriendHandler(friendService *FriendService) *FriendHandler {
	return &FriendHandler{friendService: friendService}
}

func (f *FriendHandler) GetFriendList(c echo.Context) error {
	claims, ok := c.Get("user").(*jwt_token.Claims)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "user claims not found")
	}
	userNames, err := f.friendService.GetFriendList(claims.UserID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, FriendResponse{Error: "Internal server error"})
	}
	return c.JSON(http.StatusOK, FriendListResponse{Users: userNames})
}

func (f *FriendHandler) RemoveFriend(c echo.Context) error {
	claims, ok := c.Get("user").(*jwt_token.Claims)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "user claims not found")
	}
	req := new(FriendRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, FriendResponse{Error: "Invalid request body"})
	}
	// validationをここで行う
	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, FriendResponse{Error: err.Error()})
	}
	if err := f.friendService.RemoveFriend(claims.UserID, req.Username); err != nil {
		switch err {
		case ErrSelfAction:
			return c.JSON(http.StatusBadRequest, FriendResponse{Error: "Cannot remove friend your own profile"})
		case ErrFriendNotFound:
			return c.JSON(http.StatusNotFound, FriendResponse{Error: "Friend not found"})
		default:
			return c.JSON(http.StatusInternalServerError, FriendResponse{Error: "Internal server error"})
		}
	}
	return c.JSON(http.StatusNoContent, FriendResponse{Message: "Successfully remove friend the user"})
}
