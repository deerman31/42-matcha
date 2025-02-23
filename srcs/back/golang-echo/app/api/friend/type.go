package friend


type FriendResponse struct {
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}
type FriendListResponse struct {
	Users []string `json:"users"`
}

type FriendRequest struct {
	Username string `json:"username" validate:"required,username"`
}