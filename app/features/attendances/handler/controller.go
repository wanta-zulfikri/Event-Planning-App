package handler

import (
	"net/http"
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
		var input RequestCreateAttendances
		tokenString := c.Request().Header.Get("Authorization") 
		claims, err := middlewares.ValidateJWT2(tokenString)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Missing or Malformed JWT"+err.Error(), nil))
		}

		userid := claims.ID 
		username := claims.Username 
		eventid, err := strconv.ParseUint(c.Param("id"), 10, 64) 
		if err != nil {
			c.Logger().Error("Failed to parse ID from URL param: ", err)
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))
		}
        if err := c.Bind(&input); err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))
		}
		file, err := c.FormFile("event_picture")
		var event_picture string
		if err != nil && err != http.ErrMissingFile {
			c.Logger().Error("Failed to get event_picture form file: ", err)
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))
		} else if file != nil {
			event_picture, err = helper.UploadImage(c, file)
			if err != nil {
				c.Logger().Error("Failed to upload event_picture: ", err)
				return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal Server Error", nil))
			}
		}  
		newAttendances := attendances.Core{  
			ID:              userid,
			EventID:         uint(eventid),	
			Description:     input.Description, 
			HostedBy:        username, 
			Date:            input.Date,
			Time:            input.Time, 
			Status:          input.Status, 
			Category:        input.Category, 
			Location:        input.Location, 
			EventPicture:    event_picture,
		} 
			
		err = ac.x.CreateAttendance(newAttendances) 
		if err != nil {
			c.Logger().Error("Failed to create attendances:", err)
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal Server Error", nil))
		}
		return c.JSON(http.StatusCreated, helper.ResponseFormat(http.StatusCreated, "Success Created attendances", nil))
	}
}


func (ac *AttendancesController) GetAttendance() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		_, err := middlewares.ValidateJWT2(tokenString)
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Missing or Malformed JWT. "+err.Error(), nil))
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

		
		file, err := c.FormFile("event_picture")
		var event_picture string
		if err != nil && err != http.ErrMissingFile {
			c.Logger().Error("Failed to get event_picture from file: ", err)
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))
		} else if file != nil {
			event_picture, err = helper.UploadImage(c, file)
			if err != nil {
				c.Logger().Error("Failed to upload event_picture: ", err)
				return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal Server Error", nil))
			}
		}
        
		attendances, err := ac.x.GetAttendance(id)
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusNotFound, "The requested resource was not found.", nil))
		} 


		var response []ResponseGetAttendances
		for _, Attendance := range attendances {
			response = append(response, ResponseGetAttendances{
				ID             :   Attendance.ID,          
				EventID        :   Attendance.EventID,
				Title          :   Attendance.Title,
				Description    :   Attendance.Description, 
				HostedBy       :   Attendance.HostedBy,
				Date           :   Attendance.Date,
				Time           :   Attendance.Time,
				Status         :   Attendance.Status,
				Location       :   Attendance.Location,
				EventPicture   :   event_picture,
				Category       :   Attendance.Category,
			})
		}

		return c.JSON(http.StatusOK, helper.DataResponse{
			Code:    http.StatusOK,
			Message: "Successful operation.",
			Data:    response,
		})
		return c.JSON(http.StatusCreated, helper.ResponseFormat(http.StatusCreated, "success Get attendances", nil))
	}
}
