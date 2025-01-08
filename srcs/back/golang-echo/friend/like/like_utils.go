package like

import (
	"database/sql"

	"github.com/lib/pq"
)

func getUserID(tx *sql.Tx, username string) (int, error) {
	var id int
	err := tx.QueryRow(queryGetUserID, username).Scan(&id)
	if err == sql.ErrNoRows {
		return 0, ErrUserNotFound
	}
	return id, err
}

func checkFriendship(tx *sql.Tx, userID1, userID2 int) (bool, error) {
	var exists bool
	minID, maxID := min(userID1, userID2), max(userID1, userID2)
	err := tx.QueryRow(queryCheckFriendExists, minID, maxID).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func checkLikeExists(tx *sql.Tx, likerID, likedID int) (bool, error) {
	var exists bool
	err := tx.QueryRow(queryCheckLikeExists, likerID, likedID).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func createLike(tx *sql.Tx, likerID, likedID int) error {
	_, err := tx.Exec(queryInsertLike, likerID, likedID)
	if pqErr, ok := err.(*pq.Error); ok {
		if pqErr.Code == "23505" { // unique_violation のエラーコード
			return ErrAlreadyLiked
		}
	}
	return err
}

func createFriendship(tx *sql.Tx, userID1, userID2 int) error {
	if _, err := tx.Exec(queryDeleteLike, userID2, userID1); err != nil {
		return err
	}
	minID, maxID := min(userID1, userID2), max(userID1, userID2)
	_, err := tx.Exec(queryInsertFriend, minID, maxID)
	if pqErr, ok := err.(*pq.Error); ok {
		if pqErr.Code == "23505" { // unique_violation のエラーコード
			return ErrAlreadyLiked
		}
	}
	return err
}
