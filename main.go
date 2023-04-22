package main

import (
	"fmt"

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

	routes.Route(e, userController)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", cfg.Port)))
}
