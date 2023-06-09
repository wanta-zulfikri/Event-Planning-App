package handler

import (
	"math"
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

func (ec *EventController) GetEvent() echo.HandlerFunc {
	return func(c echo.Context) error {
		eventid, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			c.Logger().Error("Failed to parse ID from URL param: ", err)
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))
		}

		event, err := ec.s.GetEvent(uint(eventid))
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusNotFound, helper.ResponseFormat(http.StatusNotFound, "The requested resource was not found.", nil))
		}

		response := ResponseGetEvent{
			ID:            event.ID,
			Title:         event.Title,
			Description:   event.Description,
			Hosted_by:     event.Hostedby,
			Date:          event.EventDate,
			Time:          event.EventTime,
			Status:        event.Status,
			Category:      event.Category,
			Location:      event.Location,
			Event_picture: event.Image,
			Transactions:  make([]ResponseTransactions, 0),
			Reviews:       make([]ResponseReviews, 0),
		}

		for _, t := range event.Transactions {
			transaction := ResponseTransactions{
				Username:    t.Username,
				UserPicture: t.Image,
			}
			response.Transactions = append(response.Transactions, transaction)
		}

		for _, r := range event.Reviews {
			review := ResponseReviews{
				Username:    r.Username,
				UserPicture: r.Image,
				Review:      r.Review,
			}
			response.Reviews = append(response.Reviews, review)
		}

		return c.JSON(http.StatusOK, GetEventResponse{
			Code:    http.StatusOK,
			Message: "Successful operation.",
			Data:    response,
		})
	}
}

func (ec *EventController) CreateEventWithTickets() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input RequestCreateEventWithTickets
		tokenString := c.Request().Header.Get("Authorization")
		claims, err := middlewares.ValidateJWT2(tokenString)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Missing or Malformed JWT"+err.Error(), nil))
		}

		id := claims.ID
		username := claims.Username

		if err := c.Bind(&input); err != nil {
			c.Logger().Error("Failed to bind input: ", err)
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))
		}

		//masuk service
		eventTickets := make([]events.TicketCore, len(input.Tickets))
		for i, ticket := range input.Tickets {
			eventTickets[i] = events.TicketCore{
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

		err = ec.s.CreateEventWithTickets(newEvent, id)
		if err != nil {
			c.Logger().Error("Failed to create event with tickets: ", err)
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal Server Error", nil))
		}

		response := EventResponse{
			Code:    http.StatusCreated,
			Message: "Success created an event",
			Data: EventData{
				Title:       newEvent.Title,
				Description: newEvent.Description,
				HostedBy:    newEvent.Hostedby,
				Date:        newEvent.EventDate,
				Time:        newEvent.EventTime,
				Status:      newEvent.Status,
				Category:    newEvent.Category,
				Location:    newEvent.Location,
				Picture:     newEvent.Image,
				Tickets:     make([]TicketResponse, len(newEvent.Tickets)),
			},
		}

		for i, ticket := range newEvent.Tickets {
			response.Data.Tickets[i] = TicketResponse{
				Category: ticket.TicketCategory,
				Price:    ticket.TicketPrice,
				Quantity: ticket.TicketQuantity,
			}
		}

		return c.JSON(http.StatusCreated, response)
	}
}

func (ec *EventController) GetEvents() echo.HandlerFunc {
	return func(c echo.Context) error {
		category := c.QueryParam("category")
		var events []events.Event
		var err error

		if category != "" {
			events, err = ec.s.GetEventsByCategory(category)
			if err != nil {
				c.Logger().Error(err.Error())
				return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusNotFound, "The requested resource was not found.", nil))
			}
		} else {
			events, err = ec.s.GetEvents()
			if err != nil {
				c.Logger().Error(err.Error())
				return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusNotFound, "The requested resource was not found.", nil))
			}
		}

		if len(events) == 0 {
			if err != nil {
				c.Logger().Error(err.Error())
				return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusNotFound, "The requested resource was not found.", nil))
			}
		}

		formattedEvents := []ResponseGetEvents{}
		for _, event := range events {
			formattedEvent := ResponseGetEvents{
				ID:            event.ID,
				Title:         event.Title,
				Description:   event.Description,
				Hosted_by:     event.Hostedby,
				Date:          event.EventDate,
				Time:          event.EventTime,
				Status:        event.Status,
				Category:      event.Category,
				Location:      event.Location,
				Event_picture: event.Image,
			}
			formattedEvents = append(formattedEvents, formattedEvent)
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

		total := len(formattedEvents)
		totalPages := int(math.Ceil(float64(total) / float64(perPageInt)))

		startIndex := (pageInt - 1) * perPageInt
		endIndex := startIndex + perPageInt
		if endIndex > total {
			endIndex = total
		}

		response := formattedEvents[startIndex:endIndex]

		pages := Pagination{
			Page:       pageInt,
			PerPage:    perPageInt,
			TotalPages: totalPages,
			TotalItems: total,
		}

		return c.JSON(http.StatusOK, EventsResponse{
			Code:       http.StatusOK,
			Message:    "Successful operation.",
			Data:       response,
			Pagination: pages,
		})
	}
}

