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

type RequestPayTransaction struct {
	EventID             uint                 `json:"event_id"`
	PaymentType         string               `json:"payment_type"`
	Transaction_Details []TransactionDetails `json:"transaction_details"`
	BankTransfer        *BankTransferDetails `json:"bank_transfer,omitempty"`
}

type TransactionDetails struct {
	OrderID     string `json:"order_id"`
	GrossAmount uint   `json:"gross_amount"`
}

type BankTransferDetails struct {
	Bank string `json:"bank"`
}

func (tc *TransactionController) PayTransaction() echo.HandlerFunc {
	return func(c echo.Context) error {
		var request RequestPayTransaction
		tokenString := c.Request().Header.Get("Authorization")
		claims, err := middlewares.ValidateJWT2(tokenString)
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Missing or Malformed JWT"+err.Error(), nil))
		}

		userID := claims.ID
		if err := c.Bind(&request); err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))
		}

		if len(request.Transaction_Details) == 0 {
			c.Logger().Error("Item description can not be empty")
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Item description can not be empty", nil))
		}

		eventID := request.EventID
		paymentType := request.PaymentType
		var input transactions.RequestPayTransaction
		for _, item := range request.Transaction_Details {
			input.Transaction_Details = append(input.Transaction_Details, transactions.TransactionDetails{
				OrderID:     item.OrderID,
				GrossAmount: item.GrossAmount,
			})
		}

		// Call midtrans API to get virtual account details
		vaaccount, err := tc.s.PayTransaction(userID, eventID, paymentType, input)
		if err != nil {
			c.Logger().Error("Failed to create transaction: ", err)
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal Server Error", nil))
		}

		return c.JSON(http.StatusCreated, helper.DataResponse{
			Code:    http.StatusCreated,
			Message: "Success Created a Transaction.",
			Data:    map[string]interface{}{"va_account": vaaccount},
		})
	}
}

// func (tc *TransactionController) GetTransaction() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		tokenString := c.Request().Header.Get("Authorization")
// 		claims, err := middlewares.ValidateJWT2(tokenString)
// 		if err != nil {
// 			c.Logger().Error(err.Error())
// 			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Missing or Malformed JWT"+err.Error(), nil))
// 		}

// 		attendee_email := claims.Email
// 		attendee := claims.Username
// 		transactionid, err := strconv.ParseUint(c.Param("id"), 10, 64)
// 		if err != nil {
// 			c.Logger().Error("Failed to parse ID from URL param: ", err)
// 			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))
// 		}

// 		transaction, err := tc.s.GetTransaction(uint(transactionid))
// 		if err != nil {
// 			c.Logger().Error(err.Error())
// 			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal Server Error", nil))
// 		}

// 		response := TransactionResponse{
// 			Code:    http.StatusOK,
// 			Message: "Successful Operation",
// 			Data: ResponseGetTransaction{
// 				Invoice:           transaction.Invoice,
// 				Seller:            transaction.Username,
// 				SEmail:            transaction.Email,
// 				Attendee:          attendee,
// 				AEmail:            attendee_email,
// 				Title:             transaction.Title,
// 				EventDate:         transaction.EventDate,
// 				EventTime:         transaction.EventTime,
// 				PurchaseStartDate: transaction.PurchaseStartDate.Format("2006-01-02 15:04:05"),
// 				PurchaseEndDate:   transaction.PurchaseEndDate.Format("2006-01-02 15:04:05"),
// 				Status:            transaction.Status,
// 				StatusDate:        transaction.StatusDate.Format("2006-01-02 15:04:05"),
// 				ItemDescription:   make([]ResponseTickets, 0),
// 				GrandTotal:        transaction.GrandTotal,
// 				PaymentMethod:     transaction.PaymentMethod,
// 			},
// 		}

// 		for _, t := range transaction.Transaction_Tickets {
// 			item := ResponseTickets{
// 				TicketCategory: t.TicketCategory,
// 				TicketPrice:    t.TicketPrice,
// 				TicketQuantity: t.TicketQuantity,
// 				Subtotal:       t.Subtotal,
// 			}
// 			response.Data.ItemDescription = append(response.Data.ItemDescription, item)
// 		}

// 		return c.JSON(http.StatusOK, response)
// 	}
// }

// func (tc *TransactionController) CreateTransaction() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		var request RequestCreateTransaction
// 		tokenString := c.Request().Header.Get("Authorization")
// 		claims, err := middlewares.ValidateJWT2(tokenString)
// 		if err != nil {
// 			c.Logger().Error(err.Error())
// 			return c.JSON(http.StatusUnauthorized, helper.ResponseFormat(http.StatusUnauthorized, "Missing or Malformed JWT"+err.Error(), nil))
// 		}

// 		user_id := claims.ID
// 		if err := c.Bind(&request); err != nil {
// 			c.Logger().Error(err.Error())
// 			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Bad Request", nil))
// 		}

// 		if len(request.ItemDescription) == 0 {
// 			c.Logger().Error("Item description can not be empty")
// 			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "Item description can not be empty", nil))
// 		}

// 		event_id := request.EventID
// 		grandtotal := request.GrandTotal
// 		paymentmethod := request.PaymentMethod
// 		var input transactions.Transaction
// 		for _, t := range request.ItemDescription {
// 			input.Transaction_Tickets = append(input.Transaction_Tickets, transactions.Transaction_Tickets{
// 				TicketID:       t.TicketID,
// 				TicketCategory: t.TicketCategory,
// 				TicketPrice:    t.TicketPrice,
// 				TicketQuantity: t.TicketQuantity,
// 				Subtotal:       t.Subtotal,
// 			})
// 		}

// 		err = tc.s.CreateTransaction(user_id, event_id, grandtotal, paymentmethod, input)
// 		if err != nil {
// 			c.Logger().Error("Failed to create transaction: ", err)
// 			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, "Internal Server Error", nil))
// 		}

// 		return c.JSON(http.StatusCreated, helper.ResponseFormat(http.StatusCreated, "Success Created a Transaction", nil))
// 	}
// }
