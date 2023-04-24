package handler

import (
	"math"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/wanta-zulfikri/Event-Planning-App/app/features/tickets"
	"github.com/wanta-zulfikri/Event-Planning-App/helper"
	"github.com/wanta-zulfikri/Event-Planning-App/middlewares"
)

type TicketController struct {
	s tickets.Service
}

func New(h tickets.Service) tickets.Handler {
	return &TicketController{s: h}
}

func (tc *TicketController) GetTickets() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Unauthorized. Token is missing.", nil))
		}

		_, err := middlewares.ValidateJWT(tokenString)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Unauthorized. "+err.Error(), nil))
		}

		tickets, err := tc.s.GetTickets()
		if err != nil {
			c.Logger().Error("Failed to get all tickets", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    http.StatusInternalServerError,
				"message": "Internal Server Error",
			})
		}

		if len(tickets) == 0 {
			c.Logger().Error("Failed to get all tickets", err.Error())
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"code":    http.StatusNotFound,
				"message": "Events not found",
			})
		}

		formattedTickets := []ResponseGetTickets{}
		for _, ticket := range tickets {
			formattedTicket := ResponseGetTickets{
				TicketType:     ticket.TicketType,
				TicketCategory: ticket.TicketCategory,
				TicketPrice:    ticket.TicketPrice,
				TicketQuantity: ticket.TicketQuantity,
			}
			formattedTickets = append(formattedTickets, formattedTicket)
		}

		page := c.QueryParam("page")
		perPage := c.QueryParam("per_page")
		if page != "" || perPage == "" {
			perPage = "3"
		}
		pageInt := 1
		if page != "" {
			pageInt, _ = strconv.Atoi(page)
		}
		perPageInt, _ := strconv.Atoi(perPage)

		total := len(formattedTickets)
		totalPages := int(math.Ceil(float64(total) / float64(perPageInt)))

		startIndex := (pageInt - 1) * perPageInt
		endIndex := startIndex + perPageInt
		if endIndex > total {
			endIndex = total
		}

		data := formattedTickets[startIndex:endIndex]

		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":        http.StatusOK,
			"message":     "Successful Operation",
			"page":        pageInt,
			"per_page":    perPageInt,
			"total_pages": totalPages,
			"total_items": total,
			"data":        data,
		})
	}
}

func (tc *TicketController) CreateTicket() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Unauthorized. Token is missing.", nil))
		}

		_, err := middlewares.ValidateJWT(tokenString)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Unauthorized. "+err.Error(), nil))
		}

		eventID, err := strconv.ParseUint(c.QueryParam("eventId"), 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request. Event ID is missing.", nil))
		}

		var input RequestCreateTicket
		if err := c.Bind(&input); err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))
		}

		newTicket := tickets.Core{
			TicketType:     input.TicketType,
			TicketCategory: input.TicketCategory,
			TicketPrice:    input.TicketPrice,
			TicketQuantity: input.TicketQuantity,
			EventID:        uint(eventID),
		}

		err = tc.s.CreateTicket(newTicket, eventID)
		if err != nil {
			c.Logger().Error("Failed to create ticket: ", err)
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal Server Error", nil))
		}
		return c.JSON(http.StatusCreated, helper.ResponseFormat(http.StatusCreated, "Ticket created successfully", nil))
	}
}

func (tc *TicketController) GetTicket() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Unauthorized. Token is missing.", nil))
		}

		_, err := middlewares.ValidateJWT(tokenString)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Unauthorized. "+err.Error(), nil))
		}

		inputID := c.Param("id")
		if inputID == "" {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))
		}

		id, err := strconv.ParseUint(inputID, 10, 32)
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))
		}

		event, err := tc.s.GetTicket(uint(id))
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal Server Error", nil))
		}
		return c.JSON(http.StatusOK, helper.ResponseFormat(http.StatusOK, "Success get a ticket", event))
	}
}

func (tc *TicketController) UpdateTicket() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Unauthorized. Token is missing.", nil))
		}

		_, err := middlewares.ValidateJWT(tokenString)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Unauthorized. "+err.Error(), nil))
		}

		inputID := c.Param("id")
		if inputID == "" {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))
		}

		id, err := strconv.ParseUint(inputID, 10, 32)
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))
		}

		var input RequestUpdateTicket
		if err := c.Bind(&input); err != nil {
			c.Logger().Error("Failed to bind input: ", err)
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))
		}

		updatedTicket := tickets.Core{
			ID:             uint(id),
			TicketType:     input.TicketType,
			TicketCategory: input.TicketCategory,
			TicketPrice:    input.TicketPrice,
			TicketQuantity: input.TicketQuantity,
		}

		err = tc.s.UpdateTicket(updatedTicket.ID, updatedTicket)
		if err != nil {
			c.Logger().Error("Failed to update event: ", err)
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal Server Error", nil))
		}

		return c.JSON(http.StatusOK, helper.ResponseFormat(http.StatusCreated, "Event updated successfully", nil))
	}
}

func (tc *TicketController) DeleteTicket() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Unauthorized. Token is missing.", nil))
		}

		_, err := middlewares.ValidateJWT(tokenString)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Unauthorized. "+err.Error(), nil))
		}

		inputID := c.Param("id")
		if inputID == "" {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))
		}

		id, err := strconv.ParseUint(inputID, 10, 32)
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))
		}

		err = tc.s.DeleteTicket(uint(id))
		if err != nil {
			c.Logger().Error("Error deleting profile", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal Server Error", nil))
		}
		return c.JSON(http.StatusOK, helper.ResponseFormat(http.StatusOK, "Success deleted an account", nil))
	}
}
