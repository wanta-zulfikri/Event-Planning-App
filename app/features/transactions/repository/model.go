package repository

import (
	"time"

	"gorm.io/gorm"
)

// untuk membuat many to many, dibuat relasi antara Transaction dan Ticket : wajib jamak dalam penamaan.

type Transaction struct {
	gorm.Model
	Invoice            string
	PurchaseStartDate  time.Time
	PurchaseEndDate    time.Time
	Status             string
	StatusDate         time.Time
	Subtotal           uint
	GrandTotal         uint
	UserID             uint
	EventID            uint
	TransactionTickets []TransactionTickets
}

type TransactionTickets struct {
	TransactionID  uint
	TicketID       uint
	TicketCategory string
	TicketPrice    string
	Quantity       uint
	Subtotal       uint
	// Ticket         Ticket // >>> detail transaction to get ticketprice, preload ke ticket
}

type Ticket struct {
	gorm.Model
	EventID            uint
	TicketCategory     string
	TicketPrice        uint
	TicketQuantity     uint
	TransactionTickets []TransactionTickets
}
