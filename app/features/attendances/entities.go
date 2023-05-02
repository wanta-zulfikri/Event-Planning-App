package attendances

import (
	"github.com/labstack/echo/v4"
)

type Core struct { 
	EventID        uint 
	Title 			string  
	Description 	string 
	HostedBy  		string 
	Date 			string 
	Time 			string  
	Status 			string  
	Category 		string  
	Location		string  
	EventPicture 	string
}

type Repository interface {
	CreateAttendance(newAttendance Core) (Core, error)
	GetAttendance() ([]Core, error)
}

type Service interface {
	CreateAttendance(newAttendance Core) error
	GetAttendance() ([]Core, error)
}

type Handler interface {
	CreateAttendance() echo.HandlerFunc
	GetAttendance() echo.HandlerFunc
}
