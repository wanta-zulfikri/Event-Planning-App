package attendances

import (
	"github.com/labstack/echo/v4"
)

type Core struct {
	ID            uint
	UserID        uint
	EventID       uint
	EventCategory string
	TicketType    string
	Quantity      string
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
