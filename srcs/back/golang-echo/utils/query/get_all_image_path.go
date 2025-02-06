package query

import (
	"database/sql"
	"os"
)

const query = `
SELECT profile_image_path1, profile_image_path2, profile_image_path3, profile_image_path4, profile_image_path5 FROM user_image WHERE user_id = $1
`

func GetAllImagePath(tx *sql.Tx, userID int) ([]string, error) {
	var paths [5]sql.NullString
	err := tx.QueryRow(query, userID).Scan(&paths[0], &paths[1], &paths[2], &paths[3], &paths[4])
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	var retAllImagePath []string
	for _, path := range paths {
		if path.Valid && path.String != "" {
			retAllImagePath = append(retAllImagePath, path.String)
		} else {
			retAllImagePath = append(retAllImagePath, os.Getenv("DEFAULT_IMAGE"))
		}
	}
	return retAllImagePath, nil
}
