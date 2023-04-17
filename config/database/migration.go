package database

import "Event-Planning-App/config"

func Migrate(c *config.Config) {
	db, err := GetConnection(c)
	if err != nil {
		panic(err)
	}
	if err := db.AutoMigrate(); err != nil {
		panic(err)
	}
}
