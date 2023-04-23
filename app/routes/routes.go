package routes

import (
	"github.com/wanta-zulfikri/Event-Planning-App/app/features/events"
	"github.com/wanta-zulfikri/Event-Planning-App/app/features/users"
	"github.com/wanta-zulfikri/Event-Planning-App/config/common"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	usersPath  = "/users"
	eventsPath = "/events"
)

var jwtSecret = []byte(common.JWTSecret)

func Route(e *echo.Echo, uc users.Handler, ec events.Handler) {
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	authMiddleware := middleware.JWT(jwtSecret)

	authRoutes := e.Group("")
	authRoutes.POST("/register", uc.Register())
	authRoutes.POST("/login", uc.Login())

	usersRoutes := e.Group(usersPath, authMiddleware)
	usersRoutes.GET("", uc.GetProfile())
	usersRoutes.PUT("", uc.UpdateProfile())
	usersRoutes.DELETE("", uc.DeleteProfile())

	eventsRoutes := e.Group(eventsPath, authMiddleware)
	eventsRoutes.GET("", ec.GetEvents())
	eventsRoutes.POST("", ec.CreateEvent())
	eventsRoutes.GET("/:id", ec.GetEvent())
	eventsRoutes.PUT("/:id", ec.UpdateEvent())
	eventsRoutes.DELETE("/:id", ec.DeleteEvent())

	//tickets

	//attendances

	//transactions

	//reviews
}
