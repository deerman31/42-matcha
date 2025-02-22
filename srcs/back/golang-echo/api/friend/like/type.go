package like

type LikeType string

const (
	LikedUsers LikeType = "liked_users"
	LikerUsers LikeType = "liker_users"
)

type LikeRequest struct {
	Username string `json:"username" validate:"required,username"`
}

type LikeResponse struct {
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}

type LikeListResponse struct {
	Users []string `json:"users"`
}
