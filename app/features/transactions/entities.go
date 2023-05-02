package transactions

import (
	"github.com/labstack/echo/v4"
)

// type Transaction struct {
// 	ID                  uint
// 	UserID              uint
// 	EventID             uint
// 	Invoice             string
// 	Username            string
// 	Email               string
// 	Attendee            string
// 	AEmail              string
// 	Title               string
// 	EventDate           string
// 	EventTime           string
// 	PurchaseStartDate   time.Time
// 	PurchaseEndDate     time.Time
// 	Status              string
// 	StatusDate          time.Time
// 	GrandTotal          uint
// 	PaymentMethod       string
// 	PaymentType         string
// 	Transaction_Tickets []Transaction_Tickets
// }

type Transaction_Tickets struct {
	TransactionID  uint
	TicketID       string
	TicketCategory string
	TicketPrice    int64
	TicketQuantity uint
	Subtotal       uint
}

type Transaction struct {
	UserID              uint
	EventID             uint
	PaymentType         string
	Transaction_Details []TransactionDetails
	Transaction_Tickets []Transaction_Tickets
	BankTransfer        *BankTransferDetails
}

type TransactionDetails struct {
	OrderID     string
	GrossAmount uint
}

type BankTransferDetails struct {
	Bank string `json:"bank"`
}

type Repository interface {
	// CreateTransaction(Transaction) error
	// GetTransaction(transactionid uint) (Transaction, error)
	PayTransaction(user_id uint, event_id uint, paymenttype string, request Transaction) (string, error)
}

type Service interface {
	// CreateTransaction(user_id, event_id, grandtotal uint, paymentmethod string, request Transaction) error
	// GetTransaction(transactionid uint) (Transaction, error)
	PayTransaction(user_id uint, event_id uint, paymenttype string, request Transaction) (string, error)
}

type Handler interface {
	// CreateTransaction() echo.HandlerFunc
	// GetTransaction() echo.HandlerFunc
	PayTransaction() echo.HandlerFunc
}
