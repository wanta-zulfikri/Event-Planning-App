package repository

import (
	"time"

	tickets "github.com/wanta-zulfikri/Event-Planning-App/app/features/tickets/repository"
)

type Transaction struct {
	ID                uint `gorm:"primaryKey"`
	PurchaseStartDate time.Time
	PurchaseEndDate   time.Time
	Status            string
	StatusDate        time.Time
	GrandTotal        uint
	UserID            uint
	EventID           uint
	Tickets           []tickets.Ticket `gorm:"many2many:transactiontickets;"`
}

type TransactionTicket struct {
	TransactionID  uint   `gorm:"primaryKey"`
	TicketID       uint   `gorm:"primaryKey"`
	TicketCategory string `gorm:"type:varchar(20)"`
	TicketPrice    uint
	TicketQuantity uint
	Subtotal       uint
}
