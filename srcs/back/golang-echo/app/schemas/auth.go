package schemas

const (
	RegisterSuccessMessage = "user created successfully. please check your email to verify your account"
	LogoutSuccessMessage   = "user logout successfully"
	PasswordNoMatchMessage = "password and confirm password do not match"
)

// 各エンドポイントの成功時のレスポンスデータ
// @Description ログイン成功時のレスポンス
type LoginResponse struct {
	IsPreparation bool   `json:"is_preparation" example:"false"`                 // 初回ログインが済んでいるかどうか
	AccessToken   string `json:"access_token" example:"eyJhbGciOiJIUzI1NiIs..."` // JWTアクセストークン
}

// RegisterRequest 新規登録リクエスト
// @Description ユーザー登録時のリクエスト内容
type RegisterRequest struct {
	// ユーザー名（3~30文字の英数字とアンダースコア）
	Username string `json:"username" validate:"required,username" example:"ykusano"`
	// メールアドレス
	Email string `json:"email" validate:"required,email" example:"ykusano@test.com"`
	// パスワード（8~30文字の英大小文字・数字・記号を含む）
	Password string `json:"password" validate:"required,password" example:"Password123!"`
	// パスワード（確認用）
	RePassword string `json:"repassword" validate:"required,password" example:"Password123!"`
}

// LoginRequest ログインリクエスト
// @Description ログイン時のリクエスト内容
type LoginRequest struct {
	// ユーザー名
	Username string `json:"username" validate:"required,username" example:"ykusano"`
	// パスワード
	Password string `json:"password" validate:"required,password" example:"Password123!"`
}

type User struct {
	ID            int
	Username      string
	PasswordHash  string
	IsOnline      bool
	IsRegistered  bool
	IsPreparation bool
}
