package config

import (
	"github.com/wanta-zulfikri/Event-Planning-App/app/features/events"
	"github.com/wanta-zulfikri/Event-Planning-App/app/features/users/repository"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	migrator := db.Migrator()
	migrator.CreateTable(&repository.User{}, &events.Events{})
}
