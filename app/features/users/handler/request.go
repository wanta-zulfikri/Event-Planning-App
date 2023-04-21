package handler

type RegisterInput struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateInput struct {
	Username string `form:"username"`
	Email    string `form:"email"`
	Password string `form:"password"`
	Image    string `form:"image"`
}
