package handler

type LoginResponse struct {
	Token string `json:"token"`
}

type UserResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Image    string `json:"image"`
}
