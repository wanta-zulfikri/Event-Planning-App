package services

import (
	"github.com/pandudpn/go-payment-gateway/gateway/midtrans"
	"github.com/wanta-zulfikri/Event-Planning-App/app/features/transactions"
)

type TransactionService struct {
	r transactions.Repository
}

func New(r transactions.Repository) transactions.Service {
	return &TransactionService{r: r}
}

func (ts *TransactionService) PayTransaction(userID uint, eventID uint, paymentType string, request transactions.Transaction) (*midtrans.ChargeResponse, error) {
	vaaccount, err := ts.r.PayTransaction(userID, eventID, paymentType, request)
	if err != nil {
		return nil, err
	}

	return vaaccount, nil
}

// func (ts *TransactionService) GetTransaction(transactionid uint) (transactions.Transaction, error) {
// 	transaction, err := ts.r.GetTransaction(transactionid)
// 	if err != nil {
// 		return transactions.Transaction{}, err
// 	}

// 	return transaction, nil
// }

// func (ts *TransactionService) CreateTransaction(user_id, event_id, grandtotal uint, paymentmethod string, request transactions.Transaction) error {
// 	Transaction := transactions.Transaction{
// 		UserID:              user_id,
// 		EventID:             event_id,
// 		Invoice:             helper.GenerateInvoice(),
// 		PurchaseStartDate:   time.Now(),
// 		PurchaseEndDate:     time.Now().Add(24 * time.Hour),
// 		Status:              "pending",
// 		StatusDate:          time.Now(),
// 		Transaction_Tickets: request.Transaction_Tickets,
// 		GrandTotal:          grandtotal,
// 		PaymentMethod:       paymentmethod,
// 	}

// 	err := ts.r.CreateTransaction(Transaction)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
