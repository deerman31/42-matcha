package dev

import "database/sql"

type FiveThousandRegisterService struct {
	db *sql.DB
}

func NewFiveThousandRegisterService(db *sql.DB) *FiveThousandRegisterService {
	return &FiveThousandRegisterService{db: db}
}
