package browse

import (
	"database/sql"
	"fmt"
	"time"
)

type myInfo struct {
	Age       int
	Gender    string
	Sexuality string
	Area      string
	Tag_ids   []int
	Latitude  float64
	Longitude float64
}

// tx *sql.Tx
func getMyInfo(tx *sql.Tx, myID int) (myInfo, error) {
	age, gender, sexuality, area, err := getUserInfo(tx, myID)
	if err != nil {
		return myInfo{Age: age, Gender: gender, Sexuality: sexuality, Area: area, Tag_ids: nil, Latitude: 0.0, Longitude: 0.0}, err
	}
	tags, err := getTagIds(tx, myID)
	if err != nil {
		return myInfo{Age: age, Gender: gender, Sexuality: sexuality, Area: area, Tag_ids: nil, Latitude: 0.0, Longitude: 0.0}, err
	}
	Latitude, Longitude, err := getLocation(tx, myID)
	if err != nil {
		return myInfo{Age: age, Gender: gender, Sexuality: sexuality, Area: area, Tag_ids: nil, Latitude: 0.0, Longitude: 0.0}, err
	}
	return myInfo{Age: age, Gender: gender, Sexuality: sexuality, Area: area, Tag_ids: tags, Latitude: Latitude, Longitude: Longitude}, nil
}

func getTagIds(tx *sql.Tx, userID int) ([]int, error) {
	const query = `
        SELECT t.tag_id 
        FROM user_tags ut 
        JOIN tags t ON ut.tag_id = t.id 
        WHERE ut.user_id = $1
        ORDER BY t.tag_name
    `
	// クエリを実行
	rows, err := tx.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to query user tags: %v", err)
	}
	defer rows.Close()

	// タグ名を格納するスライス
	var tags []int

	// 結果を処理
	for rows.Next() {
		var tagID int
		if err := rows.Scan(&tagID); err != nil {
			return nil, fmt.Errorf("failed to scan tag name: %v", err)
		}
		tags = append(tags, tagID)
	}

	// rows.Next()のエラーチェック
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating tag rows: %v", err)
	}

	return tags, nil

}

func getLocation(tx *sql.Tx, userID int) (float64, float64, error) {
	const query = `
SELECT 
    CASE 
        WHEN is_gps = TRUE THEN ST_X(location::geometry)
        ELSE ST_X(location_alternative::geometry)
    END as longitude,
    CASE 
        WHEN is_gps = TRUE THEN ST_Y(location::geometry)
        ELSE ST_Y(location_alternative::geometry)
    END as latitude
FROM user_location 
WHERE user_id = $1;
`
	var Latitude, Longitude float64
	if err := tx.QueryRow(query, userID).Scan(&Longitude, &Latitude); err != nil {
		return 0.0, 0.0, err
	}
	return Latitude, Longitude, nil
}

// "SELECT %s FROM %s WHERE %s = $1;", param.FieldName, param.TableName, param.Where)
func getUserInfo(tx *sql.Tx, myID int) (int, string, string, string, error) {
	const query = `
SELECT birthdate, gender, sexuality, area 
FROM user_info 
WHERE user_id = $1;`

	var birthdate, gender, sexuality, area string
	if err := tx.QueryRow(query, myID).Scan(&birthdate, &gender, &sexuality, &area); err != nil {
		return 0, "", "", "", err
		//return c.JSON(http.StatusInternalServerError, map[string]string{"error": fmt.Sprintf("failed to query user tags: %v", err)})
	}
	age := getAge(birthdate)

	return age, gender, sexuality, area, nil
}

func getAgeHelper(birthdate time.Time) int {
	now := time.Now()
	age := now.Year() - birthdate.Year()
	// 今年の誕生日がまだ来ていない場合は1歳引く
	if now.YearDay() < birthdate.YearDay() {
		age -= 1
	}
	return age
}
func getAge(birthdate string) int {
	t, _ := time.Parse("2006-01-02", birthdate)
	return getAgeHelper(t)
}
