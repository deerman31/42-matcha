package browse

import (
	"database/sql"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"time"
)

type BrowseService struct {
	db *sql.DB
}

func NewBrowseService(db *sql.DB) *BrowseService {
	return &BrowseService{db: db}
}

func (b *BrowseService) GetBrowseUser(req BrowseRequest, myID int) ([]userInfo, error) {
	tx, err := b.db.Begin()
	if err != nil {
		return nil, ErrTransactionFailed
	}
	defer tx.Rollback() // エラーが発生した場合はロールバック

	matchUsers, err := GetMatchingUsers(tx, myID, req)
	if err != nil {
		return nil, err
	}
	var userInfos []userInfo
	for _, u := range matchUsers {
		var user userInfo
		user.UserName = u.Username
		user.Age = getAgeHelper(u.Birthdate)
		user.DistanceKm = int(u.DistanceKm)
		user.CommonTagCount = u.CommonTagCount
		user.FameRating = u.FameRating
		imageURI, err := convertImageToDataURI(u.ProfileImagePath1)
		if err != nil {
			return nil, err
		}
		user.ImageURI = imageURI
		userInfos = append(userInfos, user)
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %v", err)
	}
	return userInfos, nil
}

func convertImageToDataURI(imagePath *string) (string, error) {
	if imagePath == nil {
		return "", nil
	}
	imageData, err := os.ReadFile(*imagePath)
	if err != nil {
		return "", fmt.Errorf("failed to read image: %v", err)
	}
	// MIMEタイプを検出
	mimeType := http.DetectContentType(imageData)

	// Base64エンコード
	base64Data := base64.StdEncoding.EncodeToString(imageData)
	// データURIスキーマを作成
	return fmt.Sprintf("data:%s;base64,%s", mimeType, base64Data), nil
}

func getAgeHelper(birthdate time.Time) int {
	now := time.Now()
	age := now.Year() - birthdate.Year()
	// 今年の誕生日がまだ来ていない場合は1歳引く
	if now.YearDay() < birthdate.YearDay() {
		age -= 1
	}
	return age
}
