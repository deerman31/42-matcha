package friend

import (
	"database/sql"
	"errors"
)

type FriendService struct {
	db *sql.DB
}

// エラーメッセージを一箇所に集約
var (
	ErrUserNotFound = errors.New("user not found")
	ErrSelfAction   = errors.New("cannot perform this action on your own profile")
	// ErrAlreadyFriends    = errors.New("already friends with this user")
	// ErrAlreadyLiked      = errors.New("you have already liked this user")
	ErrFriendNotFound    = errors.New("friend not found")
	ErrTransactionFailed = errors.New("transaction failed")
)

func NewFriendService(db *sql.DB) *FriendService {
	return &FriendService{db: db}
}

// 自分のfriendのリストを返す
func (f *FriendService) GetFriendList(userID int) ([]string, error) {
	rows, err := f.db.Query(queryGetFriendList, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var usernames []string
	for rows.Next() {
		var username string
		if err := rows.Scan(&username); err != nil {
			return nil, err
		}
		usernames = append(usernames, username)
	}
	return usernames, rows.Err()
}

// 自分のfriendを外す
func (f *FriendService) RemoveFriend(myID int, friendName string) error {
	tx, err := f.db.Begin()
	if err != nil {
		return ErrTransactionFailed
	}
	defer tx.Rollback() // エラーが発生した場合はロールバック

	friendID, err := getUserID(tx, friendName)
	if err != nil {
		return err
	}
	if myID == friendID {
		return ErrSelfAction
	}
	/*
		user_friendsテーブルはfriendの重複を防ぐためにuser_id1に小さいid,
		user_id2に大きいidを入れるようにしているため,下記のような処理が必要
	*/
	minID := min(myID, friendID)
	maxID := max(myID, friendID)
	result, err := tx.Exec(queryDeleteFriend, minID, maxID)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrFriendNotFound
	}
	return tx.Commit()
}

func getUserID(tx *sql.Tx, username string) (int, error) {
	var id int
	err := tx.QueryRow(queryGetUserID, username).Scan(&id)
	if err == sql.ErrNoRows {
		return 0, ErrUserNotFound
	}
	return id, err
}
