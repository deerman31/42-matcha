package profile

import (
	"database/sql"
	fakeaccount "golang-echo/profile/fake_account"
	"golang-echo/profile/view"

	"github.com/labstack/echo/v4"
)

func ProfileRoutes(protected *echo.Group, db *sql.DB) {
	profile := protected.Group("/profile")

	fakeaccount.FakeAccountRoutes(profile, db)

	viewRoute := profile.Group("/view")
	//view.POST("/set", view.SetProfileViewed(db))
	viewRoute.POST("/set", view.SetProfileViewed(db))
	viewRoute.GET("/get-viewed", view.GetViewedUsers(db))
	viewRoute.GET("/get-viewer", view.GetViewerUsers(db))

}