func (ec *EventController) GetEventsByAttendance() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization") 
		claims, err := middlewares.ValidateJWT2(tokenString) 
		if err != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Missing or Malformed JWT. "+err.Error(), nil))
		} 

		userid := claims.ID 
		events, err := ec.s.GetEventsByAttendance(userid) 
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusNotFound, "The requested resource was not found.", nil))
		} 

		
		if len(events) == 0 {
			if err != nil {
				c.Logger().Error(err.Error())
				return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusNotFound, "The requested resource was not found.", nil))
			}
		} 

		
		formattedEvents := []ResponseGetAttendances{}
		for _, event := range events {
			formattedEvent := ResponseGetAttendances{
				ID:            event.ID,
				Title:         event.Title,
				Description:   event.Description,
				HostedBy:      event.Hostedby,
				Date:          event.EventDate,
				Time:          event.EventTime,
				Status:        event.Status,
				Category:      event.Category,
				Location:      event.Location,
				Image:         event.Image,
			}
			formattedEvents = append(formattedEvents, formattedEvent)
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

		total := len(formattedEvents)
		totalPages := int(math.Ceil(float64(total) / float64(perPageInt)))

		startIndex := (pageInt - 1) * perPageInt
		endIndex := startIndex + perPageInt
		if endIndex > total {
			endIndex = total
		}

		Response := formattedEvents[startIndex:endIndex] 

		pages := Pagination{
			Page:       pageInt,
			PerPage:    perPageInt,
			TotalPages: totalPages,
			TotalItems: total,
		}

		return c.JSON(http.StatusOK,Getattendance{
			Code:       http.StatusOK,
			Message:    "Successful operation.",
			Data:       Response,
			Pagination: pages,
		})
	}
}


func (ec *EventController) GetEventsByUserID() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		claims, err := middlewares.ValidateJWT2(tokenString)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Missing or Malformed JWT. "+err.Error(), nil))
		}

		userid := claims.ID
		events, err := ec.s.GetEventsByUserID(userid)
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusNotFound, "The requested resource was not found.", nil))
		}

		if len(events) == 0 {
			if err != nil {
				c.Logger().Error(err.Error())
				return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusNotFound, "The requested resource was not found.", nil))
			}
		}

		formattedEvents := []ResponseGetEvents{}
		for _, event := range events {
			formattedEvent := ResponseGetEvents{
				ID:            event.ID,
				Title:         event.Title,
				Description:   event.Description,
				Hosted_by:     event.Hostedby,
				Date:          event.EventDate,
				Time:          event.EventTime,
				Status:        event.Status,
				Category:      event.Category,
				Location:      event.Location,
				Event_picture: event.Image,
			}
			formattedEvents = append(formattedEvents, formattedEvent)
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

		total := len(formattedEvents)
		totalPages := int(math.Ceil(float64(total) / float64(perPageInt)))

		startIndex := (pageInt - 1) * perPageInt
		endIndex := startIndex + perPageInt
		if endIndex > total {
			endIndex = total
		}

		response := formattedEvents[startIndex:endIndex]

		pages := Pagination{
			Page:       pageInt,
			PerPage:    perPageInt,
			TotalPages: totalPages,
			TotalItems: total,
		}

		return c.JSON(http.StatusOK, EventsResponse{
			Code:       http.StatusOK,
			Message:    "Successful operation.",
			Data:       response,
			Pagination: pages,
		})
	}
}

func (ec *EventController) UpdateEvent() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input RequestUpdateEvent
		tokenString := c.Request().Header.Get("Authorization")
		claims, err := middlewares.ValidateJWT2(tokenString)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Missing or Malformed JWT. "+err.Error(), nil))
		}

		username := claims.Username
		id := claims.ID
		if err := c.Bind(&input); err != nil {
			c.Logger().Error("Failed to bind input from request body: ", err)
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))
		}

		file, err := c.FormFile("event_picture")
		var event_picture string
		if err != nil && err != http.ErrMissingFile {
			c.Logger().Error("Failed to get event_picture from form file: ", err)
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))
		} else if file != nil {
			event_picture, err = helper.UploadImage(c, file)
			if err != nil {
				c.Logger().Error("Failed to upload event_picture: ", err)
				return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal Server Error", nil))
			}
		}

		updatedEvent := events.Core{
			ID:          uint(id),
			Title:       input.Title,
			Description: input.Description,
			EventDate:   input.EventDate,
			EventTime:   input.EventTime,
			Status:      input.Status,
			Category:    input.Category,
			Location:    input.Location,
			Image:       event_picture,
			Hostedby:    username,
		}

		err = ec.s.UpdateEvent(updatedEvent.ID, updatedEvent)
		if err != nil {
			c.Logger().Error("Failed to update event: ", err)
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal Server Error", nil))
		}

		response := ResponseUpdateEvent{
			Title:         updatedEvent.Title,
			Description:   updatedEvent.Description,
			Hosted_by:     updatedEvent.Hostedby,
			Date:          updatedEvent.EventDate,
			Time:          updatedEvent.EventTime,
			Status:        updatedEvent.Status,
			Category:      updatedEvent.Category,
			Location:      updatedEvent.Location,
			Event_picture: updatedEvent.Image,
		}

		return c.JSON(http.StatusOK, helper.DataResponse{
			Code:    http.StatusOK,
			Message: "Success updated an event.",
			Data:    response,
		})
	}
}

func (ec *EventController) DeleteEvent() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		_, err := middlewares.ValidateJWT2(tokenString)
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Missing or Malformed JWT. "+err.Error(), nil))
		}

		id, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))
		}

		err = ec.s.DeleteEvent(uint(id))
		if err != nil {
			c.Logger().Error("Error deleting profile", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal Server Error", nil))
		}

		return c.JSON(http.StatusOK, helper.ResponseFormat(http.StatusOK, "Success deleted an event", nil))
	}
}
