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
	// 仮登録
	route.POST("/register", handler.RegisterHandler)
	// Login
	route.POST("/login", handler.LoginHandler)
	// Logout
	route.POST("/logout", handler.LogoutHandler)

	// メールを使った本登録のエンドポイント
	route.GET("/verify-email/:token", handler.VerifyEmailHandler)

	route.POST("/reset-password-email", handler.ResetPasswordEmailHandler)

	route.POST("/reset-password/:token", handler.ResetPasswordHandler)
}
