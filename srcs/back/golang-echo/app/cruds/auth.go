package cruds

import (
	"database/sql"
	"errors"
	"fmt"
	"golang-echo/app/schemas"
	myErrors "golang-echo/app/schemas/errors"
	"net/http"
)

func CheckDuplicateUserCredentials(tx *sql.Tx, username, email string) (int, error) {
	// 1つのクエリで両方をチェック
	const checkDuplicateCredentialsQuery = `
        SELECT 
            EXISTS(SELECT 1 FROM users WHERE username = $1) as username_exists,
            EXISTS(SELECT 1 FROM users WHERE email = $2) as email_exists
    `
	var usernameExists, emailExists bool
	err := tx.QueryRow(checkDuplicateCredentialsQuery, username, email).Scan(&usernameExists, &emailExists)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("failed to check credentials: %w", err)
	}

	// 存在チェックの順序を明確に
	switch {
	case usernameExists:
		return http.StatusConflict, fmt.Errorf("username %s is already taken", username)
	case emailExists:
		return http.StatusConflict, fmt.Errorf("email %s is already registered", email)
	default:
		return http.StatusOK, nil
	}
}

func CreateUser(tx *sql.Tx, req *schemas.RegisterRequest) (int, error) {

	// 新規ユーザーを登録するためのクエリ
	const insertNewUserQuery = `
        INSERT INTO users (
            username, 
            email, 
            password_hash
        ) VALUES ($1, $2, $3)
		 RETURNING id
    `
	var userID int
	// QueryRowを使用してRETURNINGの結果を取得
	err := tx.QueryRow(insertNewUserQuery, req.Username, req.Email, req.Password).Scan(&userID)
	if err != nil {
		return 0, err
	}
	return userID, nil
}

func UpdateUserStatusOn(tx *sql.Tx, userID int) error {
	const updateUserOnlineStatusQuery = `
        UPDATE users 
        SET is_online = TRUE 
        WHERE id = $1
	`
	result, err := tx.Exec(updateUserOnlineStatusQuery, userID)
	if err != nil {
		return myErrors.ErrTransactionFailed
	}
	// 更新が成功したか確認
	rows, err := result.RowsAffected()
	if err != nil {
		return myErrors.ErrTransactionFailed
	}
	// userが見つからなかった場合
	if rows == 0 {
		return myErrors.ErrUserNotFound
	}
	return nil
}

func SearchUserDB(tx *sql.Tx, username string) (*schemas.User, int, error) {

	// ユーザー名からユーザー情報を取得するクエリ
	const selectUserByUsernameQuery = `
        SELECT id, username, password_hash, is_online, is_registered, is_preparation
        FROM users 
        WHERE username = $1
        LIMIT 1
    `

	user := &schemas.User{}
	if err := tx.QueryRow(selectUserByUsernameQuery, username).Scan(
		&user.ID,
		&user.Username,
		&user.PasswordHash,
		&user.IsOnline,
		&user.IsRegistered,
		&user.IsPreparation,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, http.StatusNotFound, errors.New("user not found")
		}
		return nil, http.StatusInternalServerError, fmt.Errorf("database error: %v", err)
	}
	return user, 0, nil
}

func UserOnlineStatusOff(tx *sql.Tx, myID int) error {
	const updateUserOfflineStatusQuery = `
        UPDATE users 
        SET is_online = FALSE 
        WHERE id = $1
    `
	result, err := tx.Exec(updateUserOfflineStatusQuery, myID)
	if err != nil {
		return myErrors.ErrTransactionFailed
	}
	// 更新が成功したか確認
	rows, err := result.RowsAffected()
	if err != nil {
		return myErrors.ErrTransactionFailed
	}
	// userが見つからなかった場合
	if rows == 0 {
		return myErrors.ErrUserNotFound
	}
	return nil
}

func UpdateUserStatusRegister(tx *sql.Tx, myID int) error {
	const Query = `
        UPDATE users 
        SET is_registered = TRUE 
        WHERE id = $1
    `
	result, err := tx.Exec(Query, myID)
	if err != nil {
		return myErrors.ErrTransactionFailed
	}
	// 更新が成功したか確認
	rows, err := result.RowsAffected()
	if err != nil {
		return myErrors.ErrTransactionFailed
	}
	// userが見つからなかった場合
	if rows == 0 {
		return myErrors.ErrUserNotFound
	}
	return nil
}
