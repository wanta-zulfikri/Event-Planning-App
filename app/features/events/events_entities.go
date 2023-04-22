package events

import "mime/multipart"

type Core struct {
	Id             uint
	UserID         uint
	Title          string `validate:"required"`
	Details        string `validate:"required"`
	Hosted_by      string `validate:"required"`
	Date           string `validate:"required"`
	Time           string `validate:"required"`
	Status         string
	Event_category string `validate:"required"`
	Location       string
	Image          string
}

type EventService interface {
	Add(newEvent Core, file *multipart.FileHeader) error 
	GetAll(page int, name string) ([]Core, error) 
	Update(userid int, id int, updateEvent Core, file *multipart.FileHeader) error 
	MyEvent(userid int, page int) ([]Core, error) 
	GetEventById(id int) (Core, error) 
	DeleteEvent(userid int, id int) error  
} 

type EventData interface { 
	Insert (input Core) error 
	SelectAll(limit, offset int, name string) ([]Core, error) 
	Update(userid uint, id uint, input Core) error 
	MyEvent(userid int , limit, offset int) ([]Core, error) 
	GetEventById(id uint) (Core, error) 
	DeleteEvent(userid int, id int) error  
}
