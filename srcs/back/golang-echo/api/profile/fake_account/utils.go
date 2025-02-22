package fakeaccount

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

func checkFakeAccount(tx *sql.Tx, reporterID, fakeAccountID int) (bool, error) {
	const query = `
    SELECT EXISTS (
        SELECT 1 
        FROM report_fake_accounts 
        WHERE repoter_id = $1 AND fake_account_id = $2
    )
`
	var exists bool
	err := tx.QueryRow(query, reporterID, fakeAccountID).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil

}
