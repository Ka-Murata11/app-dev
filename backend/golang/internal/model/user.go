package model

type User struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	Job    string `json:"job"`
}

type GetUsersResponse struct {
	Users []User `json:"users"`
}
