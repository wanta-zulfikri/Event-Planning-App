package repository

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	Invoice           string
	PurchaseStartDate time.Time
	PurchaseEndDate   time.Time
	Status            string
	StatusDate        time.Time
	// Tickets           []*tickets.Ticket `gorm:"many2many:transaction_tickets;"`
	Subtotal   uint
	GrandTotal uint
	UserID     uint `gorm:"foreignKey:ID"`
	EventID    uint
}
