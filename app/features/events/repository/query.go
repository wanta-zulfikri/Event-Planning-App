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
	input := Events{
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
		log.Println("Terjadi error saat membuat daftar event baru", err.Error())
		return events.Core{}, err
	}
	return newEvent, nil
}
