package tickets

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Core struct {
	ID             uint
	TicketType     string
	TicketCategory string
	TicketPrice    string
	TicketQuantity string
}

type Ticket struct {
	gorm.Model
	TicketType     string `gorm:"type:varchar(20)"`
	TicketCategory string `gorm:"type:varchar(20)"`
	TicketPrice    string `gorm:"type:varchar(20)"`
	TicketQuantity string `gorm:"type:varchar(20)"`
}

type Repository interface {
	GetTickets() ([]Core, error)
	CreateTicket(newTicket Core) (Core, error)
	GetTicket(id uint) (Core, error)
	UpdateTicket(id uint, updatedTicket Core) error
	DeleteTicket(id uint) error
}

type Service interface {
	GetTickets() ([]Core, error)
	CreateTicket(newTicket Core) error
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
