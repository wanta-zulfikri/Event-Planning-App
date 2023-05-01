package repository

import (
	"errors"
	"fmt"

	"github.com/wanta-zulfikri/Event-Planning-App/app/features/transactions"
	"gorm.io/gorm"
)

type TransactionRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (tr *TransactionRepository) GetTransaction(transactionid uint) (transactions.Transaction, error) {
	var transaction transactions.Transaction
	if err := tr.db.Where("id = ?", transactionid).Preload("Transaction_Tickets").First(&transaction).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return transactions.Transaction{}, errors.New("Transaction not found")
		}
		return transactions.Transaction{}, fmt.Errorf("Failed to retrieve transaction from database: %w", err)
	}
	return transaction, nil
}

func (tr *TransactionRepository) CreateTransaction(request transactions.Transaction) error {
	err := tr.db.Create(&request).Error
	if err != nil {
		return err
	}
	return nil
}
