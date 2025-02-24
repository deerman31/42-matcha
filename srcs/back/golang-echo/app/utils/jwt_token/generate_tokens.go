package jwt_token

import (
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// settingを保持するsturct
type Config struct {
	accessSecretKey       string
	accessTokenLimit      int
	verifyEmailSecretKey  string
	verifyEmailTokenLimit int
}

// シングルトンとしてのconfig
var config *Config

// パッケージの初期化
func init() {
	accessSecretKey := os.Getenv("JWT_SECRET_KEY")
	if accessSecretKey == "" {
		panic("JWT_SECRET_KEY is not set")
	}
	verifyEmailSecretKey := os.Getenv("VERIFY_EMAIL_KEY")
	if verifyEmailSecretKey == "" {
		panic("VERIFY_EMAIL_KEY is not set")
	}
	accessTokenLimitStr := os.Getenv("ACCESS_TOKEN_LIMIT")
	accessTokenLimit, err := strconv.Atoi(accessTokenLimitStr)
	if err != nil {
		accessTokenLimit = 24 // デフォルト値
	}
	verifyEmailTokenLimitStr := os.Getenv("VERIFY_EMAIL_TOKEN_LIMIT")
	verifyEmailTokenLimit, err := strconv.Atoi(verifyEmailTokenLimitStr)
	if err != nil {
		verifyEmailTokenLimit = 1 // デフォルト値
	}
	config = &Config{
		accessSecretKey:       accessSecretKey,
		accessTokenLimit:      accessTokenLimit,
		verifyEmailSecretKey:  verifyEmailSecretKey,
		verifyEmailTokenLimit: verifyEmailTokenLimit,
	}
}

type Claims struct {
	UserID int `json:"user_id"`
	// JWT標準クレーム（有効期限など）を継承
	jwt.RegisteredClaims
}

func GenerateAccessToken(userID int) (string, error) {
	return generateToken(userID, config.accessSecretKey, config.accessTokenLimit)
}

func GenerateVerifyEmailToken(userID int) (string, error) {
	return generateToken(userID, config.verifyEmailSecretKey, config.verifyEmailTokenLimit)
}

func generateToken(userID int, secretKey string, tokenLimit int) (string, error) {
	expiresAt := calculateTokenExpiry(tokenLimit)
	token, err := signNewToken(userID, expiresAt, secretKey)
	if err != nil {
		return "", err
	}
	return token, nil
}

func calculateTokenExpiry(tokenLimit int) time.Time {
	return time.Now().Add(time.Duration(tokenLimit) * time.Hour)
}

func signNewToken(userID int, expiresAt time.Time, secretKey string) (string, error) {
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
