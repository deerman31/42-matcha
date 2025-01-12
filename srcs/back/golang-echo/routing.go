package main

import (
	"database/sql"
	"golang-echo/auth"
	"golang-echo/browse"
	"golang-echo/friend"
	"golang-echo/gps"
	"golang-echo/middle"
	"golang-echo/profile"
	"golang-echo/research"
	"golang-echo/tags"
	"golang-echo/users"
	"net/http"

	"github.com/labstack/echo/v4"
)

func helloWorldHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func routing(e *echo.Echo, db *sql.DB) {
	g := e.Group("/api")
	g.GET("", helloWorldHandler)
	g.POST("/register", auth.Register(db))
	g.POST("/login", auth.Login(db))
	g.POST("/logout", auth.Logout(db))
	//test用
	//dev.DevRoutes(g, db)

	protected := e.Group("/api")
	protected.Use(middle.JWTMiddleware())
	users.UserRoutes(protected, db)
	browse.BrowseRoutes(protected, db)
	research.ResearchRoutes(protected, db)

	tags.TagRoutes(protected, db)
	gps.GpsRoutes(protected, db)

	profile.ProfileRoutes(protected, db)
	friend.FriendRoutes(protected, db)

}
