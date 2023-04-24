package repository

import (
	reviews "github.com/wanta-zulfikri/Event-Planning-App/app/features/reviews/repository"
	tickets "github.com/wanta-zulfikri/Event-Planning-App/app/features/tickets/repository"
	transactions "github.com/wanta-zulfikri/Event-Planning-App/app/features/transactions/repository"
	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Title        string                     `gorm:"type:varchar(100)"`
	Description  string                     `gorm:"type:varchar(225)"`
	EventDate    string                     `gorm:"type:varchar(20)"`
	EventTime    string                     `gorm:"type:varchar(20)"`
	Status       string                     `gorm:"type:varchar(20)"`
	Category     string                     `gorm:"type:varchar(20)"`
	Location     string                     `gorm:"type:varchar(100)"`
	Image        string                     `gorm:"type:varchar(100)"`
	Username     string                     `gorm:"type:varchar(100)"`
	UserID       uint                       `gorm:"onUpdate:CASCADE,onDelete:SET NULL"`
	Tickets      []tickets.Ticket           `gorm:"foreignKey:EventID"`
	Transactions []transactions.Transaction `gorm:"foreignKey:EventID"`
	Reviews      []reviews.Review           `gorm:"foreignKey:EventID"`
}
