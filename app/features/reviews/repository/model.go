package repository

import (
	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	UserID   uint
	Username string
	EventID  uint
	Review   string `gorm:"type:varchar(255)"`
}
