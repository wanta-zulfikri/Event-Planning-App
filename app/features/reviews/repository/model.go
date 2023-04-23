package repository

import "gorm.io/gorm"

type Review struct {
	gorm.Model 
	ID      uint     `gorm:"type:varchar(100)"`
	UserID  uint     `gorm:"type:varchar(100)"`
	EventID uint     `gorm:"type:varchar(100)"`
	Review  string   `gorm:"type:varchar(255)"` 
	DeletedAt gorm.DeletedAt `gorm:"index"`
} 


