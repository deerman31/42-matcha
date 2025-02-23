package schemas

const (
	ServErrMessage        = "internal server error"
	InvalidRequestMessage = "invalid request body"
)

type Response struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
