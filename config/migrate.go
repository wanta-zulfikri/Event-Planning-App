package config

import (
	events "github.com/wanta-zulfikri/Event-Planning-App/app/features/events/repository"
	reviews "github.com/wanta-zulfikri/Event-Planning-App/app/features/reviews/repository"
	tickets "github.com/wanta-zulfikri/Event-Planning-App/app/features/tickets/repository"
	transactions "github.com/wanta-zulfikri/Event-Planning-App/app/features/transactions/repository"
	users "github.com/wanta-zulfikri/Event-Planning-App/app/features/users/repository"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&users.User{},
		&events.Event{},
		&transactions.Transaction{},
		&transactions.Transaction_Tickets{},
		&transactions.Payment{},
		&tickets.Ticket{},
		&reviews.Review{},
	)
	return err
}
