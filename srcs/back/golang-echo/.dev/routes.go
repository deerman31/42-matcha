package dev

import (
	"database/sql"

	"github.com/labstack/echo/v4"
)

func DevRoutes(g *echo.Group, db *sql.DB) {
	dev := g.Group("/dev")

	handler := NewFiveThousandRegisterHandler(NewFiveThousandRegisterService(db))

	dev.POST("", handler.AllRegister)
}
