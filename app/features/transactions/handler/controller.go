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

func NewTransactionController(s transactions.Service) *TransactionController {
	return &TransactionController{s: s}
}

type Tickets struct {
	TicketCategory string `json:"ticket_category"`
	TicketPrice    uint   `json:"ticket_price"`
	TicketQuantity uint   `json:"ticket_quantity"`
}

func (tc *TransactionController) CreateTransaction() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		claims, err := middlewares.ValidateJWT2(tokenString)
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Missing or Malformed JWT"+err.Error(), nil))
		}

		userID := claims.ID
		eventID, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			c.Logger().Error("Failed to parse ID from URL param: ", err)
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))
		}

		if err := c.Bind(&input); err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))
		}

		err = tc.s.CreateTransaction(userID, uint(eventID), input)
		if err != nil {
			c.Logger().Error("Failed to create transaction: ", err)
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal Server Error", nil))
		}

		return c.JSON(http.StatusCreated, helper.ResponseFormat(http.StatusCreated, "Success created transaction", nil))
	}
}
