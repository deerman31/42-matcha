package myprofile

import (
	"database/sql"
	"fmt"

	"github.com/lib/pq"
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

func getMyTags(tx *sql.Tx, myID int) ([]string, error) {
	var tags []string

	err := tx.QueryRow(query2, myID).Scan((pq.Array)(&tags))
	if err != nil {
		if err == sql.ErrNoRows {
			return []string{}, nil // ユーザーにタグがない場合は空のスライスを返す
		}
		return nil, fmt.Errorf("failed to get tags: %w", err)
	}

	return tags, nil
}

// Location情報を格納する構造体
type UserLocation struct {
	Latitude  float64
	Longitude float64
	IsGPS     bool
}

func getUserLocation(tx *sql.Tx, userID int) (UserLocation, error) {
	var loc UserLocation

	err := tx.QueryRow(getUserLocationQuery, userID).Scan(
		&loc.Latitude,
		&loc.Longitude,
		&loc.IsGPS,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return UserLocation{}, fmt.Errorf("location not found for user: %w", err)
		}
		return UserLocation{}, fmt.Errorf("error querying user location: %w", err)
	}
	return loc, nil
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

// GetFriendCount 指定されたユーザーのフレンド総数を取得する
func getFriendCount(tx *sql.Tx, userID int) (int, error) {
	var friends int

	err := tx.QueryRow(getFriendCountQuery, userID).Scan(&friends)
	if err != nil {
		return 0, fmt.Errorf("error querying friend count: %w", err)
	}
	return friends, nil
}

// GetBlockedCount 指定されたユーザーがブロックされた総数を取得する
func getBlockedCount(tx *sql.Tx, userID int) (int, error) {
	var blocks int

	err := tx.QueryRow(getBlockedCountQuery, userID).Scan(&blocks)
	if err != nil {
		return 0, fmt.Errorf("error querying block count: %w", err)
	}
	return blocks, nil
}

// GetFakeAccountReports 指定されたユーザーが偽アカウントとして報告された総数を取得する
func getFakeAccountReports(tx *sql.Tx, userID int) (int, error) {
	var reports int

	err := tx.QueryRow(getFakeAccountReportsQuery, userID).Scan(&reports)
	if err != nil {
		return 0, fmt.Errorf("error querying fake account reports: %w", err)
	}

	return reports, nil
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
