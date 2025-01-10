package browse

import "errors"

// エラーメッセージを一箇所に集約
var (
	ErrUserNotFound = errors.New("user not found")
	//ErrSelfAction   = errors.New("cannot perform this action on your own profile")
	// ErrAlreadyFriends    = errors.New("already friends with this user")
	// ErrAlreadyLiked      = errors.New("you have already liked this user")
	ErrFriendNotFound    = errors.New("friend not found")

	ErrTransactionFailed = errors.New("transaction failed")
)
