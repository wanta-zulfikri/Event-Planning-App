package transactions

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Transaction struct {
	Invoice            string
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

type Repository interface {
	CreateTransaction(userID uint, eventID uint, transaction TransactionTickets) error
}

type Service interface {
	CreateTransaction(userID uint, eventID uint, transaction TransactionTickets) error
}

type Handler interface {
	CreateTransaction() echo.HandlerFunc
}
