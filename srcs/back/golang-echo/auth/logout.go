package auth

import (
	"fmt"
	"golang-echo/jwt_token"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func (a *AuthHandler) Logout(c echo.Context) error {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	// Authorizationヘッダーを取得
	tokenString, err := jwt_token.GetAuthToken(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, LogoutResponse{Error: err.Error()})
	}
	claims, err := verifyTokenClaims(tokenString, secretKey)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, LogoutResponse{Error: err.Error()})
	}
	userID := claims.UserID
	err = a.service.Logout(userID)
	if err != nil {
		switch err {
		case ErrUserNotFound:
			return c.JSON(http.StatusNotFound, LoginResponse{Error: "User not found"})
		default:
			return c.JSON(http.StatusInternalServerError, LoginResponse{Error: "Internal server error"})
		}
	}
	return c.JSON(http.StatusOK, LogoutResponse{Message: "User logout successfully."})
}
func (a *AuthService) Logout(myID int) error {
	// トランザクションを開始
	tx, err := a.db.Begin()
	if err != nil {
		return ErrTransactionFailed
	}
	defer tx.Rollback() // エラーが発生した場合はロールバック

	// userのis_onlineをfalseにする
	result, err := tx.Exec(updateUserOfflineStatusQuery, myID)
	if err != nil {
		return ErrTransactionFailed
	}
	// 更新が成功したか確認
	rows, err := result.RowsAffected()
	if err != nil {
		return ErrTransactionFailed
	}
	// userが見つからなかった場合
	if rows == 0 {
		return ErrUserNotFound
	}
	return tx.Commit()
}

// func Logout(db *sql.DB) echo.HandlerFunc {
// 	secretKey := os.Getenv("JWT_SECRET_KEY")
// 	return func(c echo.Context) error {
// 		// Authorizationヘッダーを取得
// 		tokenString, err := jwt_token.GetAuthToken(c)
// 		if err != nil {
// 			return c.JSON(http.StatusUnauthorized, map[string]string{
// 				"error": err.Error(),
// 			})
// 		}
// 		claims, err := verifyTokenClaims(tokenString, secretKey)
// 		if err != nil {
// 			return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
// 		}
// 		userID := claims.UserID
// 		// トランザクションを開始
// 		tx, err := db.Begin()
// 		if err != nil {
// 			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not start transaction"})
// 		}
// 		defer tx.Rollback() // エラーが発生した場合はロールバック

// 		// userのis_onlineをfalseにする
// 		result, err := tx.Exec(updateUserOfflineStatusQuery, userID)
// 		if err != nil {
// 			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
// 		}
// 		// 更新が成功したか確認
// 		rows, err := result.RowsAffected()
// 		if err != nil {
// 			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
// 		}
// 		// userが見つからなかった場合
// 		if rows == 0 {
// 			return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
// 		}
// 		// トランザクションのコミット
// 		if err = tx.Commit(); err != nil {
// 			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not commit transaction"})
// 		}
// 		return c.JSON(http.StatusOK, map[string]string{"message": "User logout successfully."})
// 	}
// }

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
			return nil, fmt.Errorf("Invalid token claims")
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
