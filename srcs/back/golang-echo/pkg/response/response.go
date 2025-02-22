package response

const (
	ServErrMessage         = "internal server error"
	InvalidRequestMessage  = "invalid request body"
)

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}