package auth

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func (a *AuthHandler) register(c echo.Context) error {
	req := new(RegisterRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, RegisterResponse{Error: "Invalid request body"})
	}
	if req.Password != req.RePassword {
		return c.JSON(http.StatusBadRequest, RegisterResponse{Error: "Password and confirm password do not match"})
	}
	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, RegisterResponse{Error: err.Error()})
	}
	err := a.service.register(req)
	if err != nil {
		switch err {
		case ErrUserNotFound:
			return c.JSON(http.StatusNotFound, LoginResponse{Error: "User not found"})
		case ErrUserNameEmailConflict:
			return c.JSON(http.StatusConflict, LoginResponse{Error: "Username or Email is already registered"})
		default:
			return c.JSON(http.StatusInternalServerError, LoginResponse{Error: "Internal server error"})
		}
	}
	return c.JSON(http.StatusCreated, RegisterResponse{Message: "User created successfully. Please check your email to verify your account."})
}

func (a *AuthService) register(req *RegisterRequest) error {
	// トランザクションを開始
	tx, err := a.db.Begin()
	if err != nil {
		return ErrTransactionFailed
	}
	defer tx.Rollback() // エラーが発生した場合はロールバック
	// usernameとemailの重複をcheck
	status, err := checkDuplicateUserCredentials(tx, req.Username, req.Email)
	if err != nil {
		if status == http.StatusConflict {
			return ErrUserNameEmailConflict
		}
		return ErrTransactionFailed
	}
	// このタイミングでパスワードをハッシュ化する
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return ErrTransactionFailed
	}
	req.Password = string(hashedBytes)
	// ユーザーの登録
	_, err = createUser(tx, req)
	if err != nil {
		return ErrTransactionFailed
	}
	return tx.Commit()
}

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
