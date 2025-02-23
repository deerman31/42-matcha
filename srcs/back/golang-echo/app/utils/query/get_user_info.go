package query

import "database/sql"

const queryGetUserInfoByUserID = `SELECT lastname, firstname, birthdate, gender, sexuality, area, self_intro FROM user_info WHERE user_id = $1`

type UserInfo struct {
	LastName   string
	FirstName  string
	BirthDate  string
	Gender     string
	Sexuality  string
	Area       string
	Self_intro string
}

func GetUserInfo(tx *sql.Tx, userID int) (UserInfo, error) {
	var userInfo UserInfo
	// var lastName, firstName, birthDate, gender, sexuality, area, self_intro string
	if err := tx.QueryRow(queryGetUserInfoByUserID, userID).
		Scan(&userInfo.LastName, &userInfo.FirstName, &userInfo.BirthDate,
			&userInfo.Gender, &userInfo.Sexuality, &userInfo.Area,
			&userInfo.Self_intro); err != nil {
		return UserInfo{}, err
	}
	return userInfo, nil
}
