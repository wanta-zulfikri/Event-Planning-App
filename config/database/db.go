package database

import (
	"fmt"
	"log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDBMySql(cfg config.AppConfig) *gorm.DB {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.DB_USER,
		cfg.DB_PASS,
		cfg.DB_HOST,
		cfg.DB_PORT,
		cfg.DB_NAME,
	)

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		log.Println("error connect to database:", err.Error())
		return nil
	}
	return db
}

// func InitialMigration(db *gorm.DB) {
// 	db.AutoMigrate(&users.Users{})
// 	db.AutoMigrate(&events.Events{})
// 	db.AutoMigrate(&eventsparticipation.Eventsparticipation{})
// 	db.AutoMigrate(&authentication.Authentication{})
// }