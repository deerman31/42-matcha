package dev

import (
	"database/sql"

	"github.com/labstack/echo/v4"
)

func DevRoutes(protected *echo.Group, db *sql.DB) {
	dev := protected.Group("/dev")

	handler := NewFiveThousandRegisterHandler(NewFiveThousandRegisterService(db))

	dev.POST("/", handler.AllRegister)
}
