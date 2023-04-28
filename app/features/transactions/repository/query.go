package repository

import (
	"github.com/wanta-zulfikri/Event-Planning-App/app/features/events/repository"
	"github.com/wanta-zulfikri/Event-Planning-App/app/features/transactions"
	"gorm.io/gorm"
)

type TransactionRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (tr *TransactionRepository) CreateTransaction(userID uint, input transactions.Core) error {
	tx := tr.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// create transaction
	newTransaction := Transaction{
		Invoice:           input.Invoice,
		PurchaseStartDate: input.PurchaseStartDate,
		PurchaseEndDate:   input.PurchaseEndDate,
		Status:            input.Status,
		StatusDate:        input.StatusDate,
		Subtotal:          input.Subtotal,
		GrandTotal:        input.GrandTotal,
		UserID:            userID,
		EventID:           input.EventID,
	}
	err := tx.Table("transactions").Create(&newTransaction).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	// create transaction_has_tickets
	tickets := make([]repository.Ticket, len(input.Tickets))
	for i, ticket := range input.Tickets {
		tickets[i] = repository.Ticket{
			TicketCategory: ticket.TicketCategory,
			TicketPrice:    ticket.TicketPrice,
			TicketQuantity: ticket.TicketQuantity,
		}
	}
	err = tx.Table("tickets").CreateInBatches(tickets, len(tickets)).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// gorm.db.Preload("Event", func(db gorm.DB)gorm.DB {
// return Select("id,start_date,end_date,name,hosted_by,image,location")
// }).Where("user_id = ? AND status='paid'", uid).Find(&res).Error; err != nil {
// 	t.log.Errorf("error db : %v", err)
// 	return nil, errorr.NewInternal("Internal server error")
// }
