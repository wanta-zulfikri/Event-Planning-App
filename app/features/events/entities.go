package events

import (
	"github.com/labstack/echo/v4"
)

type Core struct {
	ID          uint
	Title       string
	Description string
	EventDate   string
	EventTime   string
	Status      string
	Category    string
	Location    string
	Image       string
	Hostedby    string //hostedby : username didapat dari JWT Token
	UserID      uint
	Tickets     []TicketCore     `gorm:"foreignKey:EventID"` 
	Review      []ReviewCore     `gorm:"foreignKey:EventID"` 
	Attendance  []AttendanceCore `gorm:"foreignKey:EventID"`
}

type AttendanceCore struct {
	EventID          uint
	Title            string  
	Descriotion      string 
	HostedBy         string 
	Date             string
    Time             string 
	Status           string 
	Category         string 
	Location         string 
	EventPicture     string 
	
}

type ReviewCore struct {
	ID       uint
	UserID   uint
	Username string
	Image    string
	EventID  uint
	Review   string
}

type TicketCore struct {
	ID             uint
	EventID        uint
	TicketCategory string
	TicketPrice    uint
	TicketQuantity uint
}

type Repository interface {
	CreateEventWithTickets(event Core, userID uint) error
	GetEvents() ([]Core, error)
	GetEventsByCategory(category string) ([]Core, error)
	GetEventsByUserID(userid uint) ([]Core, error)
	GetEvent(eventid uint) (Core, error)
	UpdateEvent(id uint, updatedEvent Core) error
	DeleteEvent(id uint) error
}

type Service interface {
	CreateEventWithTickets(event Core, userID uint) error
	GetEvents() ([]Core, error)
	GetEventsByCategory(category string) ([]Core, error)
	GetEventsByUserID(userid uint) ([]Core, error)
	GetEvent(eventid uint) (Core, error)
	UpdateEvent(id uint, updatedEvent Core) error
	DeleteEvent(id uint) error
}

type Handler interface {
	CreateEventWithTickets() echo.HandlerFunc
	GetEvents() echo.HandlerFunc
	GetEventsByUserID() echo.HandlerFunc
	GetEvent() echo.HandlerFunc
	UpdateEvent() echo.HandlerFunc
	DeleteEvent() echo.HandlerFunc
}
