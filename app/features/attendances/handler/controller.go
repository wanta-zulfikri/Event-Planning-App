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
			c.Logger().Error("Failed to get event_picture from form file: ", err)
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))
		} else if file != nil {
			event_picture, err = helper.UploadImage(c, file)
			if err != nil {
				c.Logger().Error("Failed to upload event_picture: ", err)
				return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal Server Error", nil))
			}
		}  
		newAttendances := attendances.Core{ 
            UserID:          userid,
			EventID:         uint(eventid),	
			Description:     input.Description, 
			HostedBy:        username, 
			Date:            input.Date,
			Time:            input.Time, 
			Status:          input.Status, 
			Category:        input.Category, 
			Location:        input.Location, 
			EventPicture:    input.EventPicture,
		} 
			
		_, err := ac.x.CreateAttendance(newAttendances,id) 
		if err != nil {
			c.Logger().Error("Failed to write a review:", err)
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal Server Error", nil))
		}
		return c.JSON(http.StatusCreated, helper.ResponseFormat(http.StatusCreated, "Success Created a Review", nil))
	}
}




// func (ac *AttendancesController) GetAttendance() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		tokenString := c.Request().Header.Get("Authorization")
// 		if tokenString == "" {
// 			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Unauthorized. Token is missing.", nil)) 
// 		} 

// 		_, err := middlewares.ValidateJWT(tokenString) 
// 		if err != nil {
// 			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Unauthorized. "+err.Error(), nil))
// 		} 

		
// 		attendance, err := ac.x.GetAttendance()
// 		if err != nil {
// 			c.Logger().Error("Failed to get attendances", err.Error())
// 			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
// 				"code":    http.StatusInternalServerError,
// 				"message": "Internal Server Error",
// 			})
// 		}

// 		if len(attendance) == 0 {
// 			c.Logger().Error("Failed to get attendances", err.Error())
// 			return c.JSON(http.StatusNotFound, map[string]interface{}{
// 				"code":    http.StatusNotFound,
// 				"message": "Get attendances not found",
// 			})
// 		}

// 		formattedAttendances := []RequestGetAttendances{}
// 		for _, attendance := range attendance {
// 			formattedAttendance := RequestGetAttendances{
// 				ID:             attendance.ID, 
// 			UserID:             attendance.UserID,
// 			EventID:            attendance.EventID,
// 			EventCategory:      attendance.EventCategory, 
// 			TicketType:         attendance.TicketType, 
// 			Quantity:           attendance.Quantity,
// 			}
// 			formattedAttendances = append(formattedAttendances, formattedAttendance)
// 		}

// 		page := c.QueryParam("page")
// 		perPage := c.QueryParam("per_page")
// 		if page != "" || perPage == "" {
// 			perPage = "3"
// 		}
// 		pageInt := 1
// 		if page != "" {
// 			pageInt, _ = strconv.Atoi(page)
// 		}
// 		perPageInt, _ := strconv.Atoi(perPage)

// 		total := len(formattedAttendances)
// 		totalPages := int(math.Ceil(float64(total) / float64(perPageInt)))

// 		startIndex := (pageInt - 1) * perPageInt
// 		endIndex := startIndex + perPageInt
// 		if endIndex > total {
// 			endIndex = total
// 		}

// 		response := formattedAttendances[startIndex:endIndex]

// 		pages := Pagination{
// 			Page:       pageInt,
// 			PerPage:    perPageInt,
// 			TotalPages: totalPages,
// 			TotalItems: total,
// 		}

// 		return c.JSON(http.StatusOK, attendancesResponse{
// 			Code:       http.StatusOK,
// 			Message:    "Successful operation.",
// 			Data:       response,
// 			Pagination: pages,
// 		})
	
// 	}
// }