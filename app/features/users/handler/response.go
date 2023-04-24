package handler

type LoginResponse struct {
	Token string `json:"token"`
}

type UserResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Image    string `json:"image"`
}
