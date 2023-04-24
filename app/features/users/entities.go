package users

import (
	"github.com/labstack/echo/v4"
	"github.com/wanta-zulfikri/Event-Planning-App/app/features/events"
)

type Core struct {
	ID       uint
	Username string
	Email    string
	Password string
	Image    string
	Events   []events.Core
}

type Repository interface {
	Register(newUser Core) (Core, error)
	Login(email, password string) (Core, error)
	GetProfile(id uint) (Core, error)
	UpdateProfile(id uint, updatedUser Core) error
	DeleteProfile(id uint) error
}

type Service interface {
	Register(newUser Core) error
	Login(email string, password string) (Core, error)
	GetProfile(id uint) (Core, error)
	UpdateProfile(id uint, username, newEmail, password, image string) error
	DeleteProfile(id uint) error
}

type Handler interface {
	Register() echo.HandlerFunc
	Login() echo.HandlerFunc
	GetProfile() echo.HandlerFunc
	UpdateProfile() echo.HandlerFunc
	DeleteProfile() echo.HandlerFunc
}
