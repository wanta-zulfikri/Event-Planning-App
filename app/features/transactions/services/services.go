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

func (ts *TransactionService) CreateTransaction(userid uint, input transactions.Core) error {
	err := ts.r.CreateTransaction(userid, input)
	if err != nil {
		return err
	}

	return nil
}
