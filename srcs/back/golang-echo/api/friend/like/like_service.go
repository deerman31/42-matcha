package like

import (
	"database/sql"
	"errors"
)

type LikeService struct {
	db *sql.DB
}

// コンストラクター
func NewLikeService(db *sql.DB) *LikeService {
	return &LikeService{db: db}
}

// エラーメッセージを一箇所に集約
var (
	ErrUserNotFound      = errors.New("user not found")
	ErrSelfAction        = errors.New("cannot perform this action on your own profile")
	ErrAlreadyFriends    = errors.New("already friends with this user")
	ErrAlreadyLiked      = errors.New("you have already liked this user")
	ErrLikeNotFound      = errors.New("like not found")
	ErrTransactionFailed = errors.New("transaction failed")
)

func (s *LikeService) DoLike(userID int, userName string) error {
	// トランザクションを開始
	tx, err := s.db.Begin()
	if err != nil {
		return ErrTransactionFailed
	}
	defer tx.Rollback() // エラーが発生した場合はロールバック

	likedID, err := getUserID(tx, userName)
	if err != nil {
		return err
	}
	if userID == likedID {
		return ErrSelfAction
	}

	isFriend, err := checkFriendship(tx, userID, likedID)
	if err != nil {
		return err
	}

	if isFriend {
		return ErrAlreadyFriends
	}

	isLiked, err := checkLikeExists(tx, likedID, userID)
	if err != nil {
		return err
	}

	if isLiked {
		if err := createFriendship(tx, userID, likedID); err != nil {
			return err
		}
	} else {
		if err := createLike(tx, userID, likedID); err != nil {
			return err
		}
	}
	return tx.Commit()
}

func (s *LikeService) UnLike(userID int, userName string) error {
	// トランザクションを開始
	tx, err := s.db.Begin()
	if err != nil {
		return ErrTransactionFailed
	}
	defer tx.Rollback() // エラーが発生した場合はロールバック

	likedID, err := getUserID(tx, userName)
	if err != nil {
		return err
	}
	if userID == likedID {
		return ErrSelfAction
	}
	result, err := tx.Exec(queryDeleteLike, userID, likedID)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrLikeNotFound
	}
	return tx.Commit()
}

func (s *LikeService) GetLikeList(userID int, likeType LikeType) ([]string, error) {
	query := queryGetLikedUsers
	if likeType == LikerUsers {
		query = queryGetLikerUsers
	}
	rows, err := s.db.Query(query, userID)
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
