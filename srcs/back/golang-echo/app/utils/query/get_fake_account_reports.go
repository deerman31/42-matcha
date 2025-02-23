package query

import (
	"database/sql"
	"fmt"
)

const getFakeAccountReportsQuery = `
    SELECT 
        COUNT(DISTINCT reporter_id) as report_count
    FROM 
        report_fake_accounts
    WHERE 
        fake_account_id = $1`

// GetFakeAccountReports 指定されたユーザーが偽アカウントとして報告された総数を取得する
func GetFakeAccountReports(tx *sql.Tx, userID int) (int, error) {
	var reports int

	err := tx.QueryRow(getFakeAccountReportsQuery, userID).Scan(&reports)
	if err != nil {
		return 0, fmt.Errorf("error querying fake account reports: %w", err)
	}

	return reports, nil
}
