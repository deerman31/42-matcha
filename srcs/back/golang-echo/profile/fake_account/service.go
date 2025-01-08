package fakeaccount

import (
	"database/sql"
)

type FakeAccountService struct {
	db *sql.DB
}

func NewFakeAccountService(db *sql.DB) *FakeAccountService {
	return &FakeAccountService{db: db}
}
