package routes

import (
	"database/sql"
	"golang-echo/app/api/browse"
	"golang-echo/app/api/friend"
	"golang-echo/app/api/gps"
	myprofile "golang-echo/app/api/my_profile"
	otherusers "golang-echo/app/api/other_users"
	"golang-echo/app/api/profile"
	"golang-echo/app/api/research"
	"golang-echo/app/api/tags"
	"golang-echo/app/api/users"
	"golang-echo/app/middle"
	"net/http"

	"github.com/labstack/echo/v4"
)

func helloWorldHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func Routing(e *echo.Echo, db *sql.DB) {
	g := e.Group("/api")
	g.GET("", helloWorldHandler)
	authRoutes(g, db)

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
