package query

import "database/sql"

const queryGetUserIDByUsername = `SELECT id FROM users WHERE username = $1;`

func GetUserIDByUsername(tx *sql.Tx, username string) (int, error) {
	userID := 0
	if err := tx.QueryRow(queryGetUserIDByUsername, username).Scan(&userID); err != nil {
		return 0, err
	}
	return userID, nil
}
