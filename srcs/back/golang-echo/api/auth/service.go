package auth

import (
	"database/sql"
	"golang-echo/pkg/errors"
	"golang-echo/pkg/jwt_token"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	db *sql.DB
}

func newAuthService(db *sql.DB) *AuthService {
	return &AuthService{db: db}
}

func (a *AuthService) registerService(req *RegisterRequest) error {
	// トランザクションを開始
	tx, err := a.db.Begin()
	if err != nil {
		return errors.ErrTransactionFailed
	}
	defer tx.Rollback() // エラーが発生した場合はロールバック
	// usernameとemailの重複をcheck
	status, err := checkDuplicateUserCredentials(tx, req.Username, req.Email)
	if err != nil {
		if status == http.StatusConflict {
			return errors.ErrUserNameEmailConflict
		}
		return errors.ErrTransactionFailed
	}
	// このタイミングでパスワードをハッシュ化する
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.ErrTransactionFailed
	}
	req.Password = string(hashedBytes)
	// ユーザーの登録
	_, err = createUser(tx, req)
	if err != nil {
		return errors.ErrTransactionFailed
	}
	return tx.Commit()
}

func (a *AuthService) loginService(req *LoginRequest) (*User, string, error) {
	// トランザクションを開始
	tx, err := a.db.Begin()
	if err != nil {
		return nil, "", errors.ErrTransactionFailed
	}
	defer tx.Rollback() // エラーが発生した場合はロールバック

	user, status, err := searchUserDB(tx, req.Username)
	if err != nil {
		if status == http.StatusNotFound {
			return nil, "", errors.ErrUserNotFound
		} else {
			return nil, "", errors.ErrTransactionFailed
		}
	}
	// userがメールで認証済みかどうか確認
	if !user.isRegistered {
		return nil, "", errors.ErrStatusForbidden
	}

	if bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)) != nil {
		return nil, "", errors.ErrPasswordUnauthorized
	}

	accessToken, err := jwt_token.GenerateAccessToken(user.ID)
	if err != nil {
		return nil, "", errors.ErrTransactionFailed
	}

	if err := updateUserStatusOn(tx, user.ID); err != nil {
		return nil, "", err
	}

	// トランザクションのコミット
	if err = tx.Commit(); err != nil {
		return nil, "", errors.ErrTransactionFailed
	}
	return user, accessToken, nil
}

func (a *AuthService) logoutService(myID int) error {

	// トランザクションを開始
	tx, err := a.db.Begin()
	if err != nil {
		return errors.ErrTransactionFailed
	}
	defer tx.Rollback() // エラーが発生した場合はロールバック

	if err := userOnlineStatusOff(tx, myID); err != nil {
		return err
	}
	return tx.Commit()
}
