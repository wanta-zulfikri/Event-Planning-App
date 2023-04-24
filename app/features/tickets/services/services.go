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

func (ts *TicketService) GetTickets() ([]tickets.Core, error) {
	tickets, err := ts.r.GetTickets()
	if err != nil {
		return nil, err
	}
	return tickets, nil
}

func (ts *TicketService) CreateTicket(newTicket tickets.Core, eventid uint64) error {
	_, err := ts.r.CreateTicket(newTicket, eventid)
	if err != nil {
		return err
	}
	return nil
}

func (ts *TicketService) GetTicket(id uint) (tickets.Core, error) {
	event, err := ts.r.GetTicket(id)
	if err != nil {
		return tickets.Core{}, err
	}
	return event, nil
}

func (ts *TicketService) UpdateTicket(id uint, updatedTicket tickets.Core) error {
	updatedTicket.ID = id
	if err := ts.r.UpdateTicket(id, updatedTicket); err != nil {
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
