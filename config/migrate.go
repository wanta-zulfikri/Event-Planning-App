package config

import (
	"github.com/wanta-zulfikri/Event-Planning-App/app/features/attendances"
	"github.com/wanta-zulfikri/Event-Planning-App/app/features/events"
	"github.com/wanta-zulfikri/Event-Planning-App/app/features/tickets"
	"github.com/wanta-zulfikri/Event-Planning-App/app/features/transactions"
	"github.com/wanta-zulfikri/Event-Planning-App/app/features/users"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	migrator := db.Migrator()

	migrator.DropTable(&users.User{})
	migrator.DropTable(&events.Event{})
	migrator.DropTable(&attendances.Attendances{})
	migrator.DropTable(&tickets.Ticket{})
	migrator.DropTable(&transactions.Transaction{})

	migrator.CreateTable(&users.User{})
	migrator.CreateTable(&events.Event{})
	migrator.CreateTable(&attendances.Attendances{})
	migrator.CreateTable(&tickets.Ticket{})
	migrator.CreateTable(&transactions.Transaction{})
}
