package config

import (
	e "github.com/wanta-zulfikri/Event-Planning-App/app/features/events/repository"
	r "github.com/wanta-zulfikri/Event-Planning-App/app/features/reviews/repository"
	t "github.com/wanta-zulfikri/Event-Planning-App/app/features/tickets/repository"
	tr "github.com/wanta-zulfikri/Event-Planning-App/app/features/transactions/repository"
	u "github.com/wanta-zulfikri/Event-Planning-App/app/features/users/repository"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	migrator := db.Migrator()
	if err := migrator.DropTable(
		&u.User{},
		&e.Event{},
		&tr.Transaction{},
		&t.Ticket{},
		&r.Review{}); err != nil {
		return err
	}

	if err := migrator.CreateTable(
		&u.User{},
		&e.Event{},
		&tr.Transaction{},
		&t.Ticket{},
		&r.Review{}); err != nil {
		return err
	}
	return nil
}
