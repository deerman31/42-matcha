package friend

import (
	"database/sql"
	"golang-echo/api/friend/like"

	"github.com/labstack/echo/v4"
)

func FriendRoutes(protected *echo.Group, db *sql.DB) {
	friend := protected.Group("/friend")

	like.LikeRoutes(friend, db)

	friendHandler := NewFriendHandler(NewFriendService(db))
	friend.GET("/get", friendHandler.GetFriendList)
	friend.DELETE("/delete", friendHandler.RemoveFriend)
}
