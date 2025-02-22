package response

const (
	RegisterSuccessMessage = "user created successfully. please check your email to verify your account"
	LogoutSuccessMessage   = "user logout successfully"
	PasswordNoMatchMessage = "password and confirm password do not match"
)

// 各エンドポイントの成功時のレスポンスデータ
type RegisterData struct {
	Message string `json:"message"`
}
type LoginData struct {
	IsPreparation bool   `json:"is_preparation,omitempty"` // 初回ログインが済んでいるかどうか
	AccessToken   string `json:"access_token,omitempty"`
}
type LogoutData struct {
	Message string `json:"message"`
}
