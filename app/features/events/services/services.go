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

func (es *EventService) CreateEvent(newEvent events.Core) error {
	_, err := es.r.CreateEvent(newEvent)
	if err != nil {
		return err
	}
	return nil
}
