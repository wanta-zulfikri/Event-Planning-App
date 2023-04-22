package handler

import (
	"Event-Planning-App/app/features/events"
	"Event-Planning-App/helper"
	"Event-Planning-App/middlewares"
	"net/http"
	"strconv"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"

)

type EventHandler struct {
	srv events.EventService
}

func New(service events.EventService) *EventHandler {
	return &EventHandler{
		srv: service,
	}
} 

func (ev *EventHandler) Add(c echo.Context) error {
	addInput := EventRequest{} 
	addInput.UserID = uint(middlewares.ExtractToken(c)) 
	if err := c.Bind(&addInput); err != nil {
		c.Logger().Error("terjadi kesalahan bind", err.Error()) 
		return c.JSON(helper.ErrorResponse(err))
	} 

	newEvent := events.Core{}
	copier.Copy(&newEvent, addInput) 

	file, _ := c.FormFile("Image") 

	err := ev.srv.Add(newEvent, file) 
	if err != nil {
		c.Logger().Error("terjadi kesalahan saat add event",  err.Error())
		return c.JSON(helper.ErrorResponse(err))
	}
	return c.JSON(helper.SuccessResponse(http.StatusCreated, "add event successfully"))
}


func (ev *EventHandler) MyEvent(c echo.Context) error {
	userID := int(middlewares.ExtractToken(c)) 
	var pageNumber int = 1 
	pageParam := c.QueryParam("page") 
	if pageParam != "" {
		pageConv, errConv := strconv.Atoi(pageParam) 
		if errConv != nil {
			c.Logger().Error("terjadi kesalahan") 
			return c.JSON(http.StatusInternalServerError,helper.Response("Failed, page must number"))

		} else {
			pageNumber = pageConv
		}
	}

	data, err := ev.srv.MyEvent(userID, pageNumber)
	if err != nil {
		c.Logger().Error("terjadi kesalahan") 
		return c.JSON(http.StatusInternalServerError, helper.Response("Failed, error read data"))
	} 
	dataResponse := CoreToGetAllEventResp(data) 
	return c.JSON(http.StatusOK, helper.ResponseWithData("Success", dataResponse))
} 

func (ev *EventHandler) GetAll(c echo.Context) error {
	var pageNumber int = 1 
	pageParam := c.QueryParam("page") 
	if pageParam != "" {
		pageConv, errConv := strconv.Atoi(pageParam) 
		if errConv != nil {
			c.Logger().Error("terjadi kesalahan")  
			return c.JSON(http.StatusInternalServerError, helper.Response("Failed, page must number"))

		} else {
			pageNumber = pageConv
		}
	}

	nameParam := c.QueryParam("name") 
	data, err := ev.srv.GetAll(pageNumber, nameParam) 
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.Response("Failed, error read data"))
	} 
	dataResponse := CoreToGetAllEventResp(data)
	return c.JSON(http.StatusOK, helper.ResponseWithData("success", dataResponse))
} 

func (ev *EventHandler) Update(c echo.Context) error {
	userID := int(middlewares.ExtractToken(c)) 
	eventID, errCnv := strconv.Atoi(c.Param("id")) 
	if errCnv != nil {
		c.Logger().Error("event tidak ditemukan") 
		return errCnv
	} 

	updateInput := EventRequest{} 
	if err := c.Bind(&updateInput); err != nil {
		c.Logger().Error("terjadi kesalahan bind", err.Error())
		return c.JSON(helper.ErrorResponse(err))
	} 

	file, _ := c.FormFile("Image") 

	updateEvent := events.Core{} 
	copier.Copy(&updateEvent, &updateInput) 
	err := ev.srv.Update(userID, eventID, updateEvent, file) 
	if err != nil{
		c.Logger().Error("terjadi kesalahan update event", err.Error()) 
		return c.JSON(helper.ErrorResponse(err))
	} 
	return c.JSON(helper.SuccessResponse(http.StatusOK, "update event successfully updated"))
} 

func (ev *EventHandler) GetEventById(c echo.Context) error {
	eventID, errCnv := strconv.Atoi(c.Param("id")) 
	if errCnv != nil {
		c.Logger().Error("terjadi kesalahan ") 
		return errCnv
	}
	data, err := ev.srv.GetEventById(eventID) 
	if err != nil {
		c.Logger().Error("terjadi kesalahan", err.Error()) 
		return c.JSON(helper.ErrorResponse(err))
	} 
	res := MyEventResponse{} 
	copier.Copy(&res, &data) 
	return c.JSON(helper.SuccessResponse(http.StatusOK, "detail event successfully displayed",res))
} 


func (ev *EventHandler) DeleteEvent(C echo.Context) error {
	userID := int(middlewares.ExtractToken(C)) 
	eventID, errCnv := strconv.Atoi(C.Param("id")) 
	if errCnv != nil {
		C.Logger().Error("terjadi kesalahan") 
		return errCnv
	} 
	err := ev.srv.DeleteEvent(userID, eventID) 
	if err != nil {
		C.Logger().Error("terjadi kesalahan", err.Error()) 
		return C.JSON(helper.ErrorResponse(err)) 
	} 
	return C.JSON(helper.SuccessResponse(http.StatusOK, "detail event deleted"))
}
