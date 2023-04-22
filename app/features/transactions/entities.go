package transactions

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/wanta-zulfikri/Event-Planning-App/app/features/events"
	"github.com/wanta-zulfikri/Event-Planning-App/app/features/tickets"
	"github.com/wanta-zulfikri/Event-Planning-App/app/features/transactions/payment"
	"github.com/wanta-zulfikri/Event-Planning-App/app/features/users"
)

type Core struct {
	Invoice           string
	PurchaseStartDate time.Time
	PurchaseDueDate   time.Time
	Status            string
	StatusDate        time.Time
	Tickets           []tickets.Core
	Subtotal          uint
	GrandTotal        uint
	PaymentStatus     string
	PaymentType       string
	PaymentLink       string
	UserID            users.Core
	EventID           events.Core
}

type TransactionDetail struct {
	ID      string
	OrderId string
	EventId string
	Event   *events.Core
}

type PaymentGatewayInterface interface {
	InitializeClientMidtrans()
	CreateTransaction(snap payment.PaymentGateway) string
	CreateUrlTransactionWithGateway(snap payment.PaymentGateway) string
}

type Repository interface {
	CreateTransaction(newTransaction Core) error
}

type Service interface {
	CreateTransaction(newTransaction Core, userId int) (map[string]interface{}, error)
}

type Handler interface {
	CreateTransaction() echo.HandlerFunc
}
