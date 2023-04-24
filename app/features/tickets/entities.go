package tickets

import (
	"github.com/labstack/echo/v4"
)

type Core struct {
	ID             uint
	TicketType     string
	TicketCategory string
	TicketPrice    uint
	TicketQuantity uint
	EventID        uint
	TransactionID  uint
}

type Repository interface {
	GetTickets() ([]Core, error)
	CreateTicket(newTicket Core, eventid uint64) (Core, error)
	GetTicket(id uint) (Core, error)
	UpdateTicket(id uint, updatedTicket Core) error
	DeleteTicket(id uint) error
}

type Service interface {
	GetTickets() ([]Core, error)
	CreateTicket(newTicket Core, eventid uint64) error
	GetTicket(id uint) (Core, error)
	UpdateTicket(id uint, updatedTicket Core) error
	DeleteTicket(id uint) error
}

type Handler interface {
	GetTickets() echo.HandlerFunc
	CreateTicket() echo.HandlerFunc
	GetTicket() echo.HandlerFunc
	UpdateTicket() echo.HandlerFunc
	DeleteTicket() echo.HandlerFunc
}
