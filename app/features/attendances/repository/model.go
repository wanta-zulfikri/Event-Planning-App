package repository

import "gorm.io/gorm"

type Attendance struct {
	gorm.Model
	ID            uint
	UserID        uint 
	EventID       uint 
	EventCategory string
	TicketType    string
	Quantity      string
}
