package config

import (
	a "github.com/wanta-zulfikri/Event-Planning-App/app/features/attendances/repository"
	e "github.com/wanta-zulfikri/Event-Planning-App/app/features/events/repository"
	r "github.com/wanta-zulfikri/Event-Planning-App/app/features/reviews/repository"
	t "github.com/wanta-zulfikri/Event-Planning-App/app/features/tickets/repository"
	tr "github.com/wanta-zulfikri/Event-Planning-App/app/features/transactions/repository"
	u "github.com/wanta-zulfikri/Event-Planning-App/app/features/users/repository"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	migrator := db.Migrator()

	migrator.DropTable(&u.User{})
	migrator.DropTable(&e.Event{})
	migrator.DropTable(&tr.Transaction{})
	migrator.DropTable(&tr.TransactionTickets{})
	migrator.DropTable(&t.Ticket{})
	migrator.DropTable(&a.Attendance{})
	migrator.DropTable(&r.Review{})

	migrator.CreateTable(&u.User{})
	migrator.CreateTable(&e.Event{})
	migrator.CreateTable(&tr.Transaction{})
	migrator.CreateTable(&tr.TransactionTickets{})
	migrator.CreateTable(&t.Ticket{})
	migrator.CreateTable(&a.Attendance{})
	migrator.CreateTable(&r.Review{})
}
