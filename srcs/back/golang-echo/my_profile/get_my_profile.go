package myprofile

import (
	"golang-echo/jwt_token"
	"golang-echo/utils/query"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (m *MyProfileHandler) GetMyProfile(c echo.Context) error {
	claims, ok := c.Get("user").(*jwt_token.Claims)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "user claims not found")
	}
	users, err := m.service.GetMyProfile(claims.UserID)
	if err != nil {
		switch err {
		case ErrUserNotFound:
			return c.JSON(http.StatusNotFound, MyProfileResponse{Error: "User not found"})
		default:
			return c.JSON(http.StatusInternalServerError, MyProfileResponse{Error: "Internal server error"})
		}
	}
	return c.JSON(http.StatusOK, MyProfileResponse{MyInfo: users})
}

func (m *MyProfileService) GetMyProfile(myID int) (myInfo, error) {
	tx, err := m.db.Begin()
	if err != nil {
		return myInfo{}, ErrTransactionFailed
	}
	defer tx.Rollback() // エラーが発生した場合はロールバック

	//result := new(myInfo)
	userInfo, err := getUserInfo(tx, myID)
	if err != nil {
		return myInfo{}, err
	}
	tags, err := query.GetUserTags(tx, myID)
	if err != nil {
		return myInfo{}, err
	}

	loc, err := query.GetUserLocation(tx, myID)
	if err != nil {
		return myInfo{}, err
	}
	profileViewsCount, err := getMonthlyProfileViews(tx, myID)
	if err != nil {
		return myInfo{}, err
	}
	likesCount, err := getMonthlyLikes(tx, myID)
	if err != nil {
		return myInfo{}, err
	}
	friendsCount, err := query.GetFriendCount(tx, myID)
	if err != nil {
		return myInfo{}, err
	}
	blocksCount, err := query.GetBlockedCount(tx, myID)
	if err != nil {
		return myInfo{}, err
	}
	reportsCount, err := query.GetFakeAccountReports(tx, myID)
	if err != nil {
		return myInfo{}, err
	}
	err = tx.Commit()
	if err != nil {
		return myInfo{}, err
	}

	return myInfo{
		userInfo.userName,
		userInfo.email,
		userInfo.lastName,
		userInfo.firstName,
		userInfo.birthDate[:10],
		userInfo.gender,
		userInfo.sexuality,
		userInfo.area,
		userInfo.selfIntro,
		tags,
		loc.IsGPS,
		loc.Latitude,
		loc.Longitude,
		fameRatingCalculation(profileViewsCount, likesCount, friendsCount, blocksCount, reportsCount),
	}, nil
}

/*
目的：init.sqlを確認し、下記の要素を取得するquery定数とそのqueryを使用し、取得した値を返すgolangの関数を作成せよ。
使用データベース：postgres

取得するテーブル:
users: username, emailを取得
user_info: lastname, firstname, birthdate, gender, sexuality, area, self_introを取得
user_tags: 自分が持っているtagのリストを取得
user_location: location_alternative, is_gpsを取得
profile_views: 1ヶ月間の間に自分が閲覧された数を取得
user_likes: 1ヶ月間の間に自分がlikeされた数を取得
user_friends: 自分のfriendの数を取得
user_blocks: 自分がblockされた数を取得
report_fake_accounts: 自分が報告された数を取得する
*/
