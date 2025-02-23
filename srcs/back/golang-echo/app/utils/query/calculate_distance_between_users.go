package query

import (
	"database/sql"
	"fmt"
)

// 2人のユーザー間の距離を計算するクエリ（単位: km）
const queryGetDistanceBetweenUsers = `
	SELECT 
		-- is_gpsがtrueの場合はlocation、falseの場合はlocation_alternativeを使用
		ST_Distance(
			CASE WHEN ul1.is_gps THEN ul1.location ELSE ul1.location_alternative END,
			CASE WHEN ul2.is_gps THEN ul2.location ELSE ul2.location_alternative END
		) / 1000.0 AS distance_km
	FROM 
		user_location ul1,
		user_location ul2
	WHERE 
		ul1.user_id = $1  -- 1人目のuser_id
		AND ul2.user_id = $2;  -- 2人目のuser_id
	`

// func GetUserIDByUsername(tx *sql.Tx, username string) (int, error) {
func CalculateDistanceBetweenUsers(tx *sql.Tx, myID, otherID int) (int, error) {
	if myID == otherID {
		return 0.0, nil
	}

	var distance float64
	err := tx.QueryRow(queryGetDistanceBetweenUsers, myID, otherID).Scan(&distance)
	if err == sql.ErrNoRows {
		return 0.0, fmt.Errorf("one or both users (ID: %d, %d) not found in location table", myID, otherID)
	}
	if err != nil {
		return 1.0, fmt.Errorf("error calculating distance: %w", err)
	}
	if distance < 1.0 {
		return 1, nil
	} else {
		return int(distance), nil
	}
}
