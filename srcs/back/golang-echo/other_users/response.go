package otherusers

type OtherGetImageResponse struct {
	Image string `json:"image,omitempty"`
	Error string `json:"error,omitempty"`
}

/*
目的：paramで受け取ったuserの情報を取得し、それをレスポンスする。
1. paramからusernameを取得する
2. そのusernameからuser_idを取得する
3. user_idを使って、
user_infoから(birthdate(これは途中で年齢に変換), gender, sexuality, area, self_intro),
(* これは今回しない)user_imageから5つ(profile_image_path),
user_tagsから(tag_id(あとでこれを元にtag_nameを取得する))
user_locaitonから距離を取得
fame_ratingも取得する
*/
type OtherProfile struct {
	UserName   string   `json:"username"`
	Age        int      `json:"age"`
	Gender     string   `json:"gender"`
	Sexuality  string   `json:"sexuality"`
	Area       string   `json:"area"`
	SelfIntro  string   `json:"self_intro"`
	Tags       []string `json:"tags"`
	Distance   int      `json:"distance"`
	FameRating int      `json:"fame_rating"`
}

type OtherGetProfileResponse struct {
	OtherProfile OtherProfile `json:"other_profile,omitempty"`
	Error        string       `json:"error,omitempty"`
}

type GetOtherAllImageResponse struct {
	AllImage []string `json:"all_image,omitempty"`
	Error    string   `json:"error,omitempty"`
}
