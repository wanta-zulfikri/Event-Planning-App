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

func (er *EventRepository) GetEvents() ([]events.Event, error) {
	var cores []events.Event
	if err := er.db.Table("events").Where("deleted_at IS NULL").Find(&cores).Error; err != nil {
		return nil, err
	}
	return cores, nil
}

func (er *EventRepository) GetEventsByCategory(category string) ([]events.Event, error) {
	var cores []events.Event
	if err := er.db.Table("events").Where("category = ? AND deleted_at IS NULL", category).Find(&cores).Error; err != nil {
		return nil, err
	}
	return cores, nil
}

func (er *EventRepository)GetEventsByAttendance(userID uint) ([]events.Event, error) {
	var cores []events.Event 
	if err := er.db.Table("attendances").Where("user_id = ? = ? AND deleted_at IS NULL", userID).Find(&cores).Error; err != nil {
		return nil, err
	} 
	return cores, nil
}

func (er *EventRepository) GetEventsByUserID(userid uint) ([]events.Event, error) {
	var cores []events.Event
	if err := er.db.Table("events").Where("user_id = ? AND deleted_at IS NULL", userid).Find(&cores).Error; err != nil {
		return nil, err
	}
	return cores, nil
}

func (er *EventRepository) GetEvent(eventid uint) (events.Core, error) {
	var input Event
	result := er.db.Preload("Transactions").Preload("Reviews").Where("id = ? AND deleted_at IS NULL", eventid).Find(&input)
	if result.Error != nil {
		return events.Core{}, result.Error
	}
	if result.RowsAffected == 0 {
		return events.Core{}, result.Error
	}

	response := events.Core{
		ID:           input.ID,
		Title:        input.Title,
		Description:  input.Description,
		EventDate:    input.EventDate,
		EventTime:    input.EventTime,
		Status:       input.Status,
		Category:     input.Category,
		Location:     input.Location,
		Image:        input.Image,
		Hostedby:     input.Hostedby,
		Transactions: make([]events.Transaction, 0),
		Reviews:      make([]events.Reviews, 0),
	}

	err := er.db.Table("transactions").
		Joins("JOIN users ON transactions.user_id = users.id").
		Select("transactions.user_id, users.username, users.image").
		Where("transactions.event_id = ?", eventid).
		Scan(&response.Transactions).Error
	if err != nil {
		return events.Core{}, err
	}

	err = er.db.Table("reviews").
		Joins("JOIN users ON reviews.user_id = users.id").
		Select("reviews.user_id, users.username, users.image, reviews.review").
		Where("reviews.event_id = ?", eventid).
		Scan(&response.Reviews).Error
	if err != nil {
		return events.Core{}, err
	}

	return response, nil
}

func (er *EventRepository) UpdateEvent(id uint, updatedEvent events.Core) error {
	if err := er.db.Model(&Event{}).Where("id = ?", id).Updates(map[string]interface{}{
		"title":       updatedEvent.Title,
		"description": updatedEvent.Description,
		"event_date":  updatedEvent.EventDate,
		"event_time":  updatedEvent.EventTime,
		"status":      updatedEvent.Status,
		"category":    updatedEvent.Category,
		"location":    updatedEvent.Location,
		"image":       updatedEvent.Image,
		"hostedby":    updatedEvent.Hostedby,
		"updated_at":  time.Now(),
	}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
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

	if err := er.db.Model(&Event{}).Where("id = ?", id).Updates(map[string]interface{}{
		"DeletedAt": input.DeletedAt,
	}).Error; err != nil {
		return err
	}

	return nil
}
