package events

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Core struct {
	Id          uint
	UserID      uint
	Title       string `validate:"required"`
	Description string `validate:"required"`
	Hosted_by   string `validate:"required"`
	EventDate   string `validate:"required"`
	EventTime   string `validate:"required"`
	Status      string
	Category    string `validate:"required"`
	Location    string
	Image       string
}

type Event struct {
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

type Repository interface {
	CreateEvent(newEvent Core) (Core, error)
}

type Service interface {
	CreateEvent(newEvent Core) error
}

type Handler interface {
	CreateEvent() echo.HandlerFunc
}
