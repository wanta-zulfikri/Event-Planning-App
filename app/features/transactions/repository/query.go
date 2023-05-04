package repository

import (
	"errors"
	"fmt"
	"strings"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/wanta-zulfikri/Event-Planning-App/app/features/transactions"
	"github.com/wanta-zulfikri/Event-Planning-App/config/common"
	"gorm.io/gorm"
)

type TransactionRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (tr *TransactionRepository) Payment(invoice string, grossAmount uint) (transactions.Payment, error) {
	var c = coreapi.Client{}
	c.New(common.MidstransServerKey, midtrans.Sandbox)

	request := &coreapi.ChargeReq{
		PaymentType: coreapi.PaymentTypeBankTransfer,
		BankTransfer: &coreapi.BankTransferDetails{
			Bank: midtrans.BankBca,
		},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  invoice,
			GrossAmt: int64(grossAmount),
		},
	}

	//chargeResponse
	resp, err := c.ChargeTransaction(request)
	if err != nil {
		fmt.Println(err)
	}

	banks := make([]string, len(resp.VaNumbers))
	for i, bank := range resp.VaNumbers {
		banks[i] = bank.Bank
	}
	banksStr := strings.Join(banks, ",")

	vaNumbers := make([]string, len(resp.VaNumbers))
	for i, vaNumber := range resp.VaNumbers {
		vaNumbers[i] = vaNumber.VANumber
	}
	vaNumbersStr := strings.Join(vaNumbers, ",")

	// simpan chargeResponse dalam database
	payment := transactions.Payment{
		Transaction_ID:     resp.TransactionID,
		Order_ID:           resp.OrderID,
		Gross_Amount:       resp.GrossAmount,
		Payment_Type:       resp.PaymentType,
		Bank:               banksStr,
		Transaction_Time:   resp.TransactionTime,
		Transaction_Status: resp.TransactionStatus,
		Va_Numbers:         vaNumbersStr,
	}
	result := tr.db.Create(&payment)
	if result.Error != nil {
		err := fmt.Errorf("failed to create payment: %v", result.Error)
		return transactions.Payment{}, err
	}

	return payment, nil
}

func (tr *TransactionRepository) GetTransaction(invoice string) (transactions.Transaction, error) {
	if invoice == "" {
		return transactions.Transaction{}, errors.New("Missing invoice parameter")
	}
	var transaction transactions.Transaction
	if err := tr.db.
		Where("transactions.invoice = ?", invoice).
		Preload("Transaction_Tickets").
		Joins("JOIN users ON users.id = transactions.user_id").
		Joins("JOIN events ON events.id = transactions.event_id").
		Select("transactions.*, users.username, users.email, events.title, events.event_date, events.event_time").
		First(&transaction).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return transactions.Transaction{}, errors.New("Transaction not found")
		}
		return transactions.Transaction{}, fmt.Errorf("Failed to retrieve transaction from database: %w", err)
	}
	return transaction, nil
}

func (tr *TransactionRepository) CreateTransaction(request transactions.Transaction) (transactions.Transaction, error) {
	var err error
	tx := tr.db.Begin()

	if err = tx.Model(&request).
		Create(map[string]interface{}{
			"user_id":             request.UserID,
			"event_id":            request.EventID,
			"invoice":             request.Invoice,
			"purchase_start_date": request.PurchaseStartDate,
			"purchase_end_date":   request.PurchaseEndDate,
			"status":              request.Status,
			"status_date":         request.StatusDate,
			"grand_total":         request.GrandTotal,
			"payment_method":      request.PaymentMethod,
		}).Error; err != nil {
		tx.Rollback()
		return transactions.Transaction{}, err
	}

	if err = tx.Last(&request).Error; err != nil {
		tx.Rollback()
		return transactions.Transaction{}, err
	}

	for i := range request.Transaction_Tickets {
		request.Transaction_Tickets[i].TransactionID = request.ID
		if err = tx.Create(&request.Transaction_Tickets[i]).Error; err != nil {
			tx.Rollback()
			return transactions.Transaction{}, err
		}
	}

	if err = tx.Commit().Error; err != nil {
		return transactions.Transaction{}, err
	}

	return request, nil
}
