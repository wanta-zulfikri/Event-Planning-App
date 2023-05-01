package transactions

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Transaction struct {
	ID                  uint
	UserID              uint
	EventID             uint
	Invoice             string
	Username            string
	Email               string
	Attendee            string
	AEmail              string
	Title               string
	EventDate           string
	EventTime           string
	PurchaseStartDate   time.Time
	PurchaseEndDate     time.Time
	Status              string
	StatusDate          time.Time
	GrandTotal          uint
	PaymentMethod       string
	Transaction_Tickets []Transaction_Tickets
}

type Transaction_Tickets struct {
	TransactionID  uint
	TicketID       uint
	TicketCategory string
	TicketPrice    uint
	TicketQuantity uint
	Subtotal       uint
}

type Repository interface {
	CreateTransaction(Transaction) error
	GetTransaction(transactionid uint) (Transaction, error)
}

type Service interface {
	CreateTransaction(user_id, event_id, grandtotal uint, paymentmethod string, request Transaction) error
	GetTransaction(transactionid uint) (Transaction, error)
}

type Handler interface {
	CreateTransaction() echo.HandlerFunc
	GetTransaction() echo.HandlerFunc
}
