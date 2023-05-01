package repository

import (
	"gorm.io/gorm"
)

type Ticket struct {
	gorm.Model
	EventID        uint   `gorm:"references:ID"`
	TicketCategory string `gorm:"type:varchar(20)"`
	TicketPrice    uint
	TicketQuantity uint
}
