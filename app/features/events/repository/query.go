package repository

import (
	"errors"
	"time"

	"github.com/wanta-zulfikri/Event-Planning-App/app/features/events"
	"gorm.io/gorm"
)

type EventRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *EventRepository {
	return &EventRepository{db: db}
}

func (er *EventRepository) GetEvents() ([]events.Core, error) {
	var cores []events.Core
	if err := er.db.Table("events").Find(&cores).Error; err != nil {
		return nil, err
	}
	return cores, nil
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
		return events.Core{}, err
	}

	createdEvent := events.Core{
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

func (er *EventRepository) GetEvent(id uint) (events.Core, error) {
	var input events.Event
	result := er.db.Where("id = ?", id).Find(&input)
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
	}, nil
}

func (er *EventRepository) UpdateEvent(id uint, updatedEvent events.Core) error {
	input := events.Event{}
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
	input.UpdatedAt = time.Now()

	if err := er.db.Save(&input).Error; err != nil {
		return err
	}
	return nil
}

func (er *EventRepository) DeleteEvent(id uint) error {
	input := events.Event{}
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
