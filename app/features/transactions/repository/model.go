package repository

import (
	"time"

	"github.com/wanta-zulfikri/Event-Planning-App/app/features/events"
	"github.com/wanta-zulfikri/Event-Planning-App/app/features/users"
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	Invoice           string
	PurchaseStartDate time.Time
	PurchaseEndDate   time.Time
	Status            string
	StatusDate        time.Time
	Tickets           []Ticket `gorm:"polymorphic:Owner;"`
	Subtotal          uint
	GrandTotal        uint
	PaymentStatus     string      `json:"payment_status" gorm:"size:20"`
	PaymentType       string      `json:"payment_type" gorm:"size:50"`
	PaymentLink       string      `json:"payment_link" gorm:"size:255"`
	UserEmail         users.Core  `gorm:"foreignkey:UserEmail"`
	Event             events.Core `gorm:"foreignkey:EventID"`
}

type Ticket struct {
	gorm.Model
	TicketCategory  string
	TicketAvailable uint
	TicketQuantity  uint
	TicketPrice     uint
	Event           events.Core `gorm:"foreignkey:EventID"`
}

type TransactionDetail struct {
	ID      string       `json:"id" gorm:"primaryKey;size:255"`
	OrderId string       `json:"order_id" gorm:"size:255"`
	EventId string       `json:"event_id" gorm:"size:255"`
	Event   *events.Core `json:"event,omitempty"`
}
