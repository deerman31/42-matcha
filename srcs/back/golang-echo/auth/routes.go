package auth

import (
	"database/sql"

	"github.com/labstack/echo/v4"
)

func AuthRoutes(g *echo.Group, db *sql.DB) {
	handler := NewAuthHandler(NewAuthService(db))
	g.POST("/register", handler.register)
	g.POST("/login", handler.Login)
	g.POST("/logout", handler.Logout)
}
