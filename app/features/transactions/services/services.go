package services

import (
	"errors"
	"time"

	"time"

	"github.com/wanta-zulfikri/Event-Planning-App/app/features/transactions"
	"github.com/wanta-zulfikri/Event-Planning-App/app/helper"
)

type TransactionService struct {
	r transactions.Repository
}

func New(r transactions.Repository) transactions.Service {
	return &TransactionService{r: r}
}

func (ts *TransactionService) CreateTransaction(userID uint, eventID uint, transaction transactions.TransactionTicketCore) error {
	// melakukan validasi input
	if input.TicketQuantity < 1 {
		return errors.New("invalid ticket quantity")
	}

	// melakukan penghitungan subtotal
	subtotal := input.TicketPrice * input.TicketQuantity

	// memasukkan data ke struct TransactionTickets
	ticket := TransactionTickets{
		TicketCategory: input.TicketCategory,
		TicketPrice:    input.TicketPrice,
		TicketQuantity: input.TicketQuantity,
		Subtotal:       subtotal,
	}

	// memasukkan data ke struct Transaction
	transaction := Transaction{
		Invoice:            helper.GenerateInvoice(),
		PurchaseStartDate:  time.Now(),
		PurchaseEndDate:    time.Now().Add(24 * time.Hour),
		Status:             "pending",
		StatusDate:         time.Now(),
		GrandTotal:         1000,
		UserID:             1,
		EventID:            1,
		TransactionTickets: []TransactionTickets{ticket},
	}

	// memanggil repository untuk menyimpan data ke database
	err := ts.r.CreateTransaction(transaction).Error
	if err != nil {
		return err
	}

	return nil
}
