package myprofile

import (
	"database/sql"

	"github.com/labstack/echo/v4"
)

func MyProfileRoutes(protected *echo.Group, db *sql.DB) {
	myProfile := protected.Group("/my-profile")

	handler := NewMyProfileHandler(NewMyProfileService(db))

	myProfile.GET("", handler.GetMyProfile)
}
