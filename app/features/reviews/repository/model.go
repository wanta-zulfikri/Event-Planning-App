package repository

import "gorm.io/gorm"

type Review struct {
	gorm.Model
	UserID  uint   `gorm:"not null"`
	EventID uint   `gorm:"not null;foreignKey:EventID"`
	Review  string `gorm:"not null"`
}
