package events

import (
	"github.com/labstack/echo/v4"
)

type Core struct {
	ID           uint
	Title        string
	Description  string
	EventDate    string
	EventTime    string
	Status       string
	Category     string
	Location     string
	Image        string
	Hostedby     string //hostedby : username didapat dari JWT Token
	UserID       uint
	Transactions []Transaction
	Attendances  []Attendances
	Reviews      []Reviews
	Tickets      []TicketCore
}

type TicketCore struct {
	ID             uint
	EventID        uint
	TicketCategory string
	TicketPrice    uint
	TicketQuantity uint
}

type Attendances struct {
	Username     string
	User_picture string
}

type Reviews struct {
	Username     string
	User_picture string
	Review       string
}

type Transaction struct {
	ID      uint
	UserID  uint
	EventID uint
}

type Repository interface {
	CreateEventWithTickets(event Core, userID uint) error
	GetEvents() ([]Core, error)
	GetEventsByCategory(category string) ([]Core, error)
	GetEventsByUserID(userid uint) ([]Core, error)
	GetEvent(eventid uint) (Core, error)
	UpdateEvent(id uint, updatedEvent Core) error
	DeleteEvent(id uint) error
}

type Service interface {
	CreateEventWithTickets(event Core, userID uint) error
	GetEvents() ([]Core, error)
	GetEventsByCategory(category string) ([]Core, error)
	GetEventsByUserID(userid uint) ([]Core, error)
	GetEvent(eventid uint) (Core, error)
	UpdateEvent(id uint, updatedEvent Core) error
	DeleteEvent(id uint) error
}

type Handler interface {
	CreateEventWithTickets() echo.HandlerFunc
	GetEvents() echo.HandlerFunc
	GetEventsByUserID() echo.HandlerFunc
	GetEvent() echo.HandlerFunc
	UpdateEvent() echo.HandlerFunc
	DeleteEvent() echo.HandlerFunc
}
