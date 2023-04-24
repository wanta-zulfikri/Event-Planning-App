package events

import (
	"github.com/labstack/echo/v4"
)

type Core struct {
	ID          uint
	Title       string
	Description string
	EventDate   string
	EventTime   string
	Status      string
	Category    string
	Location    string
	Image       string
	Hostedby    string //hostedby : username didapat dari jwt
	UserID      uint
}

type Repository interface {
	GetEvents() ([]Core, error)
	CreateEvent(newEvent Core, id uint) (Core, error)
	GetEvent(id uint) (Core, error)
	UpdateEvent(id uint, updatedEvent Core) error
	DeleteEvent(id uint) error
}

type Service interface {
	GetEvents() ([]Core, error)
	CreateEvent(newEvent Core, id uint) error
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
