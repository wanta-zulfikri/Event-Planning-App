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

func (s *EventService) CreateEventWithTickets(event events.Core, userID uint) error {
	err := s.r.CreateEventWithTickets(event, userID)
	if err != nil {
		return err
	}

	return nil
}

func (es *EventService) GetEvents() ([]events.Core, error) {
	events, err := es.r.GetEvents()
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (es *EventService) GetEventsByUserID(userid uint) ([]events.Core, error) {
	events, err := es.r.GetEventsByUserID(userid)
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (es *EventService) GetEvent(eventid uint) (events.Core, error) {
	event, err := es.r.GetEvent(eventid)
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
