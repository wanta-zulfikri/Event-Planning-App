package repository

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(255);not null"`
	Email    string `gorm:"primaryKey;type:varchar(255);unique"`
	Password string `gorm:"type:varchar(100);not null"`
}
