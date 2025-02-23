package block

import "database/sql"

func checkBlockStatus(tx *sql.Tx, blockerID, blockedID int) (bool, error) {
	const query = `
	SELECT EXISTS (
    SELECT 1 
    FROM user_blocks 
    WHERE blocker_id = $1 
    AND blocked_id = $2
) as is_blocked;`
	var exists bool
	err := tx.QueryRow(query, blockerID, blockedID).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

// blockしているかを調べる
func IsBlock(tx *sql.Tx, myID, otherID int) (bool, error) {
	exists, err := checkBlockStatus(tx, myID, otherID)
	if err != nil {
		return false, err
	}
	return exists, nil
}

// blockされているかを調べる
func IsBlocked(tx *sql.Tx, myID, otherID int) (bool, error) {
	exists, err := checkBlockStatus(tx, otherID, myID)
	if err != nil {
		return false, err
	}
	return exists, nil
}
