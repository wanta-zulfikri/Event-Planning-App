package repository

import "gorm.io/gorm"

type Event struct {
	gorm.Model 
	UserID uint 
	Title string 
	
}