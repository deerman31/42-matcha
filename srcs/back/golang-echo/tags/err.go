package tags

import "errors"

// エラーメッセージを一箇所に集約
var (
	ErrUserNotFound = errors.New("user not found")


	ErrTag = errors.New("tag error")

	ErrTransactionFailed = errors.New("transaction failed")
)
