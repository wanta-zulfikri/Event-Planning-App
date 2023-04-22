package repository

import "gorm.io/gorm"

type Ticket struct {
	gorm.Model
	UserID  uint `gorm:"not null"`
	EventID uint `gorm:"not null;foreignKey:EventID"`
}
