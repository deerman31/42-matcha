package auth

import (
	"database/sql"

	"github.com/labstack/echo/v4"
)

func AuthRoutes(g *echo.Group, db *sql.DB) {
	route := g.Group("/auth")

	handler := newAuthHandler(newAuthService(db))
	route.POST("/register", handler.registerHandler)
	route.POST("/login", handler.loginHandler)
	route.POST("/logout", handler.logoutHandler)

	// メール認証のエンドポイント
	// パスワード変更のエンドポイント
}
