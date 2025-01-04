package jwt_token

import (
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID int `json:"user_id"`
	// JWT標準クレーム（有効期限など）を継承
	jwt.RegisteredClaims
}

// func GenerateAccessToken(userID int, secretKey string) (string, error) {
func GenerateAccessToken(userID int) (string, error) {
	secretKey := os.Getenv("JWT_SECRET_KEY")

	// AccessTokenを生成
	accessExpiresAt := calculateAccessTokenExpiry()

	accessToken, err := signNewToken(userID, accessExpiresAt, secretKey)
	if err != nil {
		return "", err
	}
	return accessToken, nil
}

func signNewToken(userID int, expiresAt time.Time, secretKey string) (string, error) {
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	//token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func calculateAccessTokenExpiry() time.Time {
	accessTokenLimitStr := os.Getenv("ACCESS_TOKEN_LIMIT")
	accessTokenLimit, err := strconv.Atoi(accessTokenLimitStr)
	if err != nil {
		accessTokenLimit = 24
	}
	return time.Now().Add(time.Duration(accessTokenLimit) * time.Hour)
}
