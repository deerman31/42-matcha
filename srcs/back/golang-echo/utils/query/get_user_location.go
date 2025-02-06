package query

import (
	"database/sql"
	"fmt"
)

// PostGISのgeography型からポイントを取得するクエリ
const getUserLocationQuery = `
SELECT 
    CASE 
        WHEN is_gps = true THEN ST_Y(location::geometry)
        ELSE ST_Y(location_alternative::geometry)
    END as latitude,
    CASE 
        WHEN is_gps = true THEN ST_X(location::geometry)
        ELSE ST_X(location_alternative::geometry)
    END as longitude,
    is_gps
FROM 
    user_location 
WHERE 
    user_id = $1
    `

// Location情報を格納する構造体
type UserLocation struct {
	Latitude  float64
	Longitude float64
	IsGPS     bool
}

func GetUserLocation(tx *sql.Tx, userID int) (UserLocation, error) {
	var loc UserLocation

	err := tx.QueryRow(getUserLocationQuery, userID).Scan(
		&loc.Latitude,
		&loc.Longitude,
		&loc.IsGPS,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return UserLocation{}, fmt.Errorf("location not found for user: %w", err)
		}
		return UserLocation{}, fmt.Errorf("error querying user location: %w", err)
	}
	return loc, nil
}
