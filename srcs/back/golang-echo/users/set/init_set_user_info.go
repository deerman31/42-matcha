package set

import (
	"database/sql"
	"golang-echo/jwt_token"
	"net/http"
	"os"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

const (
	setUserInfoQuery = `
		UPDATE user_info
		SET 
			lastname = $1,
			firstname = $2,
			birthdate = $3,
			gender = $4,
			sexuality = $5,
			area = $6,
			self_intro = $7
		WHERE user_id = $8
		RETURNING id`

	updateUserPreparationQuery = `
		UPDATE users
		SET is_preparation = $1
		WHERE id = $2
		RETURNING id
	`
)

type InitSetUserInfoRequest struct {
	LastName  string `json:"lastname" validate:"required,name"`
	FirstName string `json:"firstname" validate:"required,name"`
	BirthDate string `json:"birthdate" validate:"required,birthdate"`
	//IsGpsEnabled bool   `json:"isGpsEnabled"`
	Gender    string `json:"gender" validate:"required,oneof=male female"`
	Sexuality string `json:"sexuality" validate:"required,oneof=male female male/female"`
	Area      string `json:"area" validate:"required,area"`
	SelfIntro string `json:"self_intro" validate:"required,self_intro"`
}

func InitSetUserInfo(db *sql.DB) echo.HandlerFunc {
	uploadDir := os.Getenv("IMAGE_UPLOAD_DIR1")
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

		req := new(InitSetUserInfoRequest)
		req.LastName = c.FormValue("lastname")
		req.FirstName = c.FormValue("firstname")
		req.BirthDate = c.FormValue("birthdate")
		//req.IsGpsEnabled = c.FormValue("isGpsEnabled") == "true"
		req.Gender = c.FormValue("gender")
		req.Sexuality = c.FormValue("sexuality")
		req.Area = c.FormValue("area")
		req.SelfIntro = c.FormValue("self_intro")

		// validationをここで行う
		if err := c.Validate(req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		}
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

		if _, err := tx.Exec(setUserInfoQuery, req.LastName, req.FirstName, req.BirthDate, req.Gender, req.Sexuality, req.Area, req.SelfIntro, userID); err != nil {
			os.Remove(filePath)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		// DB更新
		if status, err := executeImageUpdate(tx, filePath, userID, ImageOne); err != nil {
			// エラー時は新しくアップロードしたファイルを削除
			os.Remove(filePath)
			return c.JSON(status, map[string]string{"error": err.Error()})
		}

		if _, err := tx.Exec(updateUserPreparationQuery, true, userID); err != nil {
			os.Remove(filePath)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		// トランザクションのコミット
		if err = tx.Commit(); err != nil {
			os.Remove(filePath)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not commit transaction"})
		}
		return c.JSON(http.StatusOK, map[string]string{"message": "Update user_info successfully."})
	}
}
