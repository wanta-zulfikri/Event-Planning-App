package events

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Core struct {
	ID          uint
	Title       string
	Description string
	Hosted_by   string
	EventDate   string
	EventTime   string
	Status      string
	Category    string
	Location    string
	Image       string
}

type Event struct {
	gorm.Model
	Title       string `gorm:"type:varchar(100)"`
	Description string `gorm:"type:varchar(225)"`
	EventDate   string `gorm:"type:varchar(20)"`
	EventTime   string `gorm:"type:varchar(20)"`
	Status      string `gorm:"type:varchar(20)"`
	Category    string `gorm:"type:varchar(20)"`
	Location    string `gorm:"type:varchar(100)"`
	Image       string `gorm:"type:varchar(100)"`
}

type Repository interface {
	GetEvents() ([]Core, error)
	CreateEvent(newEvent Core) (Core, error)
	GetEvent(id uint) (Core, error)
	UpdateEvent(id uint, updatedEvent Core) error
	DeleteEvent(id uint) error
}

type Service interface {
	GetEvents() ([]Core, error)
	CreateEvent(newEvent Core) error
	GetEvent(id uint) (Core, error)
	UpdateEvent(id uint, updatedEvent Core) error
	DeleteEvent(id uint) error
}

type Handler interface {
	GetEvents() echo.HandlerFunc
	CreateEvent() echo.HandlerFunc
	GetEvent() echo.HandlerFunc
	UpdateEvent() echo.HandlerFunc
	DeleteEvent() echo.HandlerFunc
}
