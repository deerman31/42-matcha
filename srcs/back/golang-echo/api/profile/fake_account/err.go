package fakeaccount

import "errors"

var (
	ErrUserNotFound        = errors.New("user not found")
	ErrSelfAction          = errors.New("cannot perform this action on your own profile")
	ErrAlreadyFakeAccount  = errors.New("already report fake account with this user")
	ErrFakeAccountNotFound = errors.New("report fake account not found")
	ErrTransactionFailed   = errors.New("transaction failed")
)
