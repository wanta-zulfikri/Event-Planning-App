package repository

import (
	"time"

	"github.com/wanta-zulfikri/Event-Planning-App/app/features/tickets/repository"
)

// https://gorm.io/docs/many_to_many.html#Customize-JoinTable
// many2many:transaction_tickets check config migrate, wajib menggunakan automigrate.
type Transaction struct {
	ID                 uint `gorm:"primaryKey; autoIncrement"`
	UserID             uint
	EventID            uint
	Invoice            string
	PurchaseStartDate  time.Time
	PurchaseEndDate    time.Time
	Status             string
	StatusDate         time.Time
	GrandTotal         uint
	PaymentMethod      string
	TransactionTickets []Transaction_Tickets `gorm:"many2many:transaction_tickets;"`
}

type Transaction_Tickets struct {
	ID             uint
	TransactionID  uint              //**
	Transaction    Transaction       //**
	TicketID       uint              //**
	Ticket         repository.Ticket //**
	TicketCategory string
	TicketPrice    uint
	TicketQuantity uint
	Subtotal       uint
}
