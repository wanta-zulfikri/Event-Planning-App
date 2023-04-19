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
	GetProfile(userID int) (Core, error)
	UpdateProfile(id uint, updatedUser Core) error
	DeleteProfile(userID uint) error
}

type Service interface {
	Register(newUser Core) error
	Login(email string, password string) (Core, error)
	GetProfile(userID int) (Core, error)
	UpdateProfile(id uint, username, email, password string) error
	DeleteProfile(userID uint) error
}

type Handler interface {
	Register() echo.HandlerFunc
	Login() echo.HandlerFunc
	GetProfile() echo.HandlerFunc
	UpdateProfile() echo.HandlerFunc
	DeleteProfile() echo.HandlerFunc
}
