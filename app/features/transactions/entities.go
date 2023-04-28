package transactions

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID                uint
	Invoice           string
	PurchaseStartDate time.Time
	PurchaseEndDate   time.Time
	Status            string
	StatusDate        time.Time
	Tickets           []TicketCore
	Subtotal          uint
	GrandTotal        uint
	UserID            uint
	EventID           uint
}

type TransactionTicketCore struct {
	TransactionID uint
	TicketID      uint
	Quantity      uint
}

type TicketCore struct {
	ID                 uint
	EventID            uint
	TicketCategory     string
	TicketPrice        uint
	TicketQuantity     uint
	TransactionTickets []TransactionTicketCore
}

type Repository interface {
	CreateTransaction(userid uint, input Core) error
}

type Service interface {
	CreateTransaction(userid uint, input Core) error
}

type Handler interface {
	CreateTransaction() echo.HandlerFunc
}
