package myprofile

import "database/sql"

type MyProfileService struct {
	db *sql.DB
}

func  NewMyProfileService(db *sql.DB) *MyProfileService {
	return &MyProfileService{db: db}
}
