package auth

import (
	"database/sql"
	"errors"
	"fmt"
	pkgErrors "golang-echo/pkg/errors"
	"golang-echo/pkg/jwt_token"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

func checkDuplicateUserCredentials(tx *sql.Tx, username, email string) (int, error) {
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

func createUser(tx *sql.Tx, req *RegisterRequest) (int, error) {
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

func updateUserStatusOn(tx *sql.Tx, userID int) error {
	const updateUserOnlineStatusQuery = `
        UPDATE users 
        SET is_online = TRUE 
        WHERE id = $1
	`
	result, err := tx.Exec(updateUserOnlineStatusQuery, userID)
	if err != nil {
		return pkgErrors.ErrTransactionFailed
	}
	// 更新が成功したか確認
	rows, err := result.RowsAffected()
	if err != nil {
		return pkgErrors.ErrTransactionFailed
	}
	// userが見つからなかった場合
	if rows == 0 {
		return pkgErrors.ErrUserNotFound
	}
	return nil
}

func searchUserDB(tx *sql.Tx, username string) (*User, int, error) {
	// ユーザー名からユーザー情報を取得するクエリ
	const selectUserByUsernameQuery = `
        SELECT id, username, password_hash, is_online, is_registered, is_preparation
        FROM users 
        WHERE username = $1
        LIMIT 1
    `

	user := &User{}
	if err := tx.QueryRow(selectUserByUsernameQuery, username).Scan(
		&user.ID,
		&user.Username,
		&user.PasswordHash,
		&user.isOnline,
		&user.isRegistered,
		&user.isPreparation,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, http.StatusNotFound, errors.New("user not found")
		}
		return nil, http.StatusInternalServerError, fmt.Errorf("database error: %v", err)
	}
	return user, 0, nil
}

func userOnlineStatusOff(tx *sql.Tx, myID int) error {
	const updateUserOfflineStatusQuery = `
        UPDATE users 
        SET is_online = FALSE 
        WHERE id = $1
    `
	result, err := tx.Exec(updateUserOfflineStatusQuery, myID)
	if err != nil {
		return pkgErrors.ErrTransactionFailed
	}
	// 更新が成功したか確認
	rows, err := result.RowsAffected()
	if err != nil {
		return pkgErrors.ErrTransactionFailed
	}
	// userが見つからなかった場合
	if rows == 0 {
		return pkgErrors.ErrUserNotFound
	}
	return nil
}

func verifyTokenClaims(tokenString, secretKey string) (*jwt_token.Claims, error) {
	// トークンの解析
	token, err := jwt.ParseWithClaims(tokenString, &jwt_token.Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	// まず、署名が正しいかどうかに関係なくClaimsを取得
	if token != nil { // tokenがnilでないことを確認
		claims, ok := token.Claims.(*jwt_token.Claims)
		if !ok {
			return nil, fmt.Errorf("invalid token claims")
		}

		// エラーがある場合でも、期限切れエラーのみの場合は claims を返す
		if err != nil {
			if err.Error() == "Token is expired" {
				return claims, nil
			}
		}
		return claims, nil
	}
	// tokenがnilの場合やその他のエラーの場合
	return nil, err
}
