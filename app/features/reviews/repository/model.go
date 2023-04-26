package repository

import "gorm.io/gorm"

type Review struct {
	gorm.Model
	UserID  	uint
	EventID 	uint
	Review  	string  
	Username 	string 
	Image 		string
}
