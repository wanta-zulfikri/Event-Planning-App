package repository

import (
	"errors"
	"time"

	"github.com/wanta-zulfikri/Event-Planning-App/app/features/events"
	"github.com/wanta-zulfikri/Event-Planning-App/app/features/tickets/repository"
	"gorm.io/gorm"
)

type EventRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *EventRepository {
	return &EventRepository{db: db}
}

func (er *EventRepository) CreateEventWithTickets(event events.Core, userID uint) error {
	tx := er.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// create event
	newEvent := Event{
		Title:       event.Title,
		Description: event.Description,
		EventDate:   event.EventDate,
		EventTime:   event.EventTime,
		Status:      event.Status,
		Category:    event.Category,
		Location:    event.Location,
		Image:       event.Image,
		Hostedby:    event.Hostedby,
		UserID:      userID,
	}
	err := tx.Table("events").Create(&newEvent).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	// create tickets
	tickets := make([]repository.Ticket, len(event.Tickets))
	for i, ticket := range event.Tickets {
		tickets[i] = repository.Ticket{
			TicketType:     ticket.TicketType,
			TicketCategory: ticket.TicketCategory,
			TicketPrice:    ticket.TicketPrice,
			TicketQuantity: ticket.TicketQuantity,
			EventID:        newEvent.ID,
		}
	}
	err = tx.Table("tickets").CreateInBatches(tickets, len(tickets)).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (er *EventRepository) GetEvents() ([]events.Core, error) {
	var cores []events.Core
	if err := er.db.Table("events").Where("deleted_at IS NULL").Find(&cores).Error; err != nil {
		return nil, err
	}
	return cores, nil
}

func (er *EventRepository) GetEvent(eventid, userid uint) (events.Core, error) {
	var input Event
	result := er.db.Where("id = ? AND user_id = ?", eventid, userid).Find(&input)
	if result.Error != nil {
		return events.Core{}, result.Error
	}
	if result.RowsAffected == 0 {
		return events.Core{}, result.Error
	}
	return events.Core{
		Title:       input.Title,
		Description: input.Description,
		EventDate:   input.EventDate,
		EventTime:   input.EventTime,
		Status:      input.Status,
		Category:    input.Category,
		Location:    input.Location,
		Image:       input.Image,
		Hostedby:    input.Hostedby,
	}, nil
}

func (er *EventRepository) UpdateEvent(id uint, updatedEvent events.Core) error {
	input := Event{}
	if err := er.db.Where("id = ?", id).First(&input).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		return err
	}
	input.Title = updatedEvent.Title
	input.Description = updatedEvent.Description
	input.EventDate = updatedEvent.EventDate
	input.EventTime = updatedEvent.EventTime
	input.Status = updatedEvent.Status
	input.Category = updatedEvent.Category
	input.Location = updatedEvent.Location
	input.Image = updatedEvent.Image
	input.Hostedby = updatedEvent.Hostedby
	input.UpdatedAt = time.Now()

	if err := er.db.Save(&input).Error; err != nil {
		return err
	}
	return nil
}

func (er *EventRepository) DeleteEvent(id uint) error {
	input := Event{}
	if err := er.db.Where("id = ?", id).Find(&input).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		return err
	}

	input.DeletedAt = gorm.DeletedAt{Time: time.Now(), Valid: true}

	if err := er.db.Save(&input).Error; err != nil {
		return err
	}
	return nil
}
