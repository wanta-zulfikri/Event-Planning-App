package repository

import "gorm.io/gorm"

type Attendance struct {
	gorm.Model
	ID            uint
	UserID        uint `gorm:"not null"`
	EventID       uint `gorm:"not null;foreignKey:EventID"`
	EventCategory string
	TicketType    string
	Quantity      string
}
