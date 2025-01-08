package fakeaccount

import (
	"database/sql"
	"golang-echo/jwt_token"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

func (f *FakeAccountHandler) ReportFakeAccount(c echo.Context) error {
	claims, ok := c.Get("user").(*jwt_token.Claims)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "user claims not found")
	}
	req := new(FakeAccountRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, FakeAccountResponse{Error: "Invalid request body"})
	}
	// validationをここで行う
	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, FakeAccountResponse{Error: err.Error()})
	}
	err := f.service.ReportFakeAccount(claims.UserID, req.Username)
	if err!=nil{
		switch err{
		case ErrUserNotFound:
			return c.JSON(http.StatusNotFound, FakeAccountResponse{Error: "User not found"})
		case ErrSelfAction:
			return c.JSON(http.StatusBadRequest, FakeAccountResponse{Error: "Cannot fake account your own profile"})
		case ErrAlreadyFakeAccount:
			return c.JSON(http.StatusConflict, FakeAccountResponse{Error: "Already fake account with this user"})
		default:
			return c.JSON(http.StatusInternalServerError, FakeAccountResponse{Error: "Internal server error"})
		}
	}
	return c.JSON(http.StatusCreated, FakeAccountResponse{Message: "Successfully reported fake account the user"})
}

func (f *FakeAccountService) ReportFakeAccount(myID int, userName string) error {
	tx, err := f.db.Begin()
	if err != nil {
		return ErrTransactionFailed
	}
	otherID, err := getUserID(tx, userName)
	if err != nil {
		return err
	}
	if myID == otherID {
		return ErrSelfAction
	}
	if err := createFakeReport(tx, myID, otherID); err != nil {
		return err
	}
	return tx.Commit()
}

func createFakeReport(tx *sql.Tx, repoterID, fakeAccountID int) error {
	const query = `
	INSERT INTO report_fake_accounts (repoter_id, fake_account_id)
	VALUES ($1, $2)`
	_, err := tx.Exec(query, repoterID, fakeAccountID)
	if pqErr, ok := err.(*pq.Error); ok {
		if pqErr.Code == "23505" { // unique_violation のエラーコード
			return ErrAlreadyFakeAccount
		}
	}
	return err
}
