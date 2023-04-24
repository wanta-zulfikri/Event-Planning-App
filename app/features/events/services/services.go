package services

import (
	"errors"

	"github.com/wanta-zulfikri/Event-Planning-App/app/features/events"
	"gorm.io/gorm"
)

type EventService struct {
	r events.Repository
}

func New(r events.Repository) events.Service {
	return &EventService{r: r}
}

func (es *EventService) GetEvents() ([]events.Core, error) {
	events, err := es.r.GetEvents()
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (es *EventService) CreateEvent(newEvent events.Core, id uint) error {
	_, err := es.r.CreateEvent(newEvent, id)
	if err != nil {
		return err
	}
	return nil
}

func (es *EventService) GetEvent(id uint) (events.Core, error) {
	event, err := es.r.GetEvent(id)
	if err != nil {
		return events.Core{}, err
	}
	return event, nil
}

func (es *EventService) UpdateEvent(id uint, updatedEvent events.Core) error {
	updatedEvent.ID = id
	if err := es.r.UpdateEvent(id, updatedEvent); err != nil {
		return err
	}
	return nil
}

func (es *EventService) DeleteEvent(id uint) error {
	err := es.r.DeleteEvent(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		return err
	}
	return nil
}
