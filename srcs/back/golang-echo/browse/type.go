package browse

import "time"

type userInfo struct {
	UserName       string // username
	Age            int    // 年齢
	DistanceKm     int    // 自分と相手との距離
	CommonTagCount int    // 共通タグの数
	FameRating     int    // fame_raging
	ImageURI       string
}

type BrowseResponse struct {
	UserInfos []userInfo `json:"user_infos,omitempty"`
	Error     string     `json:"error,omitempty"`
}

type MatchingUser struct {
	Username          string    `db:"username"`
	Birthdate         time.Time `db:"birthdate"`
	Area          string    `db:"area"`
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
		Min int `json:"min" validate:"required,gte=18"`              // 最小年齢バリデーション追加
		Max int `json:"max" validate:"required,lte=100,gtfield=Min"` // 最大年齢バリデーション追加
	} `json:"age_range" validate:"required"`
	DistanceRange struct {
		Min int `json:"min" validate:"required,gte=0"`
		Max int `json:"max" validate:"required,lte=100,gtfield=Min"` //最大100kmまで
	} `json:"distance_range" validate:"required"`
	MinCommonTags int            `json:"min_common_tags" validate:"required,gte=0,lte=5"`
	MinFameRating int            `json:"min_fame_rating" validate:"required,gte=0,lte=5"`
	SortOption    SortOptionType `json:"sort_option" validate:"required,oneof=age distance fame_rating tag"`
	SortOrder     SortOrder      `json:"sort_order" validate:"required,oneof=0 1"`
}

// filter.MinCommonTags
// filter.MinFameRating,
