package gps

import (
	"database/sql"
	"golang-echo/gps/get"
	"golang-echo/gps/set"

	"github.com/labstack/echo/v4"
)

func GpsRoutes(protected *echo.Group, db *sql.DB) {
	gps := protected.Group("/gps")

	setter := gps.Group("/set")
	setter.PATCH("/is-gps", set.SetIsGPS(db))
	setter.PATCH("/location", set.SetLocation(db))
	setter.PATCH("/location-alternative", set.SetLocationAlternative(db))

	getter := gps.Group("/get")
	getter.GET("/is-gps", get.GetIsGPS(db))
	getter.GET("/location", get.GetLocation(db))
	// setter.PATCH("/location-alternative", get.GetLocationAlternative(db))
}
