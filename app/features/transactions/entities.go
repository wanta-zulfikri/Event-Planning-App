package transactions

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID                 string
	PurchaseStartDate  time.Time
	PurchaseEndDate    time.Time
	Status             string
	StatusDate         time.Time
	GrandTotal         uint
	UserID             uint
	EventID            uint
	TransactionTickets []TransactionTickets
}

type TransactionTickets struct {
	TransactionID  uint
	TicketID       uint
	TicketCategory string
	TicketPrice    uint
	TicketQuantity uint
	Subtotal       uint
}

type Carts struct {
	ItemDescription []Ticket
}

type Ticket struct {
	TicketID           uint
	EventID            uint
	TicketCategory     string
	TicketPrice        uint
	TicketQuantity     uint
	TransactionTickets []TransactionTickets
}

type Repository interface {
	CreateTransaction(Core) error
}

type Service interface {
	CreateTransaction(userid uint, eventid uint, request Carts) error
}

type Handler interface {
	CreateTransaction() echo.HandlerFunc
}
