package tags

import "database/sql"

type TagService struct {
	db *sql.DB
}

func NewTagService(db *sql.DB) *TagService {
	return &TagService{db: db}
}
