package repository

import (
	"gorm.io/gorm"
)

// type Ticket struct {
// 	gorm.Model
// 	Title          string
// 	TicketType     string `gorm:"type:varchar(20)"`
// 	TicketCategory string `gorm:"type:varchar(20)"`
// 	TicketPrice    uint
// 	TicketQuantity uint
// 	EventID        uint
// }

type Ticket struct {
	gorm.Model
	EventID        uint `gorm:"references:EventID"`
	Title          string
	TicketType     string `gorm:"type:varchar(20)"`
	TicketCategory string `gorm:"type:varchar(20)"`
	TicketPrice    uint
	TicketQuantity uint
}
