package repository

import (
	"github.com/wanta-zulfikri/Event-Planning-App/app/features/tickets/repository"
)

// https://gorm.io/docs/many_to_many.html#Customize-JoinTable
// many2many:transaction_tickets check config migrate, mengharuskan menggunakan automigrate.
type Transaction struct {
	ID      uint `gorm:"primaryKey; autoIncrement"`
	UserID  uint
	EventID uint
	Tickets []repository.Ticket `gorm:"many2many:transaction_tickets;"`
}
