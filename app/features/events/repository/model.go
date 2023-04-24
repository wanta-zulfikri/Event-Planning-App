package repository

import (
	"github.com/wanta-zulfikri/Event-Planning-App/app/features/tickets/repository"
	"gorm.io/gorm"
)

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
	UserID      uint
	Tickets     []repository.Ticket `gorm:"foreignKey:EventID"`
}
