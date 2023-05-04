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

type TrasactionDetails struct {
	Order_ID     string
	Gross_Amount uint64
}

type Payment struct {
	Transaction_ID     string
	Order_ID           string
	Gross_Amount       string
	Payment_Type       string
	Bank               string
	Transaction_Time   string
	Transaction_Status string
	Va_Numbers         string
}

type Repository interface {
	CreateTransaction(Transaction) (Transaction, error)
	GetTransaction(invoice string) (Transaction, error)
	Payment(invoice string, grossAmount uint) (Payment, error)
}

type Service interface {
	CreateTransaction(user_id, event_id, grandtotal uint, paymentmethod string, request Transaction) (Transaction, error)
	GetTransaction(invoice string) (Transaction, error)
	Payment(invoive string, grossAmount uint) (Payment, error)
}

type Handler interface {
	CreateTransaction() echo.HandlerFunc
	GetTransaction() echo.HandlerFunc
	Payment() echo.HandlerFunc
}
