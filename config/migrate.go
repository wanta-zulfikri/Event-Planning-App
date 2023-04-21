package config

import (
	"github.com/wanta-zulfikri/Event-Planning-App/app/features/users/repository"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(repository.User{})
}
