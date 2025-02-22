package block

type BlockResponse struct {
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}

type BlockListResponse struct {
	Users []string `json:"users"`
}

type BlockRequest struct {
	Username string `json:"username" validate:"required,username"`
}
