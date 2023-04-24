package repository

import "gorm.io/gorm"

type Review struct {
	gorm.Model
	UserID  uint
	EventID uint
	Review  string `gorm:"type:varchar(255)"`
}
