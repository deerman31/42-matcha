package schemas

const (
	ServErrMessage        = "internal server error"
	InvalidRequestMessage = "invalid request body"
)

// Response 正常系のレスポンス
// @Description 処理成功時の共通レスポンス形式
type Response struct {
	// レスポンスメッセージ
	Message string `json:"message" example:"successfully"`
}

// ErrorResponse エラーレスポンス
// @Description エラー発生時の共通レスポンス形式
type ErrorResponse struct {
	// エラーメッセージ
	Error string `json:"error" example:"error"`
}
