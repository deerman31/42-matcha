package routes

import (
	"database/sql"
	"golang-echo/app/handlers"
	"golang-echo/app/services"

	"github.com/labstack/echo/v4"
)

func authRoutes(g *echo.Group, db *sql.DB) {
	route := g.Group("/auth")

	handler := handlers.NewAuthHandler(services.NewAuthService(db))
	route.POST("/register", handler.RegisterHandler)
	route.POST("/login", handler.LoginHandler)
	route.POST("/logout", handler.LogoutHandler)
}
