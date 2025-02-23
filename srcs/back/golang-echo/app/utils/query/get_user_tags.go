package query

import (
	"database/sql"
	"fmt"

	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

const queryGetUserTagsByUserID = `
	SELECT 
		ARRAY_AGG(t.tag_name) as tags
	FROM 
		user_tags ut
		INNER JOIN tags t ON ut.tag_id = t.id
	WHERE 
		ut.user_id = $1  -- ここにユーザーIDを指定
	GROUP BY 
		ut.user_id;
	`

func GetUserTags(tx *sql.Tx, userID int) ([]string, error) {
	var tags []string

	err := tx.QueryRow(queryGetUserTagsByUserID, userID).Scan((pq.Array)(&tags))
	if err != nil {
		if err == sql.ErrNoRows {
			return []string{}, nil // ユーザーにタグがない場合は空のスライスを返す
		}
		return nil, fmt.Errorf("failed to get tags: %w", err)
	}
	return tags, nil
}
