package routes

import (
	"github.com/wanta-zulfikri/Event-Planning-App/app/features/events"
	"github.com/wanta-zulfikri/Event-Planning-App/app/features/reviews"
	"github.com/wanta-zulfikri/Event-Planning-App/app/features/tickets"
	"github.com/wanta-zulfikri/Event-Planning-App/app/features/transactions"
	"github.com/wanta-zulfikri/Event-Planning-App/app/features/users"
	"github.com/wanta-zulfikri/Event-Planning-App/config/common"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Route(e *echo.Echo, uc users.Handler, ec events.Handler, tc tickets.Handler, tr transactions.Handler, rc reviews.Handler) {
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
	e.GET("/users/events", ec.GetEventsByUserID(), middleware.JWT([]byte(common.JWTSecret)))
	//events
	e.GET("/events", ec.GetEvents())
	e.POST("/events", ec.CreateEventWithTickets(), middleware.JWT([]byte(common.JWTSecret)))
	e.GET("/events/:id", ec.GetEvent())
	e.PUT("/events/:id", ec.UpdateEvent(), middleware.JWT([]byte(common.JWTSecret)))
	e.DELETE("/events/:id", ec.DeleteEvent(), middleware.JWT([]byte(common.JWTSecret)))
	//tickets
	e.GET("/tickets/:id", tc.GetTickets(), middleware.JWT([]byte(common.JWTSecret)))
	e.PUT("/tickets/:id", tc.UpdateTicket(), middleware.JWT([]byte(common.JWTSecret)))
	e.DELETE("/tickets/:id", tc.DeleteTicket(), middleware.JWT([]byte(common.JWTSecret)))
	//transactions
	e.POST("/transactions/:id", tr.CreateTransaction(), middleware.JWT([]byte(common.JWTSecret)))
	//reviews
	e.POST("/reviews", rc.WriteReview(), middleware.JWT([]byte(common.JWTSecret)))
	e.PUT("/reviews", rc.UpdateReview(), middleware.JWT([]byte(common.JWTSecret)))
	e.DELETE("/reviews", rc.DeleteReview(), middleware.JWT([]byte(common.JWTSecret)))
}
