package handler

// import (
// 	"fmt"
// 	"net/http"

// 	"github.com/golang-jwt/jwt"
// 	"github.com/labstack/echo/v4"
// 	"github.com/wanta-zulfikri/Event-Planning-App/app/features/transactions"
// 	"github.com/wanta-zulfikri/Event-Planning-App/helper"
// 	"github.com/wanta-zulfikri/Event-Planning-App/middlewares"
// )

// type TransactionController struct {
// 	s transactions.Service
// }

// func New(h transactions.Service) transactions.Handler {
// 	return &TransactionController{s: h}
// }

// type TransactionInput struct {
// 	event_name      string // receive this event id from unique valu of event name
// 	ticket_category string
// 	ticket_quantity uint
// 	payment_method  string
// }

// func (tc *TransactionController) CreateTransaction() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		var input = struct {
// 			Data []int `json:"data"`
// 		}{}
// 		if err := c.Bind(&input); err != nil {
// 			c.Logger().Error(err.Error())
// 			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))
// 		}
// 		fmt.Println(len(input.Data))
// 		if len(input.Data) == 0 {
// 			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Data tidak boleh kosong", nil))
// 		}
// 		userId := middlewares.GetUserID(c.Get("user").(*jwt.Token))
// 		err := tc.s.CreateTransaction(transactions.Core{}, userId)
// 		if err != nil {
// 			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal Server Error", nil))
// 		}
// 		return c.JSON(http.StatusOK, helper.ResponseFormat(http.StatusCreated, "Success created a transaction", nil))
// 	}
// }
