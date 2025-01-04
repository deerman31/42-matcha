package tags

import (
	"database/sql"

	"github.com/labstack/echo/v4"
)

func TagRoutes(protected *echo.Group, db *sql.DB) {
	route := protected.Group("/tag")

	route.GET("/get-user", GetUserTag(db))

	route.POST("/set", SetTag(db))
	route.POST("/delete", DeleteTag(db))
}
