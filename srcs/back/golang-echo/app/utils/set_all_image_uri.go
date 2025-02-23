package utils

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
)

func SetAllImageURI(allImagePath []string) ([]string, error) {
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
