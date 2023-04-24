package reviews

import (
	"github.com/labstack/echo/v4"
)

type Core struct {
	ID      uint
	UserID  uint
	EventID uint
	Review  string
}

type Repository interface {
	WriteReview(newReview Core) (Core, error)
	UpdateReview(id uint, updateReview Core) error
	DeleteReview(id uint) error
}

type Service interface {
	WriteReview(newReview Core) error
	UpdateReview(id uint, updateReview Core) error
	DeleteReview(id uint) error
}

type Handler interface {
	WriteReview() echo.HandlerFunc
	UpdateReview() echo.HandlerFunc
	DeleteReview() echo.HandlerFunc
}
