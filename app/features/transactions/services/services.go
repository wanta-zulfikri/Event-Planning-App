package services

import (
	"log"
	"time"

	"github.com/wanta-zulfikri/Event-Planning-App/app/features/transactions"
	"github.com/wanta-zulfikri/Event-Planning-App/helper"
)

type TransactionService struct {
	r transactions.Repository
}

func New(r transactions.Repository) transactions.Service {
	return &TransactionService{r: r}
}

func (ts *TransactionService) Payment(invoice string, grossAmount uint) (transactions.Payment, error) {
	payment, err := ts.r.Payment(invoice, grossAmount)
	if err != nil {
		log.Println(err)
		return transactions.Payment{}, err
	}
	return payment, nil
}

func (ts *TransactionService) GetTransaction(invoice string) (transactions.Transaction, error) {
	transaction, err := ts.r.GetTransaction(invoice)
	if err != nil {
		return transactions.Transaction{}, err
	}

	return transaction, nil
}

func (ts *TransactionService) CreateTransaction(user_id, event_id, grandtotal uint, paymentmethod string, request transactions.Transaction) (transactions.Transaction, error) {
	transaction := transactions.Transaction{
		UserID:              user_id,
		EventID:             event_id,
		Invoice:             helper.GenerateInvoice(),
		PurchaseStartDate:   time.Now(),
		PurchaseEndDate:     time.Now().Add(24 * time.Hour),
		Status:              "pending",
		StatusDate:          time.Now(),
		Transaction_Tickets: request.Transaction_Tickets,
		GrandTotal:          grandtotal,
		PaymentMethod:       paymentmethod,
	}

	createdTransaction, err := ts.r.CreateTransaction(transaction)
	if err != nil {
		return transactions.Transaction{}, err
	}

	return createdTransaction, nil
}
