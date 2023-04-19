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

type UserUpdateInput struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
