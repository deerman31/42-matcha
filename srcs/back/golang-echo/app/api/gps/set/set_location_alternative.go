package set

import (
	"database/sql"
	"golang-echo/app/utils/jwt_token"
	"net/http"

	"github.com/labstack/echo/v4"
)

const (
	setLocationAlternativeSuccessMessage = "SetLocationAlternative set successfully"

	//location_alternative
	setLocationAlternativeQuery = `
	UPDATE user_location
		SET
			location_alternative = ST_SetSRID(ST_MakePoint($1, $2), 4326)::geography
		WHERE user_id = $3`
)

// type setLocationAlternativeRequest struct {
// 	Latitude  float64
// 	Longitude float64
// }

func SetLocationAlternative(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		claims, ok := c.Get("user").(*jwt_token.Claims)
		if !ok {
			return echo.NewHTTPError(http.StatusInternalServerError, "user claims not found")
		}


		userID := claims.UserID

		var req setLocationRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
		}
		if err := validateCoordinates(req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		}

		// トランザクションを開始
		tx, err := db.Begin()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not start transaction"})
		}
		defer tx.Rollback() // エラーが発生した場合はロールバック

		if _, err := tx.Exec(setLocationAlternativeQuery, req.Longitude, req.Latitude, userID); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		// トランザクションのコミット
		if err = tx.Commit(); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not commit transaction"})
		}

		return c.JSON(http.StatusOK, map[string]string{"message": setLocationAlternativeSuccessMessage})
	}
}
