package repository

import "gorm.io/gorm"

type Event struct {
	gorm.Model
	Title       string `gorm:"type:varchar(100)"`
	Description string `gorm:"type:varchar(225)"`
	EventDate   string `gorm:"type:varchar(20)"`
	EventTime   string `gorm:"type:varchar(20)"`
	Status      string `gorm:"type:varchar(20)"`
	Category    string `gorm:"type:varchar(20)"`
	Location    string `gorm:"type:varchar(100)"`
	Image       string `gorm:"type:varchar(100)"`
	Hostedby    string `gorm:"type:varchar(100)"`
	UserID      uint
	Tickets     []Ticket `gorm:"foreignKey:EventID"`
}

type Ticket struct {
	gorm.Model
	EventID        uint `gorm:"references:EventID"`
	Title          string
	TicketType     string `gorm:"type:varchar(20)"`
	TicketCategory string `gorm:"type:varchar(20)"`
	TicketPrice    uint
	TicketQuantity uint
}
