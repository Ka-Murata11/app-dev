package model

type SignInRequest struct {
	UserID   string `json:"user_id" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type SignUpRequest struct {
	UserID   string `json:"user_id" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Job      string `json:"job"`
}
