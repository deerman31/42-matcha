package otherusers

import (
	"database/sql"

	"github.com/labstack/echo/v4"
)

func OtherUsersRoutes(protected *echo.Group, db *sql.DB) {
	other := protected.Group("/other-users")

	handler := NewOtherUsersHandler(NewOtherUsersService(db))

	get:= other.Group("/get")
	get.POST("/image", handler.OtherGetImage)
	get.GET("/profile/:name", handler.OtherGetProfile)
}
