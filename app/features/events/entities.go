package events

import (
	"github.com/labstack/echo/v4"
)

type Core struct {
	ID          int
	Title       string
	Description string
	EventDate   string
	EventTime   string
	Status      string
	Category    string
	Location    string
	Image       string
}

type Repository interface {
	CreateEvent(newEvent Core) (Core, error)
}

type Service interface {
	CreateEvent(newEvent Core) error
}

type Handler interface {
	CreateEvent() echo.HandlerFunc
}
