package handler

import (
	"net/http"
	"math"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/wanta-zulfikri/Event-Planning-App/app/features/attendances"
	"github.com/wanta-zulfikri/Event-Planning-App/helper"
	"github.com/wanta-zulfikri/Event-Planning-App/middlewares"
)

type AttendancesController struct {
	x attendances.Service
} 

func New(w attendances.Service) attendances.Handler {
	return &AttendancesController{x:w}
} 

func (ac *AttendancesController) CreateAttendance() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization") 
		if tokenString == "" {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Unauthorized. Token is missing. ", nil)) 		
		} 

		_, err := middlewares.ValidateJWT(tokenString) 
		if err != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Unauthorized. "+err.Error(), nil))
		} 
          
        var input RequestCreateAttendances 
		if err := c.Bind(&input); err != nil {
				c.Logger().Error(err.Error()) 
				return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil)) 
		} 

		newAttendances := attendances.Core {
			UserID:         input.UserID,
			EventID:        input.EventID,
			EventCategory:  input.EventCategory, 
			TicketType:     input.TicketType, 
			Quantity:       input.Quantity,
		} 

		err = ac.x.CreateAttendance(newAttendances) 
		if err != nil {
			c.Logger().Error ("Failed to create a attendance: ", err) 
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal Server Error", nil))
		} 
		return c.JSON(http.StatusCreated, helper.ResponseFormat(http.StatusCreated, "Create a attendeces successfully", nil))
	}
} 

func (ac *AttendancesController) GetAttendance() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Unauthorized. Token is missing.", nil)) 
		} 

		_, err := middlewares.ValidateJWT(tokenString) 
		if err != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Unauthorized. "+err.Error(), nil))
		} 

		
		attendance, err := ac.x.GetAttendance()
		if err != nil {
			c.Logger().Error("Failed to get attendances", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    http.StatusInternalServerError,
				"message": "Internal Server Error",
			})
		}

		if len(attendance) == 0 {
			c.Logger().Error("Failed to get attendances", err.Error())
			return c.JSON(http.StatusNotFound, map[string]interface{}{
				"code":    http.StatusNotFound,
				"message": "Get attendances not found",
			})
		}

		formattedAttendances := []RequestGetAttendances{}
		for _, attendance := range attendance {
			formattedAttendance := RequestGetAttendances{
				ID:             attendance.ID, 
			UserID:             attendance.UserID,
			EventID:            attendance.EventID,
			EventCategory:      attendance.EventCategory, 
			TicketType:         attendance.TicketType, 
			Quantity:           attendance.Quantity,
			}
			formattedAttendances = append(formattedAttendances, formattedAttendance)
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

		total := len(formattedAttendances)
		totalPages := int(math.Ceil(float64(total) / float64(perPageInt)))

		startIndex := (pageInt - 1) * perPageInt
		endIndex := startIndex + perPageInt
		if endIndex > total {
			endIndex = total
		}

		response := formattedAttendances[startIndex:endIndex]

		pages := Pagination{
			Page:       pageInt,
			PerPage:    perPageInt,
			TotalPages: totalPages,
			TotalItems: total,
		}

		return c.JSON(http.StatusOK, attendancesResponse{
			Code:       http.StatusOK,
			Message:    "Successful operation.",
			Data:       response,
			Pagination: pages,
		})
	
	}
}