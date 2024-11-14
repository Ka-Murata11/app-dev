package model

type LoginRequest struct {
	Email string `query:"email" validate:"required,email"`
	// Password string `json:"password" query:"password"`
}
