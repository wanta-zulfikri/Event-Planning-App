package services

import (
	"time"

	"github.com/wanta-zulfikri/Event-Planning-App/app/features/transactions"
)

type TransactionService struct {
	r transactions.Repository
}

func New(r transactions.Repository) transactions.Service {
	return &TransactionService{r: r}
}

func (ts *TransactionService) CreateTransaction(userid uint, eventid uint, request transactions.Carts) error {
	transaction := make([]transactions.TransactionTickets, len(request.ItemDescription))
	subtotal := uint(0)
	grandTotal := uint(0)
	for i, item := range request.ItemDescription {
		transaction[i] = transactions.TransactionTickets{
			TicketCategory: item.TicketCategory,
			TicketPrice:    item.TicketPrice,
			TicketQuantity: item.TicketQuantity,
			Subtotal:       subtotal,
		}
		subtotal = item.TicketPrice * item.TicketQuantity
		grandTotal += subtotal
	}

	Transaction := transactions.Core{
		PurchaseStartDate:  time.Now(),
		PurchaseEndDate:    time.Now().Add(24 * time.Hour),
		Status:             "pending",
		StatusDate:         time.Now(),
		GrandTotal:         grandTotal,
		UserID:             userid,
		EventID:            eventid,
		TransactionTickets: transaction,
	}

	err := ts.r.CreateTransaction(Transaction)
	if err != nil {
		return err
	}

	return nil
}
