package otherusers

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func (o *OtherUsersHandler) OtherGetImage(c echo.Context) error {
	req := new(OtherGetImageRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, OtherGetImageResponse{Error: "Invalid request body"})
	}

	// validationをここで行う
	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, OtherGetImageResponse{Error: err.Error()})
	}

	retImage, err := o.service.OtherGetImage(req.ImagePath)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, OtherGetImageResponse{Error: "Internal server error"})
	}
	return c.JSON(http.StatusOK, OtherGetImageResponse{Image: retImage})
}

func (o *OtherUsersService) OtherGetImage(imagePath string) (string, error) {
	tx, err := o.db.Begin()
	if err != nil {
		return "", ErrTransactionFailed
	}

	defer tx.Rollback() // エラーが発生した場合はロールバック

	imageData, err := os.ReadFile(imagePath)
	if err != nil {
		return "", ErrReadImageFile
	}

	// MIMEタイプを検出
	mimeType := http.DetectContentType(imageData)
	// Base64エンコード
	base64Data := base64.StdEncoding.EncodeToString(imageData)
	// データURIスキーマを作成
	dataURI := fmt.Sprintf("data:%s;base64,%s", mimeType, base64Data)
	// トランザクションのコミット
	if err = tx.Commit(); err != nil {
		return "", ErrTransactionFailed
	}
	return dataURI, nil
}
