package main

import (
	userHandler "Event-Planning-App/app/features/users/handler"
	userRepo "Event-Planning-App/app/features/users/repository"
	userLogic "Event-Planning-App/app/features/users/service"
	"Event-Planning-App/app/routes"
	"Event-Planning-App/config"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	cfg := config.InitConfiguration()
	db, _ := config.GetConnection(cfg)
	config.Migrate(db)

	userModel := userRepo.New(db)
	userService := userLogic.New(userModel)
	userController := userHandler.New(userService)

	routes.Route(e, userController)

	e.Start(":8080")
}
