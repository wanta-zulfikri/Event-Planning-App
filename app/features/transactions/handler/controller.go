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

func (tc *TransactionController) CreateTransaction() echo.HandlerFunc {
	return func(c echo.Context) error {
		var request RequestCreateTransaction
		tokenString := c.Request().Header.Get("Authorization")
		claims, err := middlewares.ValidateJWT2(tokenString)
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Missing or Malformed JWT"+err.Error(), nil))
		}

		userid := claims.ID
		eventid, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			c.Logger().Error("Failed to parse ID from URL param: ", err)
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))
		}

		if err := c.Bind(&request); err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))
		}

		if len(request.ItemDescription) == 0 {
			c.Logger().Error("Item description can not empty")
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Item description can not empty", nil))
		}

		var input transactions.Carts
		for _, t := range request.ItemDescription {
			ticket := transactions.Ticket{
				TicketCategory: t.TicketCategory,
				TicketPrice:    t.TicketPrice,
				TicketQuantity: t.TicketQuantity,
			}
			input.ItemDescription = append(input.ItemDescription, ticket)
		}

		err = tc.s.CreateTransaction(userid, uint(eventid), input)
		if err != nil {
			c.Logger().Error("Failed to create transaction: ", err)
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal Server Error", nil))
		}

		return c.JSON(http.StatusCreated, helper.ResponseFormat(http.StatusCreated, "Success created transaction", nil))
	}
}
