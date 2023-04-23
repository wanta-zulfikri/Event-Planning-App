package repository

import (
	"errors"
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

func (tr *TicketRepository) GetTickets() ([]tickets.Core, error) {
	var cores []tickets.Core
	if err := tr.db.Table("tickets").Find(&cores).Error; err != nil {
		return nil, err
	}
	return cores, nil
}

func (tr *TicketRepository) CreateTicket(newTicket tickets.Core) (tickets.Core, error) {
	input := tickets.Ticket{
		TicketType:     newTicket.TicketType,
		TicketCategory: newTicket.TicketCategory,
		TicketPrice:    newTicket.TicketPrice,
		TicketQuantity: newTicket.TicketQuantity,
	}

	err := tr.db.Table("tickets").Create(&input).Error
	if err != nil {
		return tickets.Core{}, err
	}

	createdTicket := tickets.Core{
		TicketType:     input.TicketType,
		TicketCategory: input.TicketCategory,
		TicketPrice:    input.TicketPrice,
		TicketQuantity: input.TicketQuantity,
	}
	return createdTicket, nil
}

func (tr *TicketRepository) GetTicket(id uint) (tickets.Core, error) {
	var input tickets.Ticket
	result := tr.db.Where("id = ?", id).Find(&input)
	if result.Error != nil {
		return tickets.Core{}, result.Error
	}
	if result.RowsAffected == 0 {
		return tickets.Core{}, result.Error
	}
	return tickets.Core{
		TicketType:     input.TicketType,
		TicketCategory: input.TicketCategory,
		TicketPrice:    input.TicketPrice,
		TicketQuantity: input.TicketQuantity,
	}, nil
}

func (tr *TicketRepository) UpdateTicket(id uint, updatedTicket tickets.Core) error {
	input := tickets.Ticket{}
	if err := tr.db.Where("id = ?", id).First(&input).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		return err
	}
	input.TicketType = updatedTicket.TicketType
	input.TicketCategory = updatedTicket.TicketCategory
	input.TicketPrice = updatedTicket.TicketPrice
	input.TicketQuantity = updatedTicket.TicketQuantity
	input.UpdatedAt = time.Now()

	if err := tr.db.Save(&input).Error; err != nil {
		return err
	}
	return nil
}

func (tr *TicketRepository) DeleteTicket(id uint) error {
	input := tickets.Ticket{}
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