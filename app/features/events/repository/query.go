package repository

import (
	"log"

	"github.com/wanta-zulfikri/Event-Planning-App/app/features/events"
	"gorm.io/gorm"
)

type EventRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *EventRepository {
	return &EventRepository{db: db}
}

func (er *EventRepository) CreateEvent(newEvent events.Core) (events.Core, error) {
	input := events.Event{
		Title:       newEvent.Title,
		Description: newEvent.Description,
		EventDate:   newEvent.EventDate,
		EventTime:   newEvent.EventTime,
		Status:      newEvent.Status,
		Category:    newEvent.Category,
		Location:    newEvent.Location,
		Image:       newEvent.Image,
	}

	err := er.db.Table("events").Create(&input).Error
	if err != nil {
		log.Println("Error creating new event: ", err.Error())
		return events.Core{}, err
	}

	createdEvent := events.Core{
		Id:          input.ID,
		Title:       input.Title,
		Description: input.Description,
		EventDate:   input.EventDate,
		EventTime:   input.EventTime,
		Status:      input.Status,
		Category:    input.Category,
		Location:    input.Location,
		Image:       input.Image,
	}
	return createdEvent, nil
}
