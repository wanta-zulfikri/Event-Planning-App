package main

import (
	"fmt" 
    
	attendancesHandler "github.com/wanta-zulfikri/Event-Planning-App/app/features/attendances/handler"
	attendancesRepo "github.com/wanta-zulfikri/Event-Planning-App/app/features/attendances/repository"
	attendancesLogic "github.com/wanta-zulfikri/Event-Planning-App/app/features/attendances/services"
	reviewHandler "github.com/wanta-zulfikri/Event-Planning-App/app/features/reviews/handler"
	reviewRepo "github.com/wanta-zulfikri/Event-Planning-App/app/features/reviews/repository"
	reviewLogic "github.com/wanta-zulfikri/Event-Planning-App/app/features/reviews/services"
	eventHandler "github.com/wanta-zulfikri/Event-Planning-App/app/features/events/handler"
	eventRepo "github.com/wanta-zulfikri/Event-Planning-App/app/features/events/repository"
	eventLogic "github.com/wanta-zulfikri/Event-Planning-App/app/features/events/services"
	ticketHandler "github.com/wanta-zulfikri/Event-Planning-App/app/features/tickets/handler"
	ticketRepo "github.com/wanta-zulfikri/Event-Planning-App/app/features/tickets/repository"
	ticketLogic "github.com/wanta-zulfikri/Event-Planning-App/app/features/tickets/services"
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
	ticketModel := ticketRepo.New(db)
	ticketService := ticketLogic.New(ticketModel)
	ticketController := ticketHandler.New(ticketService) 
	reviewModel := reviewRepo.New(db) 
	reviewService := reviewLogic.New(reviewModel) 
	reviewController := reviewHandler.New(reviewService) 
	attendancesModel := attendancesRepo.New(db) 
	attendancesService := attendancesLogic.New(attendancesModel) 
	attendancesController := attendancesHandler.New(attendancesService)

	routes.Route(e, userController, eventController, ticketController, reviewController, attendancesController)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", cfg.Port)))
}
