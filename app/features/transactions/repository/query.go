package repository

import (
	"errors"

	"github.com/wanta-zulfikri/Event-Planning-App/app/features/transactions"
	"github.com/wanta-zulfikri/Event-Planning-App/helper"
	"gorm.io/gorm"
)

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (tr *TransactionRepository) CreateTransaction(newTransaction transactions.Core) (transactions.Core, error) {
	input := transactions.Transaction{
		Model:             gorm.Model{},
		Invoice:           newTransaction.Invoice,
		PurchaseStartDate: newTransaction.PurchaseStartDate,
		PurchaseEndDate:   newTransaction.PurchaseEndDate,
		Status:            newTransaction.Status,
		StatusDate:        newTransaction.StatusDate,
		Tickets:           []transactions.Ticket{},
		Subtotal:          newTransaction.Subtotal,
		GrandTotal:        newTransaction.GrandTotal,
		UserID:            newTransaction.UserID,
		EventID:           newTransaction.EventID,
	}

	err := tr.db.Create(&input).Error
	if err != nil {
		return transactions.Core{}, err
	}

	for _, ticket := range newTransaction.Tickets {
		ticketInput := transactions.Ticket{
			TransactionID:  input.ID, // set foreign key ke transaksi yang baru saja dibuat
			TicketType:     ticket.TicketType,
			TicketCategory: ticket.TicketCategory,
			TicketPrice:    ticket.TicketPrice,
			TicketQuantity: ticket.TicketQuantity,
			EventID:        ticket.EventID,
		}

		err = tr.db.Create(&ticketInput).Error
		if err != nil {
			return transactions.Core{}, err
		}
	}

	createdTransaction := transactions.Core{
		ID:                input.ID,
		Invoice:           input.Invoice,
		PurchaseStartDate: input.PurchaseStartDate,
		PurchaseEndDate:   input.PurchaseEndDate,
		Status:            input.Status,
		StatusDate:        input.StatusDate,
		Subtotal:          input.Subtotal,
		GrandTotal:        input.GrandTotal,
		UserID:            input.UserID,
		EventID:           input.EventID,
		Tickets:           newTransaction.Tickets,
	}
	return createdTransaction, nil
}

func (tr *TransactionRepository) GetInvoice(Invoice string) (*transactions.Transaction, error) {
	transaction := &transactions.Transaction{}

	err := tr.db.Model(&transactions.Transaction{}).Where("invoice = ?", transaction).Take(&transaction).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, helper.ErrRecordNotFound
		}
		return nil, err
	}
	return transaction, nil
}
