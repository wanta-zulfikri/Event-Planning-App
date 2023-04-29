package repository

import (
	"github.com/wanta-zulfikri/Event-Planning-App/app/features/transactions"
	"github.com/wanta-zulfikri/Event-Planning-App/helper"
	"gorm.io/gorm"
)

type TransactionRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (tr *TransactionRepository) CreateTransaction(input transactions.Core) error {
	tx := tr.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	invoice := helper.GenerateInvoice()

	newTransaction := Transaction{
		ID:                invoice,
		PurchaseStartDate: input.PurchaseStartDate,
		PurchaseEndDate:   input.PurchaseEndDate,
		Status:            input.Status,
		StatusDate:        input.StatusDate,
		GrandTotal:        input.GrandTotal,
		UserID:            input.UserID,
		EventID:           input.EventID,
	}
	err := tx.Table("transactions").Create(&newTransaction).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	tickets := make([]transactions.TransactionTickets, len(input.TransactionTickets))
	for i, ticket := range input.TransactionTickets {
		tickets[i] = transactions.TransactionTickets{
			TransactionID:  invoice,
			TicketID:       ticket.TicketID,
			TicketCategory: ticket.TicketCategory,
			TicketPrice:    ticket.TicketPrice,
			TicketQuantity: ticket.TicketQuantity,
			Subtotal:       ticket.TicketPrice * ticket.TicketQuantity,
		}
	}
	err = tx.Table("transaction_tickets").CreateInBatches(tickets, len(tickets)).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
