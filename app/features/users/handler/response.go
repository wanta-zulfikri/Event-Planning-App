package handler

type LoginResponse struct {
	Token string `json:"token"`
}

type UserResponse struct {
	Username string `json:"user_name" form:"user_name"`
	Email    string `json:"email" form:"email"`
}
