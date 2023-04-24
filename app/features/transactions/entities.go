package transactions

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/wanta-zulfikri/Event-Planning-App/app/features/transactions/payment"
	"gorm.io/gorm"
)

type Core struct {
	ID                uint
	Invoice           string
	PurchaseStartDate time.Time
	PurchaseEndDate   time.Time
	Status            string
	StatusDate        time.Time
	Tickets           []Ticket
	Subtotal          uint
	GrandTotal        uint
	UserID            uint
	EventID           uint
}

type Transaction struct {
	gorm.Model
	Invoice           string
	PurchaseStartDate time.Time
	PurchaseEndDate   time.Time
	Status            string
	StatusDate        time.Time
	Tickets           []Ticket `gorm:"foreignKey:TransactionID"`
	Subtotal          uint
	GrandTotal        uint
	UserID            uint `gorm:"foreignKey:Email"`
	EventID           uint
}

type Ticket struct {
	gorm.Model
	TicketType     string `gorm:"type:varchar(20)"`
	TicketCategory string `gorm:"type:varchar(20)"`
	TicketPrice    uint
	TicketQuantity uint
	EventID        uint
	TransactionID  uint
}

type PaymentGatewayInterface interface {
	InitializeClientMidtrans()
	CreateTransaction(snap payment.PaymentGateway) string
	CreateUrlTransactionWithGateway(snap payment.PaymentGateway) string
}

type Repository interface {
	CreateTransaction(newTransaction Core) (Core, error)
}

type Service interface {
	CreateTransaction(userID uint, inputs []Core) error
}

type Handler interface {
	CreateTransaction() echo.HandlerFunc
}
