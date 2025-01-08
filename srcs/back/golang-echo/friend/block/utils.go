package block

import "database/sql"

func getUserID(tx *sql.Tx, username string) (int, error) {
	const queryGetUserID = `SELECT id FROM users WHERE username = $1;`
	var id int
	err := tx.QueryRow(queryGetUserID, username).Scan(&id)
	if err == sql.ErrNoRows {
		return 0, ErrUserNotFound
	}
	return id, err
}
