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

func (tr *TransactionRepository) CreateTransaction(newTransaction *transactions.Core) error {
	err := tr.db.Model(&Transaction{}).Create(newTransaction).Error
	if err != nil {
		return err
	}
	return nil
}

func (tr *TransactionRepository) GetInvoice(Invoice string) (*Transaction, error) {
	transaction := &Transaction{}

	err := tr.db.Model(&Transaction{}).Where("invoice = ?", transaction).Take(&transaction).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, helper.ErrRecordNotFound
		}
		return nil, err
	}
	return transaction, nil
}
