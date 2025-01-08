package block

import (
	"database/sql"
)

type BlockService struct {
	db *sql.DB
}

func NewBlockService(db *sql.DB) *BlockService {
	return &BlockService{db: db}
}
