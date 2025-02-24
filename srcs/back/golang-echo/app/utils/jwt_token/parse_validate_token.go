package jwt_token

import (
	"golang-echo/app/schemas/errors"

	"github.com/golang-jwt/jwt/v5"
)

func ParseAndValidateAccessToken(tokenString string) (*Claims, error) {
	return parseValidateToken(tokenString, config.accessSecretKey)
}

func ParseAndValidateVerifyEmailToken(tokenString string) (*Claims, error) {
	return parseValidateToken(tokenString, config.verifyEmailSecretKey)
}

func parseValidateToken(tokenStr, secretKey string) (*Claims, error) {
	// トークンの解析
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// 署名方式の検証
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.ErrTokenUnauthorized
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	// クレームの取得と検証
	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.ErrTokenUnauthorized
	}
	return claims, nil
}
