package otherusers

import "database/sql"

type OtherUsersService struct {
	db *sql.DB
}

func NewOtherUsersService(db *sql.DB) *OtherUsersService {
	return &OtherUsersService{db: db}
}
