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

func getAllImagePath(tx *sql.Tx, userID int) ([]string, error) {
	query := `SELECT 
        profile_image_path1, 
        profile_image_path2,
        profile_image_path3,
        profile_image_path4,
        profile_image_path5
    FROM user_image
    WHERE user_id = $1;`

	var paths [5]sql.NullString
	err := tx.QueryRow(query, userID).Scan(&paths[0], &paths[1], &paths[2], &paths[3], &paths[4])
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	var retAllImagePath []string
	for _, path := range paths {
		if path.Valid && path.String != "" {
			retAllImagePath = append(retAllImagePath, path.String)
		} else {
			retAllImagePath = append(retAllImagePath, os.Getenv("DEFAULT_IMAGE"))
		}
	}
	return retAllImagePath, nil
}

func setAllImageURI(allImagePath []string) ([]string, error) {
	var retImages []string
	for _, imagePath := range allImagePath {
		imageData, err := os.ReadFile(imagePath)
		if err != nil {
			return nil, fmt.Errorf("画像の読み込みに失敗しました")
		}
		// MIMEタイプを検出
		mimeType := http.DetectContentType(imageData)

		// Base64エンコード
		base64Data := base64.StdEncoding.EncodeToString(imageData)

		// データURIスキーマを作成
		dataURI := fmt.Sprintf("data:%s;base64,%s", mimeType, base64Data)
		retImages = append(retImages, dataURI)
	}
	return retImages, nil
}

func GetAllImage(db *sql.DB) echo.HandlerFunc {
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
		allImagePath, err := getAllImagePath(tx, userID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Failed to get existing image path",
			})
		}

		retImages, err := setAllImageURI(allImagePath)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": err.Error(),
			})
		}

		// トランザクションのコミット
		if err = tx.Commit(); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Could not commit transaction",
			})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"all_image": retImages,
		})
	}
}
