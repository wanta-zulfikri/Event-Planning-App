package users

import (
	"github.com/labstack/echo/v4"
)

type Core struct {
	ID       uint
	Username string
	Email    string
	Password string
}

type Repository interface {
	Register(newUser Core) (Core, error)
	Login(email, password string) (Core, error)
	GetProfile(email string) (Core, error)
	UpdateProfile(email string, updatedUser Core) error
	DeleteProfile(email string) error
}

type Service interface {
	Register(newUser Core) error
	Login(email string, password string) (Core, error)
	GetProfile(email string) (Core, error)
	UpdateProfile(email, username, newEmail, password string) error
	DeleteProfile(email string) error
}

type Handler interface {
	Register() echo.HandlerFunc
	Login() echo.HandlerFunc
	GetProfile() echo.HandlerFunc
	UpdateProfile() echo.HandlerFunc
	DeleteProfile() echo.HandlerFunc
}
