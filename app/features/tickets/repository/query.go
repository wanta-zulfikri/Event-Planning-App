package repository

import (
	"errors"
	"fmt"
	"time"

	"github.com/wanta-zulfikri/Event-Planning-App/app/features/tickets"
	"gorm.io/gorm"
)

type TicketRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *TicketRepository {
	return &TicketRepository{db: db}
}

func (tr *TicketRepository) GetTickets(id uint) ([]tickets.Core, error) {
	var cores []tickets.Core
	if err := tr.db.Table("tickets").Where("event_id = ? AND deleted_at IS NULL", id).Find(&cores).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []tickets.Core{}, fmt.Errorf("ticket with id %v not found", id)
		}
		return []tickets.Core{}, err
	}
	return cores, nil
}

func (tr *TicketRepository) UpdateTicket(eventID uint, updatedTickets []tickets.Core) error {
	for _, updatedTicket := range updatedTickets {
		if err := tr.db.Table("tickets").Where("event_id = ? AND ticket_category = ?", eventID, updatedTicket.TicketCategory).Updates(map[string]interface{}{
			"ticket_category": updatedTicket.TicketCategory,
			"ticket_price":    updatedTicket.TicketPrice,
			"ticket_quantity": updatedTicket.TicketQuantity,
		}).Error; err != nil {
			return err
		}
	}
	return nil
}

func (tr *TicketRepository) DeleteTicket(id uint) error {
	input := Ticket{}
	if err := tr.db.Where("id = ?", id).Find(&input).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		return err
	}

	input.DeletedAt = gorm.DeletedAt{Time: time.Now(), Valid: true}

	if err := tr.db.Save(&input).Error; err != nil {
		return err
	}
	return nil
}
