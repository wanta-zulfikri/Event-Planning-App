package handler

import (
	"fmt"
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

func (ec *EventController) GetEvents() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Unauthorized. Token is missing.", nil))
		}

		_, err := middlewares.ValidateJWT(tokenString)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Unauthorized. "+err.Error(), nil))
		}

		events, err := ec.s.GetEvents()
		if err != nil {
			c.Logger().Error("Failed to get all books", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    http.StatusInternalServerError,
				"message": "Internal Server Error",
			})
		}

		if len(events) == 0 {
			c.Logger().Error("Failed to get all events", err.Error())
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"code":    http.StatusNotFound,
				"message": "Events not found",
			})
		}

		formattedEvents := []ResponseGetEvents{}
		for _, event := range events {
			formattedEvent := ResponseGetEvents{
				Title:       event.Title,
				Description: event.Description,
				EventDate:   event.EventDate,
				EventTime:   event.EventTime,
				Status:      event.Status,
				Category:    event.Category,
				Location:    event.Location,
				Image:       event.Image,
				Username:    event.Username,
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

		data := formattedEvents[startIndex:endIndex]

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

func (ec *EventController) CreateEvent() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input RequestCreateEvent
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
			Username:    username,
		}

		err = ec.s.CreateEvent(newEvent, uint(id))
		if err != nil {
			c.Logger().Error("Failed to create event: ", err)
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal Server Error", nil))
		}
		return c.JSON(http.StatusCreated, helper.ResponseFormat(http.StatusCreated, "Event created successfully", nil))
	}
}

func (ec *EventController) GetEvent() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Unauthorized. Token is missing.", nil))
		}

		_, err := middlewares.ValidateJWT(tokenString)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Unauthorized. "+err.Error(), nil))
		}

		id, err := strconv.ParseUint(c.Param("id"), 10, 32)
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))
		}

		event, err := ec.s.GetEvent(uint(id))
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal Server Error", nil))
		}

		response := ResponseGetEvents{
			Title:       event.Title,
			Description: event.Description,
			EventDate:   event.EventDate,
			EventTime:   event.EventTime,
			Status:      event.Status,
			Category:    event.Category,
			Location:    event.Location,
			Image:       event.Image,
			Username:    event.Username,
		}

		return c.JSON(http.StatusOK, helper.ResponseFormat(http.StatusOK, "Success get an event", response))
	}
}

func (ec *EventController) UpdateEvent() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input RequestUpdateEvent
		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Unauthorized. Token is missing.", nil))
		}

		username, err := middlewares.ValidateJWTUsername(tokenString)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Unauthorized. "+err.Error(), nil))
		}

		if username == "" {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Unauthorized. Invalid token.", nil))
		}

		id, err := strconv.ParseUint(c.Param("id"), 10, 32)
		if err != nil {
			c.Logger().Error("Failed to parse ID from URL param: ", err)
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))
		}

		if err := c.Bind(&input); err != nil {
			c.Logger().Error("Failed to bind input from request body: ", err)
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))
		}

		file, err := c.FormFile("image")
		var image string
		if err != nil && err != http.ErrMissingFile {
			c.Logger().Error("Failed to get image from form file: ", err)
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))
		} else if file != nil {
			image, err = helper.UploadImage(c, file)
			if err != nil {
				c.Logger().Error("Failed to upload image: ", err)
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
			Image:       image,
			Username:    username,
		}

		err = ec.s.UpdateEvent(updatedEvent.ID, updatedEvent)
		if err != nil {
			c.Logger().Error("Failed to update event: ", err)
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal Server Error", nil))
		}

		return c.JSON(http.StatusOK, helper.ResponseFormat(http.StatusOK, "Event updated successfully", nil))
	}
}

func (ec *EventController) DeleteEvent() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Unauthorized. Token is missing.", nil))
		}

		_, err := middlewares.ValidateJWT(tokenString)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Unauthorized. "+err.Error(), nil))
		}

		id, err := strconv.ParseUint(c.Param("id"), 10, 32)
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))
		}

		err = ec.s.DeleteEvent(uint(id))
		if err != nil {
			c.Logger().Error("Error deleting profile", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal Server Error", nil))
		}
		return c.JSON(http.StatusOK, helper.ResponseFormat(http.StatusOK, "Success deleted an account", nil))
	}
}
