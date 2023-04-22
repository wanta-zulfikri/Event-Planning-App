package config

import "gorm.io/gorm"

func Migrate(db *gorm.DB) {
	migrator := db.Migrator()
	migrator.CreateTable(&repository.User{})
	migrator.CreateTable(&repository.Events{})
}
