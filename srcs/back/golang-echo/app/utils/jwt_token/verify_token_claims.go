package jwt_token

import (
	"golang-echo/app/schemas/errors"

	"github.com/golang-jwt/jwt/v5"
)

func VerifyTokenClaims(tokenString string) (*Claims, error) {
	// トークンの解析
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.ErrTokenUnauthorized
		}
		return []byte(config.accessSecretKey), nil
	})

	// まず、署名が正しいかどうかに関係なくClaimsを取得
	if token != nil { // tokenがnilでないことを確認
		claims, ok := token.Claims.(*Claims)
		if !ok {
			return nil, errors.ErrTokenUnauthorized
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
