package get

import (
	"database/sql"
	"fmt"
	"golang-echo/jwt_token"
	"net/http"

	"github.com/labstack/echo/v4"
)

const (
	query = `
	SELECT 
            u.username,
            u.email,
            ui.lastname,
            ui.firstname,
            ui.birthdate,
            ui.is_gps,
            ui.gender,
            ui.sexuality,
            ui.area,
            ui.self_intro
        FROM 
            users u
        INNER JOIN 
            user_info ui ON u.id = ui.user_id
        WHERE 
            u.id = $1`
)

type userInfoResponse struct {
	UserName  string
	Email     string
	LastName  string
	FirstName string
	BirthDate string
	IsGps     bool
	Gender    string
	Sexuality string
	Area      string
	SelfIntro string
}

func GetUserInfo(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		claims, ok := c.Get("user").(*jwt_token.Claims)
		if !ok {
			return echo.NewHTTPError(http.StatusInternalServerError, "user claims not found")
		}
		userID := claims.UserID

		res := userInfoResponse{}
		err := db.QueryRow(query, userID).Scan(
			&res.UserName,
			&res.Email,
			&res.LastName,
			&res.FirstName,
			&res.BirthDate,
			&res.IsGps,
			&res.Gender,
			&res.Sexuality,
			&res.Area,
			&res.SelfIntro,
		)
		if err != nil {
			if err == sql.ErrNoRows {
				return c.JSON(http.StatusNotFound, map[string]string{"error": "user not found"})
			}
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": fmt.Sprintf("error querying database: %v", err)})
		}
		return c.JSON(http.StatusOK, map[string]userInfoResponse{"userinfos": res})
	}
}
