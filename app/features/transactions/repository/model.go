package repository

import (
	"time"

	"github.com/wanta-zulfikri/Event-Planning-App/app/features/tickets/repository"
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	Invoice           string
	PurchaseStartDate time.Time
	PurchaseEndDate   time.Time
	Status            string
	StatusDate        time.Time
	Tickets           []*repository.Ticket `gorm:"many2many:transaction_tickets;"`
	Subtotal          uint
	GrandTotal        uint
	UserID            uint `gorm:"foreignKey:ID"`
	EventID           uint
}
