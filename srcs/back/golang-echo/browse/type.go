package browse

import "time"

type userInfo struct {
	UserName       string `json:"username"` // username
	Age            int `json:"age"`   // 年齢
	DistanceKm     int `json:"distance_km"`   // 自分と相手との距離
	CommonTagCount int `json:"common_tag_count"`   // 共通タグの数
	FameRating     int `json:"Fame_rating"`   // fame_raging
	//ImageURI       string
	ImagePath string `json:"image_path"`
}

type BrowseResponse struct {
	UserInfos []userInfo `json:"user_infos,omitempty"`
	Error     string     `json:"error,omitempty"`
}

type MatchingUser struct {
	Username          string    `db:"username"`
	Birthdate         time.Time `db:"birthdate"`
	Area              string    `db:"area"`
	ProfileImagePath1 *string   `db:"profile_image_path1"`
	CommonTagCount    int       `db:"common_tag_count"`
	DistanceKm        float64   `db:"distance_km"`
	FameRating        int       `db:"fame_rating"`
}

type SortOptionType string

const (
	Distance   SortOptionType = "distance"
	Age        SortOptionType = "age"
	FameRating SortOptionType = "fame_rating"
	Tag        SortOptionType = "tag"
)

type SortOrder int

const (
	Descending SortOrder = iota
	Ascending
)

type BrowseRequest struct {
	AgeRange struct {
		Min int `json:"min"` // 最小年齢バリデーション追加
		Max int `json:"max"` // 最大年齢バリデーション追加
	} `json:"age_range" validate:"required,age_range"`
	DistanceRange struct {
		Min int `json:"min"`
		Max int `json:"max"` //最大100kmまで
	} `json:"distance_range" validate:"required,distance_range"`
	MinCommonTags int            `json:"min_common_tags" validate:"min_common_tags"`
	MinFameRating int            `json:"min_fame_rating" validate:"min_common_tags"`
	SortOption    SortOptionType `json:"sort_option" validate:"required,oneof=age distance fame_rating tag"`
	SortOrder     SortOrder      `json:"sort_order" validate:"oneof=0 1"`
}
