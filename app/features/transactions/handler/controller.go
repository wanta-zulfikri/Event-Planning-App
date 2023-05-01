package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/wanta-zulfikri/Event-Planning-App/app/features/transactions"
	"github.com/wanta-zulfikri/Event-Planning-App/helper"
	"github.com/wanta-zulfikri/Event-Planning-App/middlewares"
)

type TransactionController struct {
	s transactions.Service
}

func New(s transactions.Service) transactions.Handler {
	return &TransactionController{s: s}
}

func (tc *TransactionController) GetTransaction() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		_, err := middlewares.ValidateJWT2(tokenString)
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Missing or Malformed JWT"+err.Error(), nil))
		}

		transactionid, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			c.Logger().Error("Failed to parse ID from URL param: ", err)
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))
		}

		transaction, err := tc.s.GetTransaction(uint(transactionid))
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal Server Error", nil))
		}

		return c.JSON(http.StatusOK, transaction)
	}
}

func (tc *TransactionController) CreateTransaction() echo.HandlerFunc {
	return func(c echo.Context) error {
		var request RequestCreateTransaction
		tokenString := c.Request().Header.Get("Authorization")
		claims, err := middlewares.ValidateJWT2(tokenString)
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Missing or Malformed JWT"+err.Error(), nil))
		}

		user_id := claims.ID
		if err := c.Bind(&request); err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))
		}

		if len(request.ItemDescription) == 0 {
			c.Logger().Error("Item description can not be empty")
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Item description can not be empty", nil))
		}

		event_id := request.EventID
		grandtotal := request.GrandTotal
		paymentmethod := request.PaymentMethod
		var input transactions.Transaction
		for _, t := range request.ItemDescription {
			input.Transaction_Tickets = append(input.Transaction_Tickets, transactions.Transaction_Tickets{
				TicketID:       t.TicketID,
				TicketCategory: t.TicketCategory,
				TicketPrice:    t.TicketPrice,
				TicketQuantity: t.TicketQuantity,
				Subtotal:       t.Subtotal,
			})
		}

		err = tc.s.CreateTransaction(user_id, event_id, grandtotal, paymentmethod, input)
		if err != nil {
			c.Logger().Error("Failed to create transaction: ", err)
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal Server Error", nil))
		}

		return c.JSON(http.StatusCreated, helper.ResponseFormat(http.StatusCreated, "Success Created a Transaction", nil))
	}
}
