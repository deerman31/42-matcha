package research

import "database/sql"

type ResearchService struct {
	db *sql.DB
}

func NewResearchService(db *sql.DB) *ResearchService {
	return &ResearchService{db: db}
}
