package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/wanta-zulfikri/Event-Planning-App/app/features/events"
	"github.com/wanta-zulfikri/Event-Planning-App/helper"
	"github.com/wanta-zulfikri/Event-Planning-App/middlewares"
)

type EventController struct {
	s events.Service
}

func New(h events.Service) events.Handler {
	return &EventController{s: h}
}

func (ec *EventController) CreateEvent() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Unauthorized. Token is missing.", nil))
		}

		_, err := middlewares.ValidateJWT(tokenString)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Unauthorized. "+err.Error(), nil))
		}

		var input RequestCreateEvent
		if err := c.Bind(&input); err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))
		}

		file, err := c.FormFile("image")
		if err != nil {
			c.Logger().Error("Failed to get image: ", err)
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))
		}

		image, err := helper.UploadImage(c, file)
		if err != nil {
			c.Logger().Error("Failed to upload image: ", err)
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal Server Error", nil))
		}

		newEvent := events.Core{
			Title:       input.Title,
			Description: input.Description,
			EventDate:   input.EventDate,
			EventTime:   input.EventTime,
			Status:      input.Status,
			Category:    input.Category,
			Location:    input.Location,
			Image:       image,
		}

		err = ec.s.CreateEvent(newEvent)
		if err != nil {
			c.Logger().Error("Failed to create event: ", err)
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal Server Error", nil))
		}
		return c.JSON(http.StatusCreated, helper.ResponseFormat(http.StatusCreated, "Event created successfully", nil))
	}
}
