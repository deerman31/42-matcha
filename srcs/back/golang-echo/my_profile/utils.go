package myprofile

import (
	"database/sql"
	"fmt"
)

type userInfo struct {
	userName  string
	email     string
	lastName  string
	firstName string
	birthDate string
	gender    string
	sexuality string
	area      string
	selfIntro string
}

func getUserInfo(tx *sql.Tx, myID int) (userInfo, error) {
	var user userInfo
	err := tx.QueryRow(userInfoQuery, myID).Scan(
		&user.userName,
		&user.email,
		&user.lastName,
		&user.firstName,
		&user.birthDate,
		&user.gender,
		&user.sexuality,
		&user.area,
		&user.selfIntro)
	if err != nil {
		if err == sql.ErrNoRows {
			return userInfo{}, ErrUserNotFound
		}
		return userInfo{}, fmt.Errorf("error querying user info: %w", err)
	}
	return user, err
}



// GetMonthlyProfileViews 指定されたユーザーの過去1ヶ月間のプロフィール閲覧数を取得する
func getMonthlyProfileViews(tx *sql.Tx, userID int) (int, error) {
	count := 0
	err := tx.QueryRow(getMonthlyProfileViewsQuery, userID).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("error querying monthly profile views: %w", err)
	}
	return count, nil
}

// GetMonthlyLikes 指定されたユーザーの過去1ヶ月間のLike数を取得する
func getMonthlyLikes(tx *sql.Tx, userID int) (int, error) {
	var likes int

	err := tx.QueryRow(getMonthlyLikesQuery, userID).Scan(&likes)
	if err != nil {
		return 0, fmt.Errorf("error querying monthly likes: %w", err)
	}
	return likes, nil
}



const (
	viewedPoint      = 1
	likedPoint       = 3
	friendPoint      = 5
	blockPoint       = 5
	fakeAccountPoint = 5
)

func fameRatingCalculation(vieweds, likeds, friends, blocks, fakeAccounts int) int {
	point := (vieweds * viewedPoint) + (likeds * likedPoint) + (friends * friendPoint) - (blocks * blockPoint) - (fakeAccounts * fakeAccountPoint)

	if point >= 100 {
		return 5
	} else if point >= 80 {
		return 4
	} else if point >= 60 {
		return 3
	} else if point >= 40 {
		return 2
	} else if point >= 20 {
		return 1
	} else {
		return 0
	}
}
