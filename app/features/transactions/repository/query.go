package repository

import (
	"fmt"

	pg "github.com/pandudpn/go-payment-gateway"
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

func (tr *TransactionRepository) PayTransaction(user_id uint, event_id uint, paymenttype string, request transactions.RequestPayTransaction) (string, error) {
	// Prepare transaction details
	transactionDetails := []transactions.TransactionDetails{}
	for _, td := range request.Transaction_Details {
		transactionDetails = append(transactionDetails, transactions.TransactionDetails{
			OrderID:     helper.GenerateInvoice(),
			GrossAmount: td.GrossAmount,
		})
	}

	// Call BankTransferCharge method to get the virtual account details
	opts := &pg.Options{
		ClientId:  helper.SandBoxClientKey,
		ServerKey: helper.SandBoxServerKey,
	}
	virtualAccount, err := helper.BankTransferCharge(opts)
	if err != nil {
		return "", fmt.Errorf("failed to get virtual account details: %w", err)
	}

	transaction := transactions.Transaction{
		UserID:              user_id,
		EventID:             event_id,
		PaymentType:         paymenttype,
		Transaction_Details: transactionDetails,
		VAAccount:           virtualAccount,
	}

	// Save the transaction to the database
	err = tr.db.Create(&transaction).Error
	if err != nil {
		return "", fmt.Errorf("failed to create transaction: %w", err)
	}

	return virtualAccount, nil
}

// func (tr *TransactionRepository) GetTransaction(transactionid uint) (transactions.Transaction, error) {
// 	var transaction transactions.Transaction
// 	if err := tr.db.
// 		Where("transactions.id = ?", transactionid).
// 		Preload("Transaction_Tickets").
// 		Joins("JOIN users ON users.id = transactions.user_id").
// 		Joins("JOIN events ON events.id = transactions.event_id").
// 		Select("transactions.*, users.username, users.email, events.title, events.event_date, events.event_time").
// 		First(&transaction).Error; err != nil {
// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			return transactions.Transaction{}, errors.New("Transaction not found")
// 		}
// 		return transactions.Transaction{}, fmt.Errorf("Failed to retrieve transaction from database: %w", err)
// 	}
// 	return transaction, nil
// }

// func (tr *TransactionRepository) CreateTransaction(request transactions.Transaction) error {
// 	var err error
// 	tx := tr.db.Begin()

// 	if err := tx.Model(&request).
// 		Create(map[string]interface{}{
// 			"user_id":             request.UserID,
// 			"event_id":            request.EventID,
// 			"invoice":             request.Invoice,
// 			"purchase_start_date": request.PurchaseStartDate,
// 			"purchase_end_date":   request.PurchaseEndDate,
// 			"status":              request.Status,
// 			"status_date":         request.StatusDate,
// 			"grand_total":         request.GrandTotal,
// 			"payment_method":      request.PaymentMethod,
// 		}).Error; err != nil {
// 		tx.Rollback()
// 		return err
// 	}

// 	tx.Last(&request)

// 	for i := range request.Transaction_Tickets {
// 		request.Transaction_Tickets[i].TransactionID = request.ID
// 		err = tx.Create(&request.Transaction_Tickets[i]).Error
// 		if err != nil {
// 			tx.Rollback()
// 			return err
// 		}
// 	}

// 	err = tx.Commit().Error
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
