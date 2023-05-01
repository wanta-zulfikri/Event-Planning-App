package services

import (
	"github.com/wanta-zulfikri/Event-Planning-App/app/features/transactions"
)

type TransactionService struct {
	r transactions.Repository
}

func New(r transactions.Repository) transactions.Service {
	return &TransactionService{r: r}
}

func (ts *TransactionService) CreateTransaction(userid uint, eventid uint, request transactions.Transaction) error {
	transaction := make([]transactions.Transaction_Tickets, len(request.Transaction_Tickets))
	for i, item := range request.Transaction_Tickets {
		transaction[i] = transactions.Transaction_Tickets{
			TicketID: item.TicketID,
		}
	}

	Transaction := transactions.Transaction{
		UserID:              userid,
		EventID:             eventid,
		Transaction_Tickets: transaction,
	}

	err := ts.r.CreateTransaction(Transaction)
	if err != nil {
		return err
	}

	return nil
}
