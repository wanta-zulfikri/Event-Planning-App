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
			ID: input.ID, 
			UserID: input.UserID,
			EventID: input.EventID,
			EventCategory: input.EventCategory,
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

		inputID := c.Param("id") 
		if inputID == "" {
			c.Logger().Error(err.Error()) 
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))
		} 

		id , err := strconv.ParseUint(inputID, 10, 32) 
		if err != nil {
			c.Logger().Error(err.Error()) 
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))
		} 

		attendances, err := ac.x.GetAttendance(uint(id)) 
		if err != nil {
			c.Logger().Error(err.Error()) 
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal Server Error", nil))
		} 
		return c.JSON(http.StatusOK, helper.ResponseFormat(http.StatusOK, "Success get a attendence", attendances))
	}
}