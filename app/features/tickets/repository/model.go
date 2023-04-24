package repository

import (
	"gorm.io/gorm"
)

type Ticket struct {
	gorm.Model
	TicketType     string `gorm:"type:varchar(20)"`
	TicketCategory string `gorm:"type:varchar(20)"`
	TicketPrice    uint
	TicketQuantity uint
	EventID        uint
	Invoice        string
}
