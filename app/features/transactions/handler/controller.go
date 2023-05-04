package handler

import (
	"net/http"

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

type RequestCreatePayment struct {
	Invoice      string `json:"invoice"`
	Gross_Amount uint   `json:"gross_amount"`
}

func (tc *TransactionController) Payment() echo.HandlerFunc {
	return func(c echo.Context) error {
		var request RequestCreatePayment
		tokenString := c.Request().Header.Get("Authorization")
		_, err := middlewares.ValidateJWT2(tokenString)
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Missing or Malformed JWT"+err.Error(), nil))
		}

		if err := c.Bind(&request); err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))
		}

		if request.Gross_Amount == 0 {
			c.Logger().Error("Gross Amount can not be zero")
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Gross Amount can not be zero", nil))
		}

		payment, err := tc.s.Payment(request.Invoice, request.Gross_Amount)
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal Server Error", nil))
		}
		return c.JSON(http.StatusOK, helper.DataResponse{
			Code:    http.StatusOK,
			Message: "Successful Operation",
			Data:    &payment,
		})
	}
}

func (tc *TransactionController) GetTransaction() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		claims, err := middlewares.ValidateJWT2(tokenString)
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Missing or Malformed JWT"+err.Error(), nil))
		}

		attendee_email := claims.Email
		attendee := claims.Username
		invoice := c.Param("id")
		if invoice == "" {
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Missing invoice parameter", nil))
		}
		transaction, err := tc.s.GetTransaction(invoice)
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal Server Error", nil))
		}

		response := TransactionResponse{
			Code:    http.StatusOK,
			Message: "Successful Operation",
			Data: ResponseGetTransaction{
				Invoice:           transaction.Invoice,
				Seller:            transaction.Username,
				SEmail:            transaction.Email,
				Attendee:          attendee,
				AEmail:            attendee_email,
				Title:             transaction.Title,
				EventDate:         transaction.EventDate,
				EventTime:         transaction.EventTime,
				PurchaseStartDate: transaction.PurchaseStartDate.Format("2006-01-02 15:04:05"),
				PurchaseEndDate:   transaction.PurchaseEndDate.Format("2006-01-02 15:04:05"),
				Status:            transaction.Status,
				StatusDate:        transaction.StatusDate.Format("2006-01-02 15:04:05"),
				ItemDescription:   make([]ResponseTickets, 0),
				GrandTotal:        transaction.GrandTotal,
				PaymentMethod:     transaction.PaymentMethod,
			},
		}

		for _, t := range transaction.Transaction_Tickets {
			item := ResponseTickets{
				TicketCategory: t.TicketCategory,
				TicketPrice:    t.TicketPrice,
				TicketQuantity: t.TicketQuantity,
				Subtotal:       t.Subtotal,
			}
			response.Data.ItemDescription = append(response.Data.ItemDescription, item)
		}

		return c.JSON(http.StatusOK, response)
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

		transaction, err := tc.s.CreateTransaction(user_id, event_id, grandtotal, paymentmethod, input)
		if err != nil {
			c.Logger().Error("Failed to create transaction: ", err)
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal Server Error", nil))
		}
		return c.JSON(http.StatusOK, helper.DataResponse{
			Code:    http.StatusOK,
			Message: "Successful Operation",
			Data:    map[string]any{"invoice": transaction.Invoice},
		})
	}
}
