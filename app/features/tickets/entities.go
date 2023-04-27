package tickets

import (
	"github.com/labstack/echo/v4"
)

type Core struct {
	ID             uint
	TicketCategory string
	TicketPrice    uint
	TicketQuantity uint
	EventID        uint
	TransactionID  uint
}

type EventCore struct {
	ID          uint
	Title       string
	Description string
	EventDate   string
	EventTime   string
	Status      string
	Category    string
	Location    string
	Image       string
	Hostedby    string //hostedby : username didapat dari JWT Token
	UserID      uint
	Tickets     []Core `gorm:"foreignKey:EventID"`
}

type Repository interface {
	GetTickets(id uint) ([]Core, error)
	UpdateTicket(eventid uint, updatedTickets []Core) error
	DeleteTicket(id uint) error
}

type Service interface {
	GetTickets(id uint) ([]Core, error)
	UpdateTicket(eventid uint, updatedTicket []Core) error
	DeleteTicket(id uint) error
}

type Handler interface {
	GetTickets() echo.HandlerFunc
	UpdateTicket() echo.HandlerFunc
	DeleteTicket() echo.HandlerFunc
}
