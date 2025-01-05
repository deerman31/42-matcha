package main

import (
	"database/sql"
	"golang-echo/auth"
	"golang-echo/gps"
	"golang-echo/middle"
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

	protected := e.Group("/api")
	protected.Use(middle.JWTMiddleware())
	users.UserRoutes(protected, db)

	tags.TagRoutes(protected, db)
	gps.GpsRoutes(protected, db)

	// protected.POST("/users/set/user_info", set.InitSetUserInfo(db))

	// protected.PATCH("/users/set/last-name", set.SetLastName(db))
	// protected.PATCH("/users/set/first-name", set.SetFirstName(db))
	// protected.PATCH("/users/set/self-intro", set.SetSelfIntro(db))
	// protected.PATCH("/users/set/area", set.SetArea(db))
	// protected.PATCH("/users/set/gender", set.SetGender(db))
	// protected.PATCH("/users/set/sexuality", set.SetSexuality(db))
	// protected.PATCH("/users/set/is-gps", set.SetIsGps(db))
	// protected.PATCH("/users/set/birthdate", set.SetBirthDate(db))

	// protected.POST("/users/set/image1", set.SetImage(db, set.ImageOne))
	// protected.POST("/users/set/image2", set.SetImage(db, set.ImageTwo))
	// protected.POST("/users/set/image3", set.SetImage(db, set.ImageThree))
	// protected.POST("/users/set/image4", set.SetImage(db, set.ImageFour))
	// protected.POST("/users/set/image5", set.SetImage(db, set.ImageFive))
}
