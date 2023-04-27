package services

import (
	"errors"

	"github.com/wanta-zulfikri/Event-Planning-App/app/features/tickets"
	"gorm.io/gorm"
)

type TicketService struct {
	r tickets.Repository
}

func New(r tickets.Repository) tickets.Service {
	return &TicketService{r: r}
}

func (ts *TicketService) GetTickets(id uint) ([]tickets.Core, error) {
	tickets, err := ts.r.GetTickets(id)
	if err != nil {
		return nil, err
	}
	return tickets, nil
}

func (ts *TicketService) UpdateTicket(eventid uint, updatedTicket []tickets.Core) error {
	if err := ts.r.UpdateTicket(eventid, updatedTicket); err != nil {
		return err
	}
	return nil
}

func (ts *TicketService) DeleteTicket(id uint) error {
	err := ts.r.DeleteTicket(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		return err
	}
	return nil
}
