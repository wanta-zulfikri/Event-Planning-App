package services

// import (
// 	"time"

// 	"github.com/wanta-zulfikri/Event-Planning-App/app/features/events"
// 	"github.com/wanta-zulfikri/Event-Planning-App/app/features/transactions"
// 	"github.com/wanta-zulfikri/Event-Planning-App/app/features/users"
// 	"github.com/wanta-zulfikri/Event-Planning-App/helper"
// )

// type TransactionService struct {
// 	m transactions.Repository
// 	u users.Repository
// 	e events.Repository
// }

// func New(r transactions.Repository, u users.Repository, e events.Repository) transactions.Service {
// 	return &TransactionService{m: r, u: u, e: e}
// }

// func (ts TransactionService) CreateTransaction(userID uint, inputs []transactions.TransactionInput) error {
// 	var (
// 		tickets     []transactions.TransactionDetailInput
// 		grandTotal  uint
// 		subtotal    uint
// 		transaction transactions.Core
// 	)

// func (ts TransactionService) CreateTransaction(newTransaction transactions.Core) (transactions.Core, error) {
// 	// Retrieve user and event data

// 	// Calculate total transaction
// 	var subtotal uint = 0
// 	for _, ticket := range newTransaction.Tickets {
// 		subtotal += ticket.TicketPrice * ticket.TicketQuantity
// 	}
// 	var grandTotal uint = subtotal

// 	const (
// 		TransactionStatusPending = "pending"
// 		TransactionStatusPaid    = "paid"
// 	)

// 	// pada saat membuat transaksi baru
// 	transactionStatus := TransactionStatusPending

// 	// Initiate the transaction
// 	statusDate := time.Now()
// 	newTransaction.Invoice = helper.GenerateInvoice()
// 	newTransaction.Status = transactionStatus
// 	newTransaction.StatusDate = statusDate
// 	newTransaction.Subtotal = subtotal
// 	newTransaction.GrandTotal = grandTotal

// 	// Create transaction
// 	createdTransaction, err := ts.m.CreateTransaction(newTransaction)
// 	if err != nil {
// 		return transactions.Core{}, err
// 	}

// 	// Setup response from my app
// 	response := transactions.Core{
// 		ID:                createdTransaction.ID,
// 		Invoice:           createdTransaction.Invoice,
// 		PurchaseStartDate: createdTransaction.PurchaseStartDate,
// 		PurchaseEndDate:   createdTransaction.PurchaseEndDate,
// 		Status:            createdTransaction.Status,
// 		StatusDate:        createdTransaction.StatusDate,
// 		Subtotal:          createdTransaction.Subtotal,
// 		GrandTotal:        createdTransaction.GrandTotal,
// 		UserID:            createdTransaction.UserID,
// 		EventID:           createdTransaction.EventID,
// 		Tickets:           newTransaction.Tickets,
// 	}

// 	return response, nil
// }

// cek event yang user pilih, jika eventid tersedia maka appen to slice event
// calculate total transaction
// initiate the transaction, then create transaction
// initiate the transaction detail with loop over the events above and append to slice eventAttends
// update availability of each event to "0", which is unavailable
// then create the transaction detail with status
// setup response from my app
