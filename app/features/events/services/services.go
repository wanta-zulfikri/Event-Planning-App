package services

import (
	"github.com/wanta-zulfikri/Event-Planning-App/app/features/events"
)

type EventService struct {
	r events.Repository
}

func New(r events.Repository) events.Service {
	return &EventService{r: r}
}

func (s *EventService) CreateEventWithTickets(event events.Core, userID uint) error {
	err := s.r.CreateEventWithTickets(nil, event, userID)
	if err != nil {
		return err
	}

	return nil
}
