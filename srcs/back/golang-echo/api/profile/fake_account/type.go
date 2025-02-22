package fakeaccount

type FakeAccountResponse struct {
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}

type FakeAccountRequest struct {
	Username string `json:"username" validate:"required,username"`
}
