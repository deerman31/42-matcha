package fakeaccount

import (
	"database/sql"

	"github.com/labstack/echo/v4"
)

func FakeAccountRoutes(protected *echo.Group, db *sql.DB) {
	fake := protected.Group("/fake-account")

	handler := NewFakeAccountHandler(NewFakeAccountService(db))
	fake.POST("/report", handler.ReportFakeAccount)
}
