package query

import (
	"database/sql"
	"fmt"
)

// ブロック数を取得するクエリ
const getBlockedCountQuery = `
    SELECT 
        COUNT(DISTINCT blocker_id) as block_count
    FROM 
        user_blocks
    WHERE 
        blocked_id = $1`

// GetBlockedCount 指定されたユーザーがブロックされた総数を取得する
func GetBlockedCount(tx *sql.Tx, userID int) (int, error) {
	var blocks int

	err := tx.QueryRow(getBlockedCountQuery, userID).Scan(&blocks)
	if err != nil {
		return 0, fmt.Errorf("error querying block count: %w", err)
	}
	return blocks, nil
}
