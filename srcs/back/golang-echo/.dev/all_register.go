package dev

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

// エラーメッセージを一箇所に集約
var (
	ErrTransactionFailed = errors.New("transaction failed")
)

const (
	// 新規ユーザーを登録するためのクエリ
	insertNewUserQuery = `
        INSERT INTO users (
            username, 
            email, 
            password_hash,
			is_registered,
			is_preparation
        ) VALUES ($1, $2, $3, $4, $5)
		 RETURNING id
    `
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

	imageOneQuery = `UPDATE user_image SET profile_image_path1 = $1 WHERE user_id = $2 RETURNING id`
)

func (f *FiveThousandRegisterHandler) AllRegister(c echo.Context) error {
	err := f.service.AllRegister()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, FiveThousandRegisterResponse{Error: "Internal server error"})
	}
	return c.JSON(http.StatusCreated, FiveThousandRegisterResponse{Message: "5000 User created successfully. Please check your email to verify your account."})
}

func (f *FiveThousandRegisterService) AllRegister() error {
	tx, err := f.db.Begin()
	if err != nil {
		return ErrTransactionFailed
	}
	defer tx.Rollback()

	for i := 0; i < 2500; i += 1 {
		fmt.Println("male", i)
		if err := register(tx, i, GMale); err != nil {
			return err
		}
	}
	for i := 0; i < 2500; i += 1 {
		fmt.Println("female", i)
		if err := register(tx, i, GFemale); err != nil {
			return err
		}
	}
	return tx.Commit()
}

func register(tx *sql.Tx, num int, gender GenderType) error {
	userName := getUserName(gender) + strconv.Itoa(num)
	email := fmt.Sprintf("%s%d@ft.com", userName, num)
	password := "ZidaneYkusano42!"
	lastName := "Kusano"
	firstName := "Yoshinari"
	birthDate := getBirthDate()
	sexuality := getSexuality()
	area := getArea()
	selfIntro := "Nice to meet you, my name is Satoshi Nakamoto!"
	imagePath := getImagePath(gender)

	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	password = string(hashedBytes)
	var userID int
	err = tx.QueryRow(insertNewUserQuery, userName, email, password, true, true).Scan(&userID)
	if err != nil {
		return err
	}
	_, err = tx.Exec(setUserInfoQuery, lastName, firstName, birthDate, gender, sexuality, area, selfIntro, userID)
	if err != nil {
		return err
	}
	_, err = tx.Exec(imageOneQuery, imagePath, userID)
	if err != nil {
		return err
	}
	return nil
}
