package repository

import "gorm.io/gorm"

type Attendance struct {
	gorm.Model 
	ID              uint
	EventID         uint 
	Title          string 
	Description    string  
	HostedBy       string 
	Date           string 
	Time           string  
	Status         string  
	Category       string   
	Location       string      
	EventPicture   string
}
