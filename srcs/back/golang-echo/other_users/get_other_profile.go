package otherusers

import (
	"golang-echo/jwt_token"
	"golang-echo/utils"
	"golang-echo/utils/query"
	"net/http"

	"github.com/labstack/echo/v4"
)

/*
目的：paramで受け取ったuserの情報を取得し、それをレスポンスする。
1. paramからusernameを取得する
2. そのusernameからuser_idを取得する
3. user_idを使って、
user_infoから(birthdate(これは途中で年齢に変換), gender, sexuality, area, self_intro),
(* これは今回しない)user_imageから5つ(profile_image_path),
user_tagsから(tag_id(あとでこれを元にtag_nameを取得する))
user_locaitonから距離を取得
fame_ratingも取得する
*/

func (o *OtherUsersHandler) OtherGetProfile(c echo.Context) error {
	claims, ok := c.Get("user").(*jwt_token.Claims)
	if !ok {
		return c.JSON(http.StatusInternalServerError, OtherGetProfileResponse{Error: "user claims not found"})
	}

	username := c.Param("name")
	if err := c.Validate(struct {
		Username string `json:"username" validate:"required,username"`
	}{Username: username}); err != nil {
		return c.JSON(http.StatusBadRequest, OtherGetProfileResponse{Error: err.Error()})
	}

	profile, err := o.service.OtherGetProfile(claims.UserID, username)
	if err != nil {
		switch err {
		case ErrFailedToGetOtherUserProfile:
			return c.JSON(http.StatusInternalServerError, OtherGetProfileResponse{Error: err.Error()})
		default:
			return c.JSON(http.StatusInternalServerError, OtherGetProfileResponse{Error: err.Error()})
		}
	}
	return c.JSON(http.StatusOK, OtherGetProfileResponse{OtherProfile: profile})
}

func (o *OtherUsersService) OtherGetProfile(myID int, userName string) (OtherProfile, error) {
	tx, err := o.db.Begin()
	if err != nil {
		return OtherProfile{}, ErrTransactionFailed
	}
	defer tx.Rollback()
	otherID, err := query.GetUserIDByUsername(tx, userName)
	if err != nil {
		return OtherProfile{}, err
	}
	userInfo, err := query.GetUserInfo(tx, otherID)
	if err != nil {
		return OtherProfile{}, err
	}
	tags, err := query.GetUserTags(tx, otherID)
	if err != nil {
		return OtherProfile{}, err
	}
	distance, err := query.CalculateDistanceBetweenUsers(tx, myID, otherID)
	if err != nil {
		return OtherProfile{}, err
	}
	fameRating, err := query.GetFameRating(tx, otherID)
	if err != nil {
		return OtherProfile{}, err
	}
	age := utils.CalculateAgeFromBirthDate(userInfo.BirthDate)

	other := OtherProfile{UserName: userName, Age: age,
		Gender: userInfo.Gender, Sexuality: userInfo.Sexuality, Area: userInfo.Area, SelfIntro: userInfo.Self_intro,
		Tags: tags, Distance: distance, FameRating: fameRating}

	return other, nil
}
