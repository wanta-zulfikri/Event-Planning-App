package repository

import "gorm.io/gorm"

type Events struct {
	gorm.Model
	Title       string `json:"title" gorm:"type:varchar(100);not null"`
	Description string `json:"description" gorm:"type:varchar(225);not null"`
	EventDate   string `json:"eventdate" gorm:"type:varchar(20);not null"`
	EventTime   string `json:"eventtime" gorm:"type:varchar(20);not null"`
	Status      string `json:"status" gorm:"type:varchar(20);not null"`
	Category    string `json:"category" gorm:"type:varchar(20);not null"`
	Location    string `json:"location" gorm:"type:varchar(100);not null"`
	Image       string `json:"iamge" gorm:"type:varchar(100);not null"`
}
