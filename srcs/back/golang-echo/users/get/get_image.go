package get

import (
	"database/sql"
	"encoding/base64"
	"fmt"
	"golang-echo/jwt_token"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func GetImage(db *sql.DB, imageNum int) echo.HandlerFunc {
	return func(c echo.Context) error {
		claims, ok := c.Get("user").(*jwt_token.Claims)
		if !ok {
			return echo.NewHTTPError(http.StatusInternalServerError, "user claims not found")
		}
		userID := claims.UserID
		// トランザクションを開始
		tx, err := db.Begin()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not start transaction"})
		}
		defer tx.Rollback() // エラーが発生した場合はロールバック

		var imagePath sql.NullString

		//query:="SELECT profile_image_path1 FROM user_info WHERE user_id = ?"
		query := fmt.Sprintf("SELECT profile_image_path%d FROM user_image WHERE user_id = $1", imageNum)
		err = tx.QueryRow(query, userID).Scan(&imagePath)
		if err != nil && err != sql.ErrNoRows {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Failed to get existing image path",
			})
		}
		retImagePath := ""
		if imagePath.Valid && imagePath.String != "" {
			retImagePath = imagePath.String
		} else {
			retImagePath = os.Getenv("DEFAULT_IMAGE")
		}
		// 画像データを送信する途中
		/*
			取得した画像pathから画像を取得
			取得した画像データをなにかしらのデータとしてフロントに送信
		*/
		imageData, err := os.ReadFile(retImagePath)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "画像の読み込みに失敗しました",
			})
		}

		// MIMEタイプを検出
		mimeType := http.DetectContentType(imageData)

		// Base64エンコード
		base64Data := base64.StdEncoding.EncodeToString(imageData)

		// データURIスキーマを作成
		dataURI := fmt.Sprintf("data:%s;base64,%s", mimeType, base64Data)
		// トランザクションのコミット
		if err = tx.Commit(); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Could not commit transaction",
			})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"image":   dataURI,
			"message": "success",
		})

	}
}
