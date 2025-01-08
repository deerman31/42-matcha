package like

import (
	"golang-echo/jwt_token"
	"net/http"

	"github.com/labstack/echo/v4"
)

type LikeHandler struct {
	likeService *LikeService
}

func NewLikeHandler(likeService *LikeService) *LikeHandler {
	return &LikeHandler{likeService: likeService}
}

func (h *LikeHandler) DoLike(c echo.Context) error {
	claims, ok := c.Get("user").(*jwt_token.Claims)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "user claims not found")
	}
	req := new(LikeRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, LikeResponse{Error: "Invalid request body"})
	}
	// validationをここで行う
	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, LikeResponse{Error: err.Error()})
	}

	if err := h.likeService.DoLike(claims.UserID, req.Username); err != nil {
		switch err {
		case ErrUserNotFound:
			return c.JSON(http.StatusNotFound, LikeResponse{Error: "User not found"})
		case ErrSelfAction:
			return c.JSON(http.StatusBadRequest, LikeResponse{Error: "Cannot like your own profile"})
		case ErrAlreadyFriends:
			return c.JSON(http.StatusConflict, LikeResponse{Error: "Already friends with this user"})
		case ErrAlreadyLiked:
			return c.JSON(http.StatusConflict, LikeResponse{Error: "You have already liked this user"})
		default:
			return c.JSON(http.StatusInternalServerError, LikeResponse{Error: "Internal server error"})
		}
	}
	return c.JSON(http.StatusCreated, LikeResponse{Message: "Successfully liked the user"})
}

func (h *LikeHandler) UnLike(c echo.Context) error {
	claims, ok := c.Get("user").(*jwt_token.Claims)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "user claims not found")
	}
	req := new(LikeRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, LikeResponse{Error: "Invalid request body"})
	}
	// validationをここで行う
	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, LikeResponse{Error: err.Error()})
	}
	if err := h.likeService.UnLike(claims.UserID, req.Username); err != nil {
		switch err {
		case ErrUserNotFound:
			return c.JSON(http.StatusNotFound, LikeResponse{Error: "User not found"})
		case ErrSelfAction:
			return c.JSON(http.StatusBadRequest, LikeResponse{Error: "Cannot unlike your own profile"})
		case ErrLikeNotFound: // 追加が必要
			return c.JSON(http.StatusNotFound, LikeResponse{Error: "Like relationship not found"})
		default:
			return c.JSON(http.StatusInternalServerError, LikeResponse{Error: "Internal server error"})
		}
	}
	return c.JSON(http.StatusNoContent, LikeResponse{Message: "Successfully unliked the user"})
}

func (h *LikeHandler) GetLikedUsers(c echo.Context) error {
	claims, ok := c.Get("user").(*jwt_token.Claims)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "user claims not found")
	}
	userNames, err := h.likeService.GetLikeList(claims.UserID, LikedUsers)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, LikeResponse{Error: "Internal server error"})
	}
	//return c.JSON(http.StatusOK, map[string][]string{string(LikedUsers): userNames})
	return c.JSON(http.StatusOK, LikeListResponse{Users: userNames})
}

func (h *LikeHandler) GetLikerUsers(c echo.Context) error {
	claims, ok := c.Get("user").(*jwt_token.Claims)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "user claims not found")
	}
	userNames, err := h.likeService.GetLikeList(claims.UserID, LikerUsers)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, LikeResponse{Error: "Internal server error"})
	}
	return c.JSON(http.StatusOK, LikeListResponse{Users: userNames})
}
