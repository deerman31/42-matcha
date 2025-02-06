package auth

type RegisterResponse struct {
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}

// トークンのレスポンス用構造体を追加
type LoginResponse struct {
	IsPreparation bool   `json:"is_preparation,omitempty"` // 初回ログインが済んでいるかどうか
	AccessToken   string `json:"access_token,omitempty"`
	Error         string `json:"error,omitempty"`
}

type LogoutResponse struct {
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}