package repository

import (
	"time"

	"github.com/wanta-zulfikri/Event-Planning-App/app/features/tickets/repository"
	"gorm.io/gorm"
)

// untuk membuat many to many, dibuat relasi antara Transaction dan Ticket : wajib jamak dalam penamaan.
// Ticket - Ticket // >>> detail transaction to get ticketprice, preload ke ticket

type Transaction struct {
	ID                 string `gorm:"primaryKey"`
	PurchaseStartDate  time.Time
	PurchaseEndDate    time.Time
	Status             string
	StatusDate         time.Time
	GrandTotal         uint
	UserID             uint
	EventID            uint
	TransactionTickets []TransactionTickets
}

type TransactionTickets struct {
	TransactionID  uint
	TicketID       repository.Ticket
	TicketCategory string
	TicketPrice    uint
	TicketQuantity uint
	Subtotal       uint
}

type Ticket struct {
	ID                 uint `gorm:"primaryKey"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          gorm.DeletedAt `gorm:"index"`
	EventID            uint
	TicketCategory     string
	TicketPrice        uint
	TicketQuantity     uint
	TransactionTickets []TransactionTickets
}
