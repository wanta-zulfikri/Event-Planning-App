package handler

import (
	"net/http"
	"strconv"
	"time"

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
		var input RequestCreateTransaction
		tokenString := c.Request().Header.Get("Authorization")
		claims, err := middlewares.ValidateJWT2(tokenString)
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Missing or Malformed JWT"+err.Error(), nil))
		}

		id := claims.ID
		eventid, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			c.Logger().Error("Failed to parse ID from URL param: ", err)
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))
		}

		if err := c.Bind(&input); err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))
		}
		if len(input.ItemDescription) == 0 {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Item description can not empty", nil))
		}

		transaction := make([]transactions.TicketCore, len(input.ItemDescription))
		subtotal := uint(0)
		grandTotal := uint(0)
		for i, item := range input.ItemDescription {
			transaction[i] = transactions.TicketCore{
				TicketCategory: item.TicketCategory,
				TicketPrice:    item.TicketPrice,
				TicketQuantity: item.TicketQuantity,
			}
			subtotal = item.TicketPrice * item.TicketQuantity
			grandTotal += subtotal
		}

		Transaction := transactions.Core{
			Invoice:           helper.GenerateInvoice(),
			PurchaseStartDate: time.Now(),
			PurchaseEndDate:   time.Now().Add(24 * time.Hour),
			Status:            "pending",
			StatusDate:        time.Now(),
			Tickets:           transaction,
			Subtotal:          subtotal,
			GrandTotal:        grandTotal,
			UserID:            id,
			EventID:           uint(eventid),
		}

		// Create transaction
		err = tc.s.CreateTransaction(id, Transaction)
		if err != nil {
			c.Logger().Error("Failed to create transaction: ", err)
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal Server Error", nil))
		}

		return c.JSON(http.StatusCreated, helper.ResponseFormat(http.StatusCreated, "Success created transaction", nil))
	}
}
