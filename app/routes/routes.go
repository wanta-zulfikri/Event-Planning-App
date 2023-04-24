package routes

import (
	"github.com/wanta-zulfikri/Event-Planning-App/app/features/attendances"
	"github.com/wanta-zulfikri/Event-Planning-App/app/features/events"
	"github.com/wanta-zulfikri/Event-Planning-App/app/features/reviews"
	"github.com/wanta-zulfikri/Event-Planning-App/app/features/tickets"
	"github.com/wanta-zulfikri/Event-Planning-App/app/features/users"
	"github.com/wanta-zulfikri/Event-Planning-App/config/common"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Route(e *echo.Echo, uc users.Handler, ec events.Handler, rc reviews.Handler, tc tickets.Handler, as attendances.Handler) {
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.Logger()) 
	//authentication
	e.POST("/register", uc.Register())
	e.POST("/login", uc.Login())
	//users
	e.GET("/users", uc.GetProfile(), middleware.JWT([]byte(common.JWTSecret)))
	e.PUT("/users", uc.UpdateProfile(), middleware.JWT([]byte(common.JWTSecret)))
	e.DELETE("/users", uc.DeleteProfile(), middleware.JWT([]byte(common.JWTSecret)))
	//events
	e.GET("/events", ec.GetEvents(), middleware.JWT([]byte(common.JWTSecret)))
	e.POST("/events", ec.CreateEvent(), middleware.JWT([]byte(common.JWTSecret)))
	e.GET("/events/:id", ec.GetEvent(), middleware.JWT([]byte(common.JWTSecret)))
	e.PUT("/events/:id", ec.UpdateEvent(), middleware.JWT([]byte(common.JWTSecret)))
	e.DELETE("/events/:id", ec.DeleteEvent(), middleware.JWT([]byte(common.JWTSecret)))
	//tickets
	e.GET("/tickets", tc.GetTickets(), middleware.JWT([]byte(common.JWTSecret)))
	e.POST("/tickets", tc.CreateTicket(), middleware.JWT([]byte(common.JWTSecret)))
	e.GET("/tickets/:id", tc.GetTicket(), middleware.JWT([]byte(common.JWTSecret)))
	e.PUT("/tickets/:id", tc.UpdateTicket(), middleware.JWT([]byte(common.JWTSecret)))
	e.DELETE("/tickets/:id", tc.DeleteTicket(), middleware.JWT([]byte(common.JWTSecret)))
	//attendancees
	e.POST("/attendances", as.CreateAttendance(), middleware.JWT([]byte(common.JWTSecret))) 
	e.GET("/attendances", as.GetAttendance(), middleware.JWT([]byte(common.JWTSecret)))
	//transactions

	//reviews 
	e.POST("/reviews", rc.UpdateReview(), middleware.JWT([]byte(common.JWTSecret))) 
	e.PUT("/reviews", rc.UpdateReview(), middleware.JWT([]byte(common.JWTSecret))) 
	e.DELETE("/reviews", rc.DeleteReview(), middleware.JWT([]byte(common.JWTSecret)))
}
