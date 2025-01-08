package block

import "errors"

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrSelfAction        = errors.New("cannot perform this action on your own profile")
	ErrAlreadyBlock      = errors.New("already block with this user")
	ErrBlockNotFound     = errors.New("block not found")
	ErrTransactionFailed = errors.New("transaction failed")
)
