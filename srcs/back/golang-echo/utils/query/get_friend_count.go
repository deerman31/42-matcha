package query

import (
	"database/sql"
	"fmt"
)

// フレンド数を取得するクエリ
const getFriendCountQuery = `
SELECT 
	(
		SELECT COUNT(*) 
		FROM user_friends 
		WHERE user_id1 = $1 
		OR user_id2 = $1
	) as friend_count`

// GetFriendCount 指定されたユーザーのフレンド総数を取得する
func GetFriendCount(tx *sql.Tx, userID int) (int, error) {
	var friends int

	err := tx.QueryRow(getFriendCountQuery, userID).Scan(&friends)
	if err != nil {
		return 0, fmt.Errorf("error querying friend count: %w", err)
	}
	return friends, nil
}
