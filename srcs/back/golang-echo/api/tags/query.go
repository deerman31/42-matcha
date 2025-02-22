package tags

import (
	"database/sql"
	"fmt"
)

func addTag(tx *sql.Tx, tagName string) error {
	const CheckExistingTag = `
        SELECT EXISTS(
			SELECT 1 FROM tags
			WHERE tag_name = $1
		) as tag_exists
    `
	var exists bool
	err := tx.QueryRow(CheckExistingTag, tagName).Scan(&exists)
	if err != nil {
		if err != sql.ErrNoRows {
			// 予期しないエラーの場合
			return err
		}
	}
	if exists {
		// タグが既に存在する場合は何もしない
		return nil
	}
	const InsertTagQuery = `
INSERT INTO tags (tag_name) 
        VALUES ($1)
        ON CONFLICT (tag_name) DO NOTHING
`
	_, err = tx.Exec(InsertTagQuery, tagName)
	return err
}

func getTagIDByName(tx *sql.Tx, name string) (int, error) {
	// タグ追加クエリ
	// ON DUPLICATE KEY UPDATEは一意制約（UNIQUE制約）に違反する場合の処理を指定する句です。
	// 同じnameのタグが存在しない → 通常の INSERT実行
	// 同じnameのタグが既に存在 → updated_atをCURRENT_TIMESTAMPに更新
	// tagのtagのidを取得するquery
	const GetTagIDByNameQuery = `SELECT id FROM tags WHERE tag_name = $1`
	var id int
	err := tx.QueryRow(GetTagIDByNameQuery, name).Scan(&id)
	return id, err
}

func addUserTag(tx *sql.Tx, userID, tagID int) error {

	const InsertUserTagQuery = `
INSERT INTO user_tags (user_id, tag_id) 
        VALUES ($1, $2)
        ON CONFLICT (user_id, tag_id) DO NOTHING
`
	_, err := tx.Exec(InsertUserTagQuery, userID, tagID)
	return err
}

func deleteUserTag(tx *sql.Tx, userID, tagID int) error {
	const DeleteUserTagQuery = `DELETE FROM user_tags WHERE user_id = $1 AND tag_id = $2`

	_, err := tx.Exec(DeleteUserTagQuery, userID, tagID)
	if err != nil {
		return fmt.Errorf("failed to delete user tag: %v", err)
	}
	return nil
}

func existsTagInUserTags(tx *sql.Tx, tagID int) (bool, error) {
	const CheckTagExistsQuery = `SELECT EXISTS (SELECT 1 FROM user_tags WHERE tag_id = $1)`

	var exists bool
	err := tx.QueryRow(CheckTagExistsQuery, tagID).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("failed to check tag existence: %v", err)
	}
	return exists, nil
}

func deleteTag(tx *sql.Tx, tagID int) error {
	const DeleteTagByIdQuery = `DELETE FROM tags WHERE id = $1`

	_, err := tx.Exec(DeleteTagByIdQuery, tagID)
	if err != nil {
		return fmt.Errorf("failed to delete user tag: %v", err)
	}
	return nil
}

func getUserTags(tx *sql.Tx, userID int) ([]string, error) {
	const GetUserTagsQuery = `
        SELECT t.tag_name 
        FROM user_tags ut 
        JOIN tags t ON ut.tag_id = t.id 
        WHERE ut.user_id = $1
        ORDER BY t.tag_name
    `

	// クエリを実行
	rows, err := tx.Query(GetUserTagsQuery, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to query user tags: %v", err)
	}
	defer rows.Close()

	// タグ名を格納するスライス
	var tags []string

	// 結果を処理
	for rows.Next() {
		var tagName string
		if err := rows.Scan(&tagName); err != nil {
			return nil, fmt.Errorf("failed to scan tag name: %v", err)
		}
		tags = append(tags, tagName)
	}

	// rows.Next()のエラーチェック
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating tag rows: %v", err)
	}

	return tags, nil
}
