package handler

import (
	"fmt"
	"net/http"
	"strconv"

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

func (ec *EventController) CreateEventWithTickets() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input RequestCreateEventWithTickets
		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Unauthorized. Token is missing.", nil))
		}
		idString, err := middlewares.ValidateJWT(tokenString)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Unauthorized. "+err.Error(), nil))
		}

		id, err := strconv.ParseUint(fmt.Sprint(idString), 10, 64)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Unauthorized. Invalid token.", nil))
		}

		username, err := middlewares.ValidateJWTUsername(tokenString)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Unauthorized. "+err.Error(), nil))
		}

		if username == "" {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Unauthorized. Invalid token.", nil))
		}

		if err := c.Bind(&input); err != nil {
			c.Logger().Error("Failed to bind input: ", err)
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))
		}

		// file, err := c.FormFile("image")
		// if err != nil {
		// 	c.Logger().Error("Failed to get image: ", err)
		// 	return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))
		// }

		// image, err := helper.UploadImage(c, file)
		// if err != nil {
		// 	c.Logger().Error("Failed to upload image: ", err)
		// 	return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal Server Error", nil))
		// }

		eventTickets := make([]events.TicketCore, len(input.Tickets))
		for i, ticket := range input.Tickets {
			eventTickets[i] = events.TicketCore{
				Title:          ticket.Title,
				TicketType:     ticket.TicketType,
				TicketCategory: ticket.TicketCategory,
				TicketPrice:    ticket.TicketPrice,
				TicketQuantity: ticket.TicketQuantity,
			}
		}

		newEvent := events.Core{
			Title:       input.Title,
			Description: input.Description,
			EventDate:   input.EventDate,
			EventTime:   input.EventTime,
			Status:      input.Status,
			Category:    input.Category,
			Location:    input.Location,
			Image:       input.Image,
			Hostedby:    username,
			Tickets:     eventTickets,
		}

		err = ec.s.CreateEventWithTickets(newEvent, uint(id))
		if err != nil {
			c.Logger().Error("Failed to create event with tickets: ", err)
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal Server Error", nil))
		}

		return c.JSON(http.StatusCreated, helper.ResponseFormat(http.StatusCreated, "Event with tickets created successfully", nil))
	}
}
