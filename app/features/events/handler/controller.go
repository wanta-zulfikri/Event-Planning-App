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
				Hostedby:    event.UserID,
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
			UserID:      input.UserID,
		}

		err = ec.s.CreateEvent(newEvent)
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

		event, err := ec.s.GetEvent(uint(id))
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal Server Error", nil))
		}
		return c.JSON(http.StatusOK, helper.ResponseFormat(http.StatusOK, "Success get an event", event))
	}
}

func (ec *EventController) UpdateEvent() echo.HandlerFunc {
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

		var input RequestUpdateEvent
		if err := c.Bind(&input); err != nil {
			c.Logger().Error("Failed to bind input: ", err)
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))
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
			UserID:      input.UserID,
		}

		file, err := c.FormFile("image")
		if err != nil {
			c.Logger().Error("Failed to get image: ", err)
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))
		}

		if file != nil {
			image, err := helper.UploadImage(c, file)
			if err != nil {
				c.Logger().Error("Failed to upload image: ", err)
				return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal Server Error", nil))
			}

			updatedEvent.Image = image
		}

		err = ec.s.UpdateEvent(updatedEvent.ID, updatedEvent)
		if err != nil {
			c.Logger().Error("Failed to update event: ", err)
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal Server Error", nil))
		}

		return c.JSON(http.StatusOK, helper.ResponseFormat(http.StatusCreated, "Event updated successfully", nil))
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

		err = ec.s.DeleteEvent(uint(id))
		if err != nil {
			c.Logger().Error("Error deleting profile", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal Server Error", nil))
		}
		return c.JSON(http.StatusOK, helper.ResponseFormat(http.StatusOK, "Success deleted an account", nil))
	}
}
