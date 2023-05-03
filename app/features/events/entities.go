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
	Reviews      []Reviews
	Tickets      []TicketCore
}

type Event struct {
	ID          uint
	Title       string
	Description string
	EventDate   string
	EventTime   string
	Status      string
	Category    string
	Location    string
	Image       string
	Hostedby    string // hostedby : username obtained from JWT Token
	UserID      uint
}

type TicketCore struct {
	ID             uint
	EventID        uint
	TicketCategory string
	TicketPrice    uint
	TicketQuantity uint
}

// kalau mau lookup ke table users,
//untuk mendapatkan user_picture,
//nama object harus sesuai dengan nama table.

type Reviews struct {
	UserID   uint
	Username string
	Image    string
	Review   string
}

type Transaction struct {
	ID       uint
	UserID   uint
	EventID  uint
	Username string
	Image    string
}

type Repository interface {
	CreateEventWithTickets(event Core, userID uint) error
	GetEvents() ([]Event, error)
	GetEventsByCategory(category string) ([]Event, error)
	GetEventsByUserID(userid uint) ([]Core, error)
	GetEvent(eventid uint) (Core, error)
	UpdateEvent(id uint, updatedEvent Core) error
	DeleteEvent(id uint) error
}

type Service interface {
	CreateEventWithTickets(event Core, userID uint) error
	GetEvents() ([]Event, error)
	GetEventsByCategory(category string) ([]Event, error)
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
