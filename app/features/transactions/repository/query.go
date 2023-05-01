package repository

import (
	_ "github.com/lib/pq"
	"github.com/wanta-zulfikri/Event-Planning-App/app/features/transactions"
	"gorm.io/gorm"
)

type TransactionRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (tr *TransactionRepository) CreateTransaction(input transactions.Transaction) error {
	err := tr.db.Create(&input).Error
	if err != nil {
		return err
	}

	// Insert transaction_tickets into transaction_tickets table
	var tickets []transactions.Transaction_Tickets
	for _, t := range input.Transaction_Tickets {
		ticket := transactions.Transaction_Tickets{
			TransactionID: input.ID,
			TicketID:      t.TicketID,
		}
		tickets = append(tickets, ticket)
	}

	err = tr.db.Model(&input).Association("Transaction_Tickets").Append(tickets)
	if err != nil {
		return err
	}

	return nil
}
