package auth

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type RegisterRequest struct {
	Username   string `json:"username" validate:"required,username"`
	Email      string `json:"email" validate:"required,email"`
	Password   string `json:"password" validate:"required,password"`
	RePassword string `json:"repassword" validate:"required,password"`
}

const (
	// 1つのクエリで両方をチェック
	checkDuplicateCredentialsQuery = `
        SELECT 
            EXISTS(SELECT 1 FROM users WHERE username = $1) as username_exists,
            EXISTS(SELECT 1 FROM users WHERE email = $2) as email_exists
    `
	// 新規ユーザーを登録するためのクエリ
	insertNewUserQuery = `
        INSERT INTO users (
            username, 
            email, 
            password_hash
        ) VALUES ($1, $2, $3)
		 RETURNING id
    `
)

func Register(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := new(RegisterRequest)
		if err := c.Bind(req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
		}
		// passwordとrepasswordが同じかをCheckする
		if req.Password != req.RePassword {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Password and confirm password do not match"})
		}
		// validationをここで行う
		// Echo のグローバルバリデータを使用
		if err := c.Validate(req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		}
		// トランザクションを開始
		tx, err := db.Begin()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not start transaction"})
		}
		defer tx.Rollback() // エラーが発生した場合はロールバック

		// usernameとemailの重複をcheck
		status, err := checkDuplicateUserCredentials(tx, req.Username, req.Email)
		if err != nil {
			return c.JSON(status, map[string]string{"error": err.Error()})
		}
		// このタイミングでパスワードをハッシュ化する
		hashedBytes, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		req.Password = string(hashedBytes)
		// ユーザーの登録
		//userID, err := createUser(tx, req)
		_, err = createUser(tx, req)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create user"})
		}
		// トランザクションをコミット
		if err = tx.Commit(); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not commit transaction"})
		}

		return c.JSON(http.StatusCreated, map[string]string{"message": "User created successfully. Please check your email to verify your account."})
	}
}

func checkDuplicateUserCredentials(tx *sql.Tx, username, email string) (int, error) {
	var usernameExists, emailExists bool
	err := tx.QueryRow(checkDuplicateCredentialsQuery, username, email).Scan(&usernameExists, &emailExists)
	if err != nil {
		// エラーメッセージをより具体的に
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
	var userID int
    // QueryRowを使用してRETURNINGの結果を取得
    err := tx.QueryRow(insertNewUserQuery, req.Username, req.Email, req.Password).Scan(&userID)
    if err != nil {
        return 0, err
    }
    return userID, nil
}
