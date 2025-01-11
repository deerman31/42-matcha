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
		if err := register(tx, i, GMale); err != nil {
			return err
		}
	}
	for i := 0; i < 2500; i += 1 {
		if err := register(tx, i, GMale); err != nil {
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
	_, err = tx.Query(query, userName, email, password, lastName, firstName, birthDate, sexuality, area, selfIntro, imagePath)
	if err != nil {
		return err
	}
	return nil
}
