package myprofile

type myInfo struct {
	UserName string `json:"username"`
	Email    string `json:"email"`

	LastName  string   `json:"lastname"`
	FirstName string   `json:"firstname"`
	BirthDate string   `json:"birthdate"`
	Gender    string   `json:"gender"`
	Sexuality string   `json:"sexuality"`
	Area      string   `json:"area"`
	SelfIntro string   `json:"self_intro"`
	Tags      []string `json:"tags"`
	IsGPS     bool     `json:"is_gps"`
	Latitude  float64  `json:"latitude"`
	Longitude float64  `json:"longitude"`

	FameRating int `json:"fame_rating"`
}

type MyProfileResponse struct {
	MyInfo myInfo `json:"my_info,omitempty"`
	Error  string `json:"error,omitempty"`
}
