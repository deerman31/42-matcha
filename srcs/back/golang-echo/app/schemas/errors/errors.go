package errors

import "errors"

var (
	ErrUserNotFound         = errors.New("user not found")
	ErrStatusForbidden      = errors.New("email not verified")
	ErrPasswordUnauthorized = errors.New("invalid password")

	ErrTransactionFailed = errors.New("transaction failed")

	ErrUserNameEmailConflict = errors.New("username or email is already registered")
)
