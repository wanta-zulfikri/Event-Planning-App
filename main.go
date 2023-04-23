package main

import (
	"fmt" 
    
	reviewHandler "github.com/wanta-zulfikri/Event-Planning-App/app/features/review/handler"
	reviewRepo "github.com/wanta-zulfikri/Event-Planning-App/app/features/review/repository"
	reviewLogic "github.com/wanta-zulfikri/Event-Planning-App/app/features/review/service"
	eventHandler "github.com/wanta-zulfikri/Event-Planning-App/app/features/events/handler"
	eventRepo "github.com/wanta-zulfikri/Event-Planning-App/app/features/events/repository"
	eventLogic "github.com/wanta-zulfikri/Event-Planning-App/app/features/events/services"
	userHandler "github.com/wanta-zulfikri/Event-Planning-App/app/features/users/handler"
	userRepo "github.com/wanta-zulfikri/Event-Planning-App/app/features/users/repository"
	userLogic "github.com/wanta-zulfikri/Event-Planning-App/app/features/users/service"
	"github.com/wanta-zulfikri/Event-Planning-App/app/routes"
	"github.com/wanta-zulfikri/Event-Planning-App/config"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	cfg := config.GetConfiguration()
	db, _ := config.GetConnection(*cfg)
	config.Migrate(db)

	userModel := userRepo.New(db)
	userService := userLogic.New(userModel)
	userController := userHandler.New(userService)
	eventModel := eventRepo.New(db)
	eventService := eventLogic.New(eventModel)
	eventController := eventHandler.New(eventService)

	routes.Route(e, userController, eventController)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", cfg.Port)))
}
