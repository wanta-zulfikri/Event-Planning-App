package events

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
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
	Tickets     []TicketCore
}

type TicketCore struct { //new struct TicketCore
	Title          string
	TicketType     string
	TicketCategory string
	TicketPrice    uint
	TicketQuantity uint
}

type Repository interface {
	CreateEventWithTickets(tx *gorm.DB, event Core, userID uint) error
	CreateEvent(tx *gorm.DB, event Core, userID uint) error
	CreateTickets(tx *gorm.DB, event Core, eventID uint) error
}

type Service interface {
	CreateEventWithTickets(event Core, userID uint) error
}

type Handler interface {
	CreateEventWithTickets() echo.HandlerFunc
}
