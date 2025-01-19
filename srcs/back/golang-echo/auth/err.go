package auth

import "errors"

var (
	ErrUserNotFound         = errors.New("user not found")
	ErrStatusForbidden      = errors.New("Email not verified")
	ErrPasswordUnauthorized = errors.New("Invalid password")
	// StatusUnauthorized
	ErrTransactionFailed = errors.New("transaction failed")

	ErrUserNameEmailConflict = errors.New("Username or Email is already registered")
)
