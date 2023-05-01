package transactions

import (
	"github.com/labstack/echo/v4"
)

type Transaction struct {
	ID                  uint
	UserID              uint
	EventID             uint
	Transaction_Tickets []Transaction_Tickets
}

type Transaction_Tickets struct {
	TransactionID uint
	TicketID      uint
}

type Repository interface {
	CreateTransaction(Transaction) error
}

type Service interface {
	CreateTransaction(userid uint, eventid uint, request Transaction) error
}

type Handler interface {
	CreateTransaction() echo.HandlerFunc
}

type RequestCreateTransaction struct {
	ItemDescription []Tickets `json:"item_description"`
}

type Tickets struct {
	TicketID uint `json:"ticket_id"`
}
