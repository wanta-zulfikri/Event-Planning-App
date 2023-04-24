package transactions

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/wanta-zulfikri/Event-Planning-App/app/features/tickets"
	"github.com/wanta-zulfikri/Event-Planning-App/app/features/transactions/payment"
)

type Core struct {
	ID                uint
	Invoice           string
	PurchaseStartDate time.Time
	PurchaseEndDate   time.Time
	Status            string
	StatusDate        time.Time
	Tickets           []tickets.Core
	Subtotal          uint
	GrandTotal        uint
	UserID            uint
	EventID           uint
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
