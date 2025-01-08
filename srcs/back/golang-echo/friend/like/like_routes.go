package like

import (
	"database/sql"

	"github.com/labstack/echo/v4"
)

func LikeRoutes(protected *echo.Group, db *sql.DB) {
	like := protected.Group("/like")

	likeHandler := NewLikeHandler(NewLikeService(db))
	like.GET("/get-liked", likeHandler.GetLikedUsers)
	like.GET("/get-liker", likeHandler.GetLikerUsers)
	like.POST("/do-like", likeHandler.DoLike)
	like.DELETE("/un-like", likeHandler.UnLike)
}
