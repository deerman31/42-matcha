package schemas

const (
	RegisterSuccessMessage = "user created successfully. please check your email to verify your account"
	LogoutSuccessMessage   = "user logout successfully"
	PasswordNoMatchMessage = "password and confirm password do not match"
)

// 各エンドポイントの成功時のレスポンスデータ
type LoginResponse struct {
	IsPreparation bool   `json:"is_preparation"` // 初回ログインが済んでいるかどうか
	AccessToken   string `json:"access_token"`
}

type User struct {
	ID            int
	Username      string
	PasswordHash  string
	IsOnline      bool
	IsRegistered  bool
	IsPreparation bool
}

type RegisterRequest struct {
	Username   string `json:"username" validate:"required,username"`
	Email      string `json:"email" validate:"required,email"`
	Password   string `json:"password" validate:"required,password"`
	RePassword string `json:"repassword" validate:"required,password"`
}

type LoginRequest struct {
	Username string `json:"username" validate:"required,username"`
	Password string `json:"password" validate:"required,password"`
}
