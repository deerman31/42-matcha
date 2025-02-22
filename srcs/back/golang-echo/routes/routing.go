package routes

import (
	"database/sql"
	"golang-echo/api/auth"
	"golang-echo/api/browse"
	"golang-echo/api/friend"
	"golang-echo/api/gps"
	myprofile "golang-echo/api/my_profile"
	otherusers "golang-echo/api/other_users"
	"golang-echo/api/profile"
	"golang-echo/api/research"
	"golang-echo/api/tags"
	"golang-echo/api/users"
	"golang-echo/middle"
	"net/http"

	"github.com/labstack/echo/v4"
)

func helloWorldHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func Routing(e *echo.Echo, db *sql.DB) {
	g := e.Group("/api")
	g.GET("", helloWorldHandler)
	auth.AuthRoutes(g, db)

	protected := e.Group("/api")
	protected.Use(middle.JWTMiddleware())
	users.UserRoutes(protected, db)
	browse.BrowseRoutes(protected, db)
	research.ResearchRoutes(protected, db)

	tags.TagRoutes(protected, db)
	gps.GpsRoutes(protected, db)

	profile.ProfileRoutes(protected, db)
	friend.FriendRoutes(protected, db)

	myprofile.MyProfileRoutes(protected, db)

	otherusers.OtherUsersRoutes(protected, db)
}
