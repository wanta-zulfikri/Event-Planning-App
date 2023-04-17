package database

import (
	"Event-Planning-App/config"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetConnection(c *config.Config) (*gorm.DB, error) {
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.Database.Username,
		c.Database.Password,
		c.Database.Host,
		c.Database.Port,
		c.Database.Name,
	)

	db, err := gorm.Open(mysql.Open(dataSource), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}

	sql, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("unable to access database: %v", err)
	}

	sql.SetMaxIdleConns(5)
	sql.SetMaxOpenConns(50)
	sql.SetConnMaxIdleTime(5 * time.Minute)
	sql.SetConnMaxLifetime(60 * time.Minute)

	err = sql.Ping()
	if err != nil {
		return nil, fmt.Errorf("unable to establish a good connection to the database: %v", err)
	}

	return db, nil
}
