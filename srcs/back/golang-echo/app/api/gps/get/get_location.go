package get

import (
	"database/sql"
	"golang-echo/app/utils/jwt_token"
	"net/http"

	"github.com/labstack/echo/v4"
)

type getLocationResponse struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func GetLocation(db *sql.DB) echo.HandlerFunc {
	const query = `
SELECT 
    CASE 
        WHEN is_gps = TRUE THEN ST_X(location::geometry)
        ELSE ST_X(location_alternative::geometry)
    END as longitude,
    CASE 
        WHEN is_gps = TRUE THEN ST_Y(location::geometry)
        ELSE ST_Y(location_alternative::geometry)
    END as latitude
FROM user_location 
WHERE user_id = $1;
`

	return func(c echo.Context) error {
		claims, ok := c.Get("user").(*jwt_token.Claims)
		if !ok {
			return echo.NewHTTPError(http.StatusInternalServerError, "user claims not found")
		}
		userID := claims.UserID

		var result getLocationResponse

		if err := db.QueryRow(query, userID).Scan(&result.Longitude, &result.Latitude); err != nil {
			if err == sql.ErrNoRows {
				return c.JSON(http.StatusNotFound, map[string]string{"error": "Location not found"})
			}
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		return c.JSON(http.StatusOK, map[string]getLocationResponse{"location": result})
	}
}
