package tags

import (
	"database/sql"

	"github.com/labstack/echo/v4"
)

func TagRoutes(protected *echo.Group, db *sql.DB) {
	route := protected.Group("/tag")

	handler := NewTagHandler(NewTagService(db))

	route.GET("/get-user", handler.GetUserTag)

	route.POST("/set", handler.SetTag)
	route.POST("/delete", handler.DeleteTag)
}
