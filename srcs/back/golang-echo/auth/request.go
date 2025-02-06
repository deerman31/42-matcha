package auth

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