package set

import (
	"database/sql"
	"fmt"
	"golang-echo/app/utils/jwt_token"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

const (
	maxFileSize         = 5 << 20 // 5MB
	allowedMIMETypeJPEG = "image/jpeg"
	allowedMIMETypePNG  = "image/png"
	allowedMIMETypeGIF  = "image/gif"
)

// 画像番号の定数定義
const (
	ImageOne = iota + 1
	ImageTwo
	ImageThree
	ImageFour
	ImageFive
)

// SQLクエリをマップで管理
var imageQueries = map[int]string{
	ImageOne:   `UPDATE user_image SET profile_image_path1 = $1 WHERE user_id = $2 RETURNING id`,
	ImageTwo:   `UPDATE user_image SET profile_image_path2 = $1 WHERE user_id = $2 RETURNING id`,
	ImageThree: `UPDATE user_image SET profile_image_path3 = $1 WHERE user_id = $2 RETURNING id`,
	ImageFour:  `UPDATE user_image SET profile_image_path4 = $1 WHERE user_id = $2 RETURNING id`,
	ImageFive:  `UPDATE user_image SET profile_image_path5 = $1 WHERE user_id = $2 RETURNING id`,
}

func SetImage(db *sql.DB, imageNum int) echo.HandlerFunc {
	uploadDir := os.Getenv(fmt.Sprintf("IMAGE_UPLOAD_DIR%d", imageNum))
	return func(c echo.Context) error {
		claims, ok := c.Get("user").(*jwt_token.Claims)
		if !ok {
			return echo.NewHTTPError(http.StatusInternalServerError, "user claims not found")
		}
		userID := claims.UserID

		// マルチパートフォームファイルを取得
		/* ここでuploadする画像データを取得する */
		file, err := c.FormFile("image")
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid file upload"})
		}
		// ファイルサイズチェック
		if err := checkFileSize(file); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		}
		// ファイル形式チェック
		if !isValidImageType(file.Header.Get("Content-Type")) {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Invalid file type. Only JPEG, PNG and GIF are allowed",
			})
		}
		// ファイル名生成（UUID + 元の拡張子）
		newFileName := generateNewFileName(file.Filename)

		// uploadするディレクトリ名とファイル名を結合する
		filePath := filepath.Join(uploadDir, newFileName)

		// ファイル保存とDB更新
		if err := saveFile(file, filePath); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Failed to save file",
			})
		}

		// トランザクションを開始
		tx, err := db.Begin()
		if err != nil {
			os.Remove(filePath)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not start transaction"})
		}
		defer tx.Rollback() // エラーが発生した場合はロールバック

		// DB更新
		if status, err := executeImageUpdate(tx, filePath, userID, imageNum); err != nil {
			// エラー時は新しくアップロードしたファイルを削除
			os.Remove(filePath)
			return c.JSON(status, map[string]string{"error": err.Error()})
		}
		// トランザクションのコミット
		if err = tx.Commit(); err != nil {
			os.Remove(filePath)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not commit transaction"})
		}
		return c.JSON(http.StatusOK, map[string]string{
			"message": fmt.Sprintf("Set image%d successfully.", imageNum),
		})
	}
}

// DB更新の共通ロジック
func executeImageUpdate(tx *sql.Tx, imagePath string, userID int, imageNum int) (int, error) {
	query, exists := imageQueries[imageNum]
	if !exists {
		return http.StatusBadRequest, fmt.Errorf("invalid image number")
	}

	result, err := tx.Exec(query, imagePath, userID)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("failed to update image path: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return http.StatusInternalServerError, err
	}
	if rows == 0 {
		return http.StatusNotFound, fmt.Errorf("user not found")
	}

	return http.StatusOK, nil
}

func generateNewFileName(fileName string) string {
	ext := filepath.Ext(fileName)
	newFileName := fmt.Sprintf("%s%s", uuid.NewString(), ext)
	return newFileName
}

// ファイルサイズをチェックする関数
func checkFileSize(file *multipart.FileHeader) error {
	if file.Size > maxFileSize {
		return fmt.Errorf("File size exceeds maximum limit of %d MB", maxFileSize/(1<<20))
	}
	return nil
}

func isValidImageType(contentType string) bool {
	validTypes := map[string]bool{
		allowedMIMETypeJPEG: true,
		allowedMIMETypePNG:  true,
		allowedMIMETypeGIF:  true,
	}
	return validTypes[contentType]
}

func saveFile(file *multipart.FileHeader, destPath string) error {
	// パスのバリデーションを追加
	// IsAbs()は引数が絶対Pathかどうかを調べる関数
	if !filepath.IsAbs(destPath) {
		return fmt.Errorf("destination path must be absolute")
	}

	/* マルチパートファイルをオープンしてReaderを取得
	エラー発生時は即座にエラーを返す
	deferを使用して関数終了時に確実にファイルをクローズ */
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	/* 指定されたパスに新しいファイルを作成
	エラー発生時は即座にエラーを返す
	deferを使用して関数終了時に確実にファイルをクローズ */
	dst, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer dst.Close()
	/* io.Copyを使用してソースファイル（src）から保存先ファイル（dst）にデータをコピー
	コピー処理中のエラーがあれば、それを返す */
	_, err = io.Copy(dst, src)
	return err
}
