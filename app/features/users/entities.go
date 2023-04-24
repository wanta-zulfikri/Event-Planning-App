package users

import (
	"github.com/labstack/echo/v4"
	"github.com/wanta-zulfikri/Event-Planning-App/app/features/events"
	"gorm.io/gorm"
)

type Core struct {
	ID       uint
	Username string
	Email    string
	Password string
	Image    string
	Events   []events.Core
}

type User struct {
	gorm.Model
	Username string         `json:"username" gorm:"type:varchar(100);not null"`
	Email    string         `json:"email" gorm:"primaryKey"`
	Password string         `json:"password" gorm:"type:varchar(100);not null"`
	Image    string         `json:"image" gorm:"type:varchar(100)"`
	Events   []events.Event `gorm:"foreignKey:UserID"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.Image == "" {
		u.Image = "https://storage.googleapis.com/grouproject/book/images.jpeg"
	}
	return nil
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
	UpdateProfile(email, username, newEmail, password, image string) error
	DeleteProfile(email string) error
}

type Handler interface {
	Register() echo.HandlerFunc
	Login() echo.HandlerFunc
	GetProfile() echo.HandlerFunc
	UpdateProfile() echo.HandlerFunc
	DeleteProfile() echo.HandlerFunc
}
